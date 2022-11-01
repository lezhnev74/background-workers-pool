package worker

import "errors"

const (
	WORKER_READY  = iota // ready to accept a batch of tasks
	WORKER_BUSY          // working on the tasks, expect results asap
	WORKER_KILLED        // worker must be GCed

	ERROR_BUSY   = "worker is busy"
	ERROR_KILLED = "worker is killed"
)

// Worker is an abstraction on top of actual worker implementation (a dedicated process on this or foreign host)
type Worker interface {
	GetState() int
	// AcceptBatch should be nonblocking, it must send data to the worker and return instantly
	// the order or results should not be assumed
	AcceptBatch(batch []Task) (<-chan TaskResult, error)
	// Kill should free any used resources, no further call should be performed
	// the worker itself will be GC-ed
	Kill()
}

// WorkerBase is a process that handles tasks
type WorkerBase struct {
	state int
}

func (w *WorkerBase) GetState() int { return w.state }

type GoWorkerFunc func(tasks []Task) <-chan TaskResult

// GoWorker is a worker that uses a go routine to perform its work
// meant for tests only, real workers start separate OS processes
type GoWorker struct {
	WorkerBase
	handleBatch GoWorkerFunc
}

func (w *GoWorker) AcceptBatch(batch []Task) (<-chan TaskResult, error) {
	if w.state == WORKER_BUSY {
		return nil, errors.New(ERROR_BUSY)
	}
	if w.state == WORKER_KILLED {
		return nil, errors.New(ERROR_KILLED)
	}

	w.state = WORKER_BUSY
	results := make(chan TaskResult)
	go func() {
		// catch all goroutine results to control when the worker finishes
		goResults := w.handleBatch(batch)
		for goResult := range goResults {
			results <- goResult
		}
		close(results)
		w.state = WORKER_READY
	}()
	return results, nil
}

func (w *GoWorker) Kill() {
	w.state = WORKER_KILLED
}

func MakeGoroutineWorker(h GoWorkerFunc) *GoWorker {
	w := GoWorker{
		WorkerBase{
			state: WORKER_READY,
		},
		h,
	}
	return &w
}

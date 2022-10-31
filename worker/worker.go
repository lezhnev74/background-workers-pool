package worker

const (
	WORKER_READY = iota // ready to accept a batch of tasks
	WORKER_BUSY         // working on the tasks, expect results asap

	ERROR_BUSY = "worker is busy"
)

type Worker interface {
	GetState() int
	// AcceptBatch should be nonblocking, it must send data to the worker and return instantly
	// the order or results should not be assumed
	AcceptBatch(batch []Task) <-chan TaskResult
}

// WorkerBase is a process that handles tasks
type WorkerBase struct {
	state int
}

func (w *WorkerBase) GetState() int { return w.state }

type GoWorkerFunc func(tasks []Task, results chan<- TaskResult)

// GoWorker is a worker that uses a go routine to perform its work
// meant for tests only, real workers start separate OS processes
type GoWorker struct {
	WorkerBase
	handleBatch GoWorkerFunc
}

func (w *GoWorker) AcceptBatch(batch []Task) <-chan TaskResult {
	results := make(chan TaskResult)
	go w.handleBatch(batch, results)
	return results
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

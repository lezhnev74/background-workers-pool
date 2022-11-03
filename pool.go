package ipc

// Pool is a family of workers that handle only specific tasks
// It is a service, it tries to fill up to the capacity
type Pool struct {
	workers        []Worker // all currently allocated workers
	newWorker      workerCreator
	maxWorkers     int // pool capacity
	tasksBatchSize int // how many tasks a worker can accept at once
}

// A method to create one more worker in the pool
type workerCreator func() Worker

func MakePool(maxWorkers, taskBatchSize int, creator workerCreator) *Pool {
	return &Pool{
		workers:        make([]Worker, 0),
		newWorker:      creator,
		tasksBatchSize: taskBatchSize,
		maxWorkers:     maxWorkers,
	}
}

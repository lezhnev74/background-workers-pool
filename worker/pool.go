package worker

// Pool is a family of workers that handle only specific tasks
type Pool struct {
	workers []WorkerBase // all currently allocated workers
}

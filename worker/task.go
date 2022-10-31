package worker

const (
	TASK_RESULT_SUCCESS = iota
	TASK_RESULT_RESCHEDULE
	TASK_RESULT_FAILURE
)

type Task struct {
	id string
}

type TaskResult struct {
	id     string
	result int
}

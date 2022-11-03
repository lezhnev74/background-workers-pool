package ipc

import (
	"github.com/google/uuid"
	"time"
)

const (
	TASK_RESULT_SUCCESS = iota
	TASK_RESULT_RESCHEDULE
	TASK_RESULT_FAILURE
)

type Task struct {
	// unique id, probably UUID
	id string
	// a category of tasks, like "calculate_quarter_report"
	name string
	// some JSON-ish payload
	payload any
}

type TaskResult struct {
	id     string
	result int
}

// ReadForWork returns up to count tasks that are ready for work
// it locks tasks for the lockTimeout
type ReadForWork func(count int, lockTimeout time.Duration) []Task

func MakeTask(name string, payload any) *Task {
	return &Task{
		id:      uuid.New().String(),
		name:    name,
		payload: payload,
	}
}

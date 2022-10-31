package worker

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateWorker(t *testing.T) {
	t.Run("Create GoWorker", func(t *testing.T) {
		w := MakeGoroutineWorker(SuccessHandler)
		require.Equal(t, WORKER_READY, w.GetState())
	})
}

func TestFullCycle(t *testing.T) {
	t.Run("Full Cycle GoWorker", func(t *testing.T) {
		w := MakeGoroutineWorker(SuccessHandler)
		resultsChan := w.AcceptBatch(tasksBatch)

		results := make([]TaskResult, 0)
		for result := range resultsChan {
			results = append(results, result)
		}
		require.Len(t, results, len(tasksBatch))
	})
}

// SuccessHandler always returns successful results
func SuccessHandler(tasks []Task, results chan<- TaskResult) {
	for _, t := range tasks {
		results <- TaskResult{
			id:     t.id,
			result: TASK_RESULT_SUCCESS,
		}
	}
	close(results)
}

var tasksBatch = []Task{
	{"1"},
	{"2"},
	{"3"},
}

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

	// Test all worker implementations
	workers := []Worker{
		MakeGoroutineWorker(SuccessHandler),
	}

	for _, w := range workers {
		t.Run("Full Cycle", func(t *testing.T) {
			// Initially the worker is ready
			require.Equal(t, WORKER_READY, w.GetState())

			// After it has accepted the bach it becomes BUSY
			resultsChan, err := w.AcceptBatch(tasksBatch)
			require.NoError(t, err)
			require.Equal(t, WORKER_BUSY, w.GetState())

			// If one tries to give it more Tasks it returns an Error
			_, err = w.AcceptBatch(tasksBatch)
			require.Error(t, err, ERROR_BUSY)

			// Read all available results from the Worker's channel
			results := make([]TaskResult, 0)
			for result := range resultsChan {
				results = append(results, result)
			}
			require.Len(t, results, len(tasksBatch))

			// After it returns all the results it becomes READY again
			require.Equal(t, WORKER_READY, w.GetState())
		})
	}
}

// SuccessHandler always returns successful results
func SuccessHandler(tasks []Task) <-chan TaskResult {
	results := make(chan TaskResult)
	go func() {
		for _, t := range tasks {
			results <- TaskResult{
				id:     t.id,
				result: TASK_RESULT_SUCCESS,
			}
		}
		close(results)
	}()
	return results
}

var tasksBatch = []Task{
	{"1"},
	{"2"},
	{"3"},
}

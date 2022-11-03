package ipc

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPoolCreation(t *testing.T) {
	t.Run("Create a pool", func(t *testing.T) {
		p := MakePool(3, 1, createGoWorker)
		require.Equal(t, 3, p.maxWorkers)
		require.Equal(t, 1, p.tasksBatchSize)
	})
}

func createGoWorker() Worker {
	return MakeGoroutineWorker(SuccessHandler)
}

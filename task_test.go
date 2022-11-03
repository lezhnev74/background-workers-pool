package ipc

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTaskCreation(t *testing.T) {
	t.Run("create a task", func(t *testing.T) {
		task := MakeTask("calculate_report", []any{"2017"})
		require.Equal(t, "calculate_report", task.name)
		require.Equal(t, []any{"2017"}, task.payload)
	})
}

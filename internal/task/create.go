package task

import (
	"context"
	"fmt"

	"github.com/yuriongit/tday/internal/domain"
)

// Create creates and saves a new task (persistence planned).
func Create(
	rootCtx context.Context,
	inputHandler *domain.TaskInputHandler,
	taskIDGen *domain.TaskIDGenerator,
	db *domain.SupabaseDB,
) error {
	// Collect inputs from user.
	task := newTask(inputHandler, taskIDGen)

	// Persist user's task
	if err := db.InsertTask(rootCtx, task); err != nil {
		return fmt.Errorf("persistence error: %s\n", err)
	}

	// Output created task.
	outputNewTask(task)

	return nil
}

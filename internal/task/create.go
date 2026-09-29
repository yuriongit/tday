package task

import (
	"context"
	"fmt"

	"github.com/yuriongit/tday/internal/app"
)

// Create creates and saves a new task (persistence planned).
func Create(
	rootCtx context.Context,
	inputHandler *app.TaskInputHandler,
	taskIDGen *app.TaskIDGenerator,
	db *app.SupabaseDB,
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

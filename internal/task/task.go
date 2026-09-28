/*
Package task offers (CRUD) functionality for tasks.
*/
package task

import (
	"context"
	"fmt"
	"time"

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

// newTask instantiates a new Task struct.
func newTask(
	inputHandler *domain.TaskInputHandler,
	taskIDGen *domain.TaskIDGenerator,
) *domain.Task {
	id := taskIDGen.Generate()
	taskInputs := collectTaskInputs(inputHandler)

	return &domain.Task{
		UUID:      id,
		CreatedAt: time.Now(),
		InputData: &taskInputs,
	}
}

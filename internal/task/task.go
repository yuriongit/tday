/*
Package task offers (CRUD) functionality for tasks.
*/
package task

import (
	"time"

	"github.com/yuriongit/tday/internal/app"
	"github.com/yuriongit/tday/internal/domain"
)

// Create creates and saves a new task (persistence planned).
func Create(
  inputHandler *app.TaskInputHandler, 
  taskIDGen *app.TaskIDGenerator,
) {
  // Collect inputs from user.
	task := newTask(inputHandler, taskIDGen)
	// Output created task.
	outputNewTask(task) 
}

func newTask(
	inputHandler *app.TaskInputHandler,
	taskIDGen *app.TaskIDGenerator,
) *domain.Task {
	id := taskIDGen.Generate()
	taskInputs := collectTaskInputs(inputHandler)

	return &domain.Task{
		UUID:        id,
		CreatedAt: time.Now(),
		InputData: &taskInputs,
	}
}

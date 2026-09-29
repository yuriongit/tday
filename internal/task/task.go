/*
Package task offers (CRUD) functionality for tasks.
*/
package task

import (
	"time"

	"github.com/yuriongit/tday/internal/domain"
)

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

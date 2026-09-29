/*
Package task offers the CRUD functionality for
tasks, currently offers the functionality to
create and read all tasks.
*/
package task

import (
	"time"

	"github.com/yuriongit/tday/internal/app"
	"github.com/yuriongit/tday/internal/domain"
)

// newTask instantiates a new Task struct.
func newTask(
	inputHandler *app.TaskInputHandler,
	taskIDGen *app.TaskIDGenerator,
) *domain.Task {
	id := taskIDGen.Generate()
	taskInputs := collectTaskInputs(inputHandler)

	return &domain.Task{
		Metadata: domain.TaskMetadata{
			UUID:      id,
			CreatedAt: time.Now(),
		},
		InputData: &taskInputs,
	}
}

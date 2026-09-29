/*
Package task offers the CRUD functionality for
tasks, currently offers the functionality to
create and read all tasks.
*/
package task

import (
	"context"
	"fmt"

	"github.com/yuriongit/tday/internal/app"
)

/*
Create creates and persists a new task.
*/
func Create(
	rootCtx context.Context,
	inputHandler *app.TaskInputHandler,
	taskIDGen *app.TaskIDGenerator,
	db *app.SupabaseDB,
) error {
	// Collect inputs from user and creates a task
	task := newTask(inputHandler, taskIDGen)

	// Persist the created task
	if err := db.InsertTask(rootCtx, task); err != nil {
		return fmt.Errorf("persistence error: %s\n", err)
	}

	// Output created task
	outputNewTask(task)

	return nil
}

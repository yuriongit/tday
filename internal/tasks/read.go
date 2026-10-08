/*
Package tasks offers the CRUD functionality for
tasks, currently offers the functionality to
create and read all tasks.
*/
package tasks

import (
	"context"

	"github.com/yuriongit/tday/internal/app"
	"github.com/yuriongit/tday/internal/ui"
)

// ListAll lists all persisted tasks.
func ListAll(rootCtx context.Context, db app.Database) error {
	tasks, err := db.QueryAllTasks(rootCtx)
	if err != nil {
		return err
	}

	ui.ColoredTaskOutput.All(tasks)

	return nil
}

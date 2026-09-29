/*
Package task offers the CRUD functionality for
tasks, currently offers the functionality to
create and read all tasks.
*/
package task

import (
	"context"

	"github.com/yuriongit/tday/internal/app"
)

// ListAll lists all persisted tasks.
func ListAll(rootCtx context.Context, db app.Database) error {
	tasks, err := db.QueryAllTasks(rootCtx)
	if err != nil {
		return err
	}

	outputAllTasks(tasks)

	return nil
}

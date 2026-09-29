package task

import (
	"context"

	"github.com/yuriongit/tday/internal/app"
)

func ListAll(rootCtx context.Context, db app.Database) error {
	tasks, err := db.QueryAllTasks(rootCtx)
	if err != nil {
		return err
	}

	outputAllTasks(tasks)

	return nil
}

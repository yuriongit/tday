package task

import (
	"context"

	"github.com/yuriongit/tday/internal/domain"
)

func ListAll(rootCtx context.Context, db domain.Database) error {
	tasks, err := db.QueryAllTasks(rootCtx)
	if err != nil {
		return err
	}

	outputAllTasks(tasks)

	return nil
}

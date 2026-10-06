package tasks

import (
	"context"

	"github.com/yuriongit/tday/internal/app"
	"github.com/yuriongit/tday/internal/domain"
	"github.com/yuriongit/tday/internal/ui"
)

func MarkDone(
	rootCtx context.Context,
	id domain.ID,
	db *app.SupabaseDB,
) error {
	if err := db.CompleteTask(rootCtx, id); err != nil {
		return err
	}

	ui.TaskOutput.Complete(id)

	return nil
}

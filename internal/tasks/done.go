package tasks

import (
	"context"

	"github.com/yuriongit/tday/internal/app"
	"github.com/yuriongit/tday/internal/domain"
	"github.com/yuriongit/tday/internal/ui"
)

/*
MarkDone attempts to mark a task as complete and
outputs the status of the operation.
*/
func MarkDone(
	rootCtx context.Context,
	id domain.ID,
	db *app.SupabaseDB,
) error {
  if err := validateID(id); err != nil {
  	return err
  }

	if err := db.CompleteTask(rootCtx, id); err != nil {
		return err
	}

	ui.TaskOutput.Complete(id)

	return nil
}

package task

import (
	"context"

	"github.com/yuriongit/tday/internal/app"
	"github.com/yuriongit/tday/internal/domain"
	"github.com/yuriongit/tday/internal/ui"
)

func Remove(rootCtx context.Context, id domain.ID, db *app.SupabaseDB) error {
	if err := db.DeleteTask(rootCtx, id); err != nil {
		return err
	}
	
	ui.TaskOutput.Remove(id)
	
	return nil
}

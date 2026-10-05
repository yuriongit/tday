package task

import (
	"context"
	"fmt"

	"github.com/yuriongit/tday/internal/app"
	"github.com/yuriongit/tday/internal/domain"
	"github.com/yuriongit/tday/internal/ui"
)

// Update updates a task.
func Update(rootCtx context.Context, id domain.ID, db *app.SupabaseDB) error {
	existingTask, err := db.QueryTask(rootCtx, id)
	if err != nil {
		return err
	}

	updatedTaskInputs := collectTaskUpdateInputs(
		app.NewTaskInputHandler(),
		existingTask,
	)

	if err := db.UpdateTask(rootCtx, updatedTaskInputs, id); err != nil {
		return err
	}

	ui.TaskOutput.Update(id)

	return nil
}

func collectTaskUpdateInputs(
	inputHandler *app.TaskInputHandler,
	prevTask *domain.Task,
) (updatedTaskInputs *domain.TaskInputData) {
	scanner := inputHandler.Scanner
	input := make(domain.TaskInputData)

	// Iterate over defined fields
	for _, fieldDef := range domain.AllFields {
		for {
			fmt.Printf(
				"i. Before %s:\n  < %q\n",
				fieldDef.Name,
				(*prevTask.InputData)[string(fieldDef.ID)],
			)
			fmt.Printf("?. %s:\n  > ", fieldDef.Name)

			if !scanner.Scan() {
				break
			}

			value := scanner.Text()

			if value == "" {
				// skip to next field because of updates
				break
			}

			// Validate using the field's validation function
			if err := fieldDef.Validate(value, true); err != nil {
				fmt.Printf("✗ %s\n", err.Error())
				// retry logic here
				continue
			}

			input[string(fieldDef.ID)] = value
			break
		}
	}

	return &input
}

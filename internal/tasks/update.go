/*
Package tasks offers the CRUD functionality for
tasks, currently offers the functionality to
create and read all tasks.
*/
package tasks

import (
	"context"
	"fmt"

	"github.com/yuriongit/tday/internal/app"
	"github.com/yuriongit/tday/internal/domain"
	"github.com/yuriongit/tday/internal/ui"
)

// Update updates a task.
func Update(rootCtx context.Context, id domain.ID, db *app.SupabaseDB) error {
	if err := validateID(id); err != nil {
		return err
	}

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

	ui.ColoredTaskOutput.Update(id)

	return nil
}

/*
collectTaskUpdateInputs collects user input to update existing task fields.
Displays the previous field value prefixed with git-style "<<" rendered entirely in red.
*/
func collectTaskUpdateInputs(
	inputHandler *app.TaskInputHandler,
	prevTask *domain.Task,
) (updatedTaskInputs *domain.TaskInputData) {
	scanner := inputHandler.Scanner
	input := make(domain.TaskInputData)

	// Iterate over defined fields
	for _, fieldDef := range domain.AllFields {
		for {
			prevVal := (*prevTask.InputData)[string(fieldDef.ID)]

			// Render the entire "Before" output block in red using git-style <<
			beforeMsg := fmt.Sprintf("  << Before: %q", prevVal)
			fmt.Println(ui.RedOutput.Render(beforeMsg))

			// Prompt line with >> arrows
			fmt.Printf(
				"%s %s\n  %s ",
				ui.NewTaskFieldLabelStyle.Faint(true).Render("?"),
				ui.NewTaskFieldLabelStyle.Render(fieldDef.Name),
				ui.BlueArrows,
			)

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
				fmt.Println(ui.ErrStyle.Render(fmt.Sprintf("✗ %s", err.Error())))
				// retry logic here
				continue
			}

			input[string(fieldDef.ID)] = value
			break
		}
	}

	return &input
}

//
// func collectTaskUpdateInputs(
// 	inputHandler *app.TaskInputHandler,
// 	prevTask *domain.Task,
// ) (updatedTaskInputs *domain.TaskInputData) {
// 	scanner := inputHandler.Scanner
// 	input := make(domain.TaskInputData)
//
// 	// Iterate over defined fields
// 	for _, fieldDef := range domain.AllFields {
// 		for {
// 			fmt.Printf(
// 				"i. Before %s:\n  < %q\n",
// 				fieldDef.Name,
// 				(*prevTask.InputData)[string(fieldDef.ID)],
// 			)
// 			fmt.Printf("?. %s:\n  > ", fieldDef.Name)
//
// 			if !scanner.Scan() {
// 				break
// 			}
//
// 			value := scanner.Text()
//
// 			if value == "" {
// 				// skip to next field because of updates
// 				break
// 			}
//
// 			// Validate using the field's validation function
// 			if err := fieldDef.Validate(value, true); err != nil {
// 				fmt.Printf("✗ %s\n", err.Error())
// 				// retry logic here
// 				continue
// 			}
//
// 			input[string(fieldDef.ID)] = value
// 			break
// 		}
// 	}
//
// 	return &input
// }

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
		return fmt.Errorf("✗ %s", err)
	}

	// Output created task
	ui.ColoredTaskOutput.New(&task.Metadata.UUID)

	return nil
}

/*
collectTaskInputs collects user input to fill out new task fields.
*/
func collectTaskInputs(
	inputHandler *app.TaskInputHandler,
) domain.TaskInputData {
	scanner := inputHandler.Scanner
	input := make(domain.TaskInputData)

	// Iterate over defined fields
	for _, fieldDef := range domain.AllFields {
		for {
			if !fieldDef.Required {
				fmt.Printf(
					"%s %s %s\n  %s ",
					ui.NewTaskFieldLabelStyle.Faint(true).Render("?"),
					ui.NewTaskFieldLabelStyle.Render(fieldDef.Name),
					ui.OptLabelStyle.Render("(opt.)"),
					ui.BlueArrow,
				)
			} else {
				fmt.Printf(
					"%s %s\n  %s ",
					ui.NewTaskFieldLabelStyle.Faint(true).Render("?"),
					ui.NewTaskFieldLabelStyle.Render(fieldDef.Name),
					ui.BlueArrow,
				)
			}

			if !scanner.Scan() {
				break
			}

			value := scanner.Text()

			// Validate using the field's validation function
			if err := fieldDef.Validate(value, false); err != nil {
				fmt.Println(ui.ErrStyle.Render(fmt.Sprintf("✗ %s", err.Error())))
				// retry logic here
				continue
			}

			input[string(fieldDef.ID)] = value
			break
		}
	}

	return input
}

// Non-colored output
// /*
// collectTaskInputs collects user input to
// fill out the task's fields. In the process
// of collecting this data, each field's value
// is updated through direct access to the
// domain.
// */
// func collectTaskInputs(
// 	inputHandler *app.TaskInputHandler,
// ) domain.TaskInputData {
// 	scanner := inputHandler.Scanner
// 	input := make(domain.TaskInputData)
//
// 	// Iterate over defined fields
// 	for _, fieldDef := range domain.AllFields {
// 		for {
// 			if !fieldDef.Required {
// 				fmt.Printf("?. %s (opt.):\n  > ", fieldDef.Name)
// 			} else {
// 				fmt.Printf("?. %s:\n  > ", fieldDef.Name)
// 			}
//
// 			if !scanner.Scan() {
// 				break
// 			}
//
// 			value := scanner.Text()
//
// 			// Validate using the field's validation function
// 			if err := fieldDef.Validate(value, false); err != nil {
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
// 	return input
// }

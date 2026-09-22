/*
Package task is responsible for all input.
*/
package task

import (
	"fmt"

	"github.com/yuriongit/tday/internal/app"
	"github.com/yuriongit/tday/internal/domain"
)

/*
collectTaskInputs collects user input to
fill out the task's fields. In the process
of collecting this data, each field's value
is updated through direct access to the
domain.
*/
func collectTaskInputs(inputHandler *app.TaskInputHandler) domain.TaskInputData {
	scanner := inputHandler.Scanner
	input := make(domain.TaskInputData)

	// Iterate over defined fields
	for _, fieldDef := range domain.AllFields {
		for {
			if !fieldDef.Required {
				fmt.Printf("? %s (opt.):\n  > ", fieldDef.Name)
			} else {
				fmt.Printf("? %s:\n  > ", fieldDef.Name)
			}

			if !scanner.Scan() {
				break
			}

			value := scanner.Text()

			// Validate using the field's validation function
			if err := fieldDef.Validate(value); err != nil {
				fmt.Printf("Invalid %s: %v\n", fieldDef.Name, err)
				// retry logic here
				continue
			}

			input[fieldDef.ID] = value
			break
		}
	}
	return input
}

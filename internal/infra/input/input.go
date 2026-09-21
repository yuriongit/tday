/*
Package input is responsible for all input.
*/
package input

import (
	"time"

	"github.com/yuriongit/tday/internal/domain"
	"github.com/yuriongit/tday/internal/infra/output"
)

/*
CollectTaskFields prompts the user for field input.
In the process of collecting this data, each field's
value is updated through direct access to the
domain.TaskInputHandler struct handed by 
instantiation.
*/
func CollectTaskFields(t *domain.TaskInputHandler) *domain.Task {
	// Iterate over fields for collection:
	for _, v := range t.Fields {
		// Prompt user for field.
		output.RequestTaskInput(t.Scanner, v)

		// Retrieve user input.
		fieldVal := domain.TaskFieldValue(t.Scanner.Text())
		// Append new value to i.Values slice.
		t.Values = append(t.Values, fieldVal)
	}
	return &domain.Task{
		CreationDate: time.Now().Format("Jan 2, 2006"),
		Fields: t.Fields,
		Values: t.Values,
	}
}

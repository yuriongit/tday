/*
Package task offers (CRUD) functionality for tasks.
*/
package task

import (
	"bufio"
	"os"

	"github.com/yuriongit/tday/internal/domain"
	"github.com/yuriongit/tday/internal/infra/input"
	"github.com/yuriongit/tday/internal/infra/output"
)

var (
	taskFields = []domain.TaskField{
		"Label",
		"Title",
		"Description",
		"Due Date",
	}
)

// Create creates a new task.
func Create() {
	inpHandler := NewTaskInputHandler()         // Create new required inputs.
	task := input.CollectTaskFields(inpHandler) // Collect inputs from user.
	output.CreatedTask(task)                    // Output created task.
}

/*
NewTaskInputHandler instantiates a new task .
*/
func NewTaskInputHandler() *domain.TaskInputHandler {
	return &domain.TaskInputHandler{
		Scanner: bufio.NewScanner(os.Stdin),
		Fields:  taskFields,
		Values:  []domain.TaskFieldValue{},
	}
}

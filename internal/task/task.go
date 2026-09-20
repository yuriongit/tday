/*
Package task offers (CRUD) functionality for tasks.
*/
package task

import (
	"bufio"
	"os"

	"github.com/yurongit/tday/internal/domain"
	"github.com/yurongit/tday/internal/infra/input"
	"github.com/yurongit/tday/internal/infra/output"
)

// Create creates a new task.
func Create() {
	// Create scanner for inputs.
	scanner := bufio.NewScanner(os.Stdin)

	// Create new required inputs.
	inp := input.NewInputs(
		*scanner,
		[]domain.Field{"Title", "Description"},
	)
	input.Collect(inp) // Collect inputs from user.

	output.TaskCreation(inp) // Output created task.
}

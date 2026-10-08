/*
Package tasks offers the CRUD functionality for
tasks, currently offers the functionality to
create and read all tasks.
*/
package tasks

import (
	"fmt"

	"github.com/yuriongit/tday/internal/domain"
)

func validateID(id domain.ID) error {

	if len(id) != domain.TaskIDLen {
		return fmt.Errorf(
			"Invalid task ID; ID must be %d characters",
			domain.TaskIDLen,
		)
	}

	return nil
}

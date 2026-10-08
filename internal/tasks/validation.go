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

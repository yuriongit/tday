package task

import (
	"bufio"
	"os"
	"tday/internal/domain"
)

func Create() {
	// create scanner for inputs
	scanner := bufio.NewScanner(os.Stdin)

	// create new required inputs
	inp := domain.NewInputs(
		*scanner,
		[]domain.Field{"Title", "Description"},
	)
	inp.Collect() // collect inputs from user

	// output created task
	out := domain.NewOutput()
	out.TaskCreation(inp.Fields, inp.Values)
}

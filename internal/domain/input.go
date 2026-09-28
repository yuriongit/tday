/*
Package domain contains the domain's core
types and dependencies.
*/
package domain

import (
	"bufio"
	"os"
)

/*
TaskInputHandler defines the structure of
task's input.
*/
type TaskInputHandler struct {
	Scanner *bufio.Scanner
}

/*
NewTaskInputHandler instantiates a new
TaskInputHandler.
*/
func NewTaskInputHandler() *TaskInputHandler {
	return &TaskInputHandler{
		Scanner: bufio.NewScanner(os.Stdin),
	}
}

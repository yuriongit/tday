package app

import (
	"bufio"
	"os"
)

/*
TaskInputHandler defines the structure of task's input.
*/
type TaskInputHandler struct {
  Scanner *bufio.Scanner
}

/*
NewTaskInputHandler instantiates a new task .
*/
func NewInputHandler() (*TaskInputHandler) {
  return &TaskInputHandler{
    Scanner: bufio.NewScanner(os.Stdin),
  }
}

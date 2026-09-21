/*
Package domain is responsible for business
logic, pure data representation, and
constants.
*/
package domain

import (
	"bufio"
)

/*
TaskField is a branded type of type string.
It defines the type of a TaskField's value.
*/
type TaskField string

/*
TaskFieldValue is a branded type of type string.
It defines the type of a Field's value.
*/
type TaskFieldValue string

/*
TaskInputHandler defines the structure of task's input.
*/
type TaskInputHandler struct {
	Scanner *bufio.Scanner
	Fields  []TaskField
	Values  []TaskFieldValue
}

type Task struct {
	CreationDate string
	Fields       []TaskField
	Values       []TaskFieldValue
}

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
Field is a branded type of type string.
It defines the type of a Field's value.
*/
type Field string

/*
FieldValue is a branded type of type string.
It defines the type of a Field's value.
*/
type FieldValue string

/*
Input defines the structure of task's input.
*/
type Input struct {
	Scanner bufio.Scanner
	Fields  []Field
	Values  []FieldValue
}

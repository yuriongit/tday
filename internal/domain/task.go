package domain

import (
	"bufio"
	"fmt"
)

type Field string
type FieldValue string

type Task struct {
	title Field
	desc  Field
}

type Input struct {
	Scanner bufio.Scanner
	Fields  []Field
	Values  []FieldValue
}

func NewInputs(
	scanner bufio.Scanner,
	fields []Field,
) *Input {
	return &Input{
		Scanner: scanner,
		Fields:  fields,
		Values:  []FieldValue{},
	}
}

func (i *Input) Collect() {
	for _, v := range i.Fields {
		// Collect user input
		fmt.Printf("? — %s:\n  > ", v)
		i.Scanner.Scan()

		val := FieldValue(i.Scanner.Text())
		i.Values = append(i.Values, val)
	}
}

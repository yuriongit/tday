/*
Package input is responsible for all input.
*/
package input

import (
	"bufio"
	"fmt"

	"github.com/yurongit/tday/internal/domain"
)

/*
NewInputs instantiates a new domain.Input struct
with the Values field automatically attached
*/
func NewInputs(
	scanner bufio.Scanner,
	fields []domain.Field,
) *domain.Input {
	return &domain.Input{
		Scanner: scanner,
		Fields:  fields,
		Values:  []domain.FieldValue{},
	}
}

/*
Collect prompts for user input and collects
the fields listed in Input
*/
func Collect(i *domain.Input) {
	for _, v := range i.Fields {
		// Collect user input
		fmt.Printf("? — %s:\n  > ", v)
		i.Scanner.Scan()

		val := domain.FieldValue(i.Scanner.Text())
		i.Values = append(i.Values, val)
	}
}

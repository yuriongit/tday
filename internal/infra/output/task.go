package output

import (
	"bufio"
	"fmt"

	"github.com/yuriongit/tday/internal/domain"
)

func RequestTaskInput(
	scanner *bufio.Scanner,
	field domain.TaskField,
) {
	fmt.Printf("? %s:\n  > ", field)
	scanner.Scan()
}

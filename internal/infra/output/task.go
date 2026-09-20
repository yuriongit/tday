package output

import (
	"bufio"
	"fmt"

	"github.com/yurongit/tday/internal/domain"
)

func RequestTaskInput(
  scanner *bufio.Scanner,
  field domain.Field,
) {
  fmt.Printf("? %s:\n  > ", field)
  scanner.Scan()
}
/*
Package output is responsible for streaming output
to the terminal.
*/
package output

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/yurongit/tday/internal/domain"
)

func clearTerminal() {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		return
	}
}

/*
TaskCreation outputs the created task and extra
relative information to the task.
*/
func TaskCreation(i *domain.Input) {
	clearTerminal()
	fmt.Println("...")
	fmt.Println("New task created,\nTask details include:")

	fmt.Println("———————")
	fmt.Printf(
		"%s | %s\n",
		time.Now().Format("03:04 PM"),
		time.Now().Format("Jan 2, 2006"),
	)

	for idx, v := range i.Values {
		fmt.Printf("• %s:\n  > %s%s%s\n",
			i.Fields[idx],
			domain.QuoteSymbol,
			v,
			domain.QuoteSymbol)
	}
	fmt.Println("———————")
}

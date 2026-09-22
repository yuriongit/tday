/*
Package output is responsible for streaming output
to the terminal.
*/
package task

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/yuriongit/tday/internal/domain"
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
outputNewTask outputs the created task and extra
relative information to the task.
*/
func outputNewTask(t *domain.Task) {
	clearTerminal()
	fmt.Println("...")
	fmt.Println("New task created!\nTask details include:")

	fmt.Println("————————————————————————————")

	fmt.Println("Metadata:")
	fmt.Printf("i. UUID:\n  > %s\n", t.UUID)
	fmt.Printf(
		"i. Created at:\n  > %s\n  > %s\n",
		time.Now().Format("Jan 2, 2006"),
		time.Now().Format(domain.TimeLayouts[1]),
	)

	fmt.Print("——————————————|")

	fmt.Println("\nData:")
	for _, field := range domain.AllFields {
  	v, exists := (*t.InputData)[field.ID]
   
		if !exists || v == "" {
		  continue
		}
		
		fmt.Printf("• %s:\n", field.Name)
		
		if field.ID == "due_at" {
  		fmt.Printf("  > %v\n", v)
		} else {
     	fmt.Printf("  > %s%v%s\n", domain.QuoteSymbol, v, domain.QuoteSymbol)
		}
	} 
	
	fmt.Println("————————————————————————————")
}

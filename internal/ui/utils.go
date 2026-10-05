package ui

// import (
// 	"os"
// 	"os/exec"
// 	"runtime"
// )

// /*
// clearTerminal clears the previous output from
// the terminal.
// */
// func clearTerminal() {
// 	var cmd *exec.Cmd
//
// 	if runtime.GOOS == "windows" {
// 		cmd = exec.Command("cmd", "/c", "cls")
// 	} else {
// 		cmd = exec.Command("clear")
// 	}
//
// 	cmd.Stdout = os.Stdout
// 	if err := cmd.Run(); err != nil {
// 		return
// 	}
// }

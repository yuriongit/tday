package domain

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

type Output struct {
	fields      []Field
	fieldValues []FieldValue
}

func NewOutput() *Output {
	return &Output{
		fields:      []Field{},
		fieldValues: []FieldValue{},
	}
}

func clearTerminal() {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}

func (o *Output) TaskCreation(
	Fields []Field,
	FieldValues []FieldValue,
) {
	clearTerminal()
	fmt.Println("...")
	fmt.Println("New task created,\nTask details include:")
	fmt.Println("——————")

	for idx, v := range FieldValues {
		fmt.Printf("%s:\n%s> %s%s%s\n",
			Fields[idx],
			Indent,
			QuoteSymbol,
			v,
			QuoteSymbol)
	}
}

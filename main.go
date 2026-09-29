/*
Package main is the entry point of TDay.

Copyright © 2026 Yuri Okeren <yuri.dev44@outlook.com>
*/
package main

import (
	"log"

	"github.com/yuriongit/tday/cmd"
	"github.com/yuriongit/tday/internal/app"
)

func main() {
	app, err := app.InitApp()
	if err != nil {
		log.Fatal(err)
	}

	// Set global app for cmd package to use
	cmd.SetApp(app)

	// Handle cleanup on exit
	defer app.Cancel()
	defer app.Database.Pool.Close()

	cmd.Execute()
}

// Delete, Update, ReadOne, Read, ReadRange(x-y)

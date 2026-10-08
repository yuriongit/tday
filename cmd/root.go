/*
Package cmd holds all of TDay's commands.

Copyright © 2026 Yuri Okeren <yuri.dev44@outlook.com>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	app "github.com/yuriongit/tday/internal/app"
	"github.com/yuriongit/tday/internal/ui"
)

var globalApp *app.App

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tday",
	Short: `My personal task-managing CLI tool`,
	Long: `TDay: A personal task-managing CLI tool for keep tracking of what I need
done for the day. It's quick, feather-weight, and simply straightforward.

TDay is meant for me to manage my tasks in a manner as simple and as 
straightfoward as it actually should be.

To get started, run: 

  tday init

`,
	PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
		// Skip app init for "init" command
		if cmd.Name() == "init" || cmd.Name() == "help" {
			return nil
		}

		// Initialize app for all other commands
		var err error
		globalApp, err = app.InitApp()
		if err != nil {
			return err
		}

		SetApp(globalApp)
		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	// Global cleanup
	if globalApp != nil {
		defer globalApp.Cancel()
		defer globalApp.Database.Pool.Close()
	}

	err := rootCmd.Execute()
	if err != nil {
    errMsg := ui.ErrStyle.Render(fmt.Sprintf("✗ %s", err.Error())) + "\n"
    fmt.Fprint(os.Stderr, errMsg)
		return
	}
}

func init() {
	// Suppress usage output
	rootCmd.SilenceUsage = true
	rootCmd.SilenceErrors = true
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.tday.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

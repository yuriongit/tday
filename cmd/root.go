/*
Package cmd holds all of TDay's commands

Copyright © 2026 Yuri Okeren <yuri.dev44@outlook.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tday",
	Short: `My personal task-managing CLI tool`,
	Long: `TDay: A personal task-managing CLI tool for keep tracking of what I need
done for the day. It's quick, feather-weight, and simply straightforward.

I've decided to create TDay for multiple reasons. I find it critical to
reduce as much extensive mouse-use where possible. Additionally, GUI-
based task-managers are shipped with too much: distracting UIs, bloat, 
a lack of simplicity, sluggishness, and a mouse-oriented UX. Tday is to
make managing my tasks as simple and as straightfoward as it actually
should be.
`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.tday.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

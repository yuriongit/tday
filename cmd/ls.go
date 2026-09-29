/*
Package cmd holds all of TDay's commands.

Copyright © 2026 Yuri Okeren <yuri.dev44@outlook.com>.
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yuriongit/tday/internal/task"
)

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	// Rename 'ls' command to 'la' for listing all tasks
	Use:   "ls",
	Short: "lists out all tasks",
	Long:  `TODO: Implement later`,
	RunE: func(_ *cobra.Command, _ []string) error {
		app := GetApp()

		return task.ListAll(app.Ctx, app.Database)
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// lsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// lsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

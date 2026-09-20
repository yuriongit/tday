/*
Package cmd holds all of TDay's commands

Copyright © 2026 Yuri Okeren <yuri.dev44@outlook.com>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yurongit/tday/internal/task"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "creates a new task",
	Long:  `TODO: Implement later`,
	Run: func(_ *cobra.Command, _ []string) {
		task.Create()
	},
}

func init() {
	rootCmd.AddCommand(newCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

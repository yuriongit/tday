/*
Package cmd holds all of TDay's commands.

Copyright © 2026 Yuri Okeren <yuri.dev44@outlook.com>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yuriongit/tday/internal/domain"
	"github.com/yuriongit/tday/internal/task"
)

// newCmd represents the new command.
var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Removes an existing task",
	Long:  `TODO: Implement later`,
	RunE: func(_ *cobra.Command, args []string) error {
		app := GetApp()
		id := domain.ID(args[0])

		return task.Remove(
			app.Ctx,
			id,
			app.Database,
		)
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

/*
Package cmd holds all of TDay's commands.

Copyright © 2026 Yuri Okeren <yuri.dev44@outlook.com>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/yuriongit/tday/internal/domain"
	"github.com/yuriongit/tday/internal/tasks"
)

// doneCmd represents the done command.
var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "Marks a task as complete",
	Long:  `TODO: Implement later`,
	RunE: func(_ *cobra.Command, args []string) error {
		app := GetApp()

		if len(args) == 0 {
			return fmt.Errorf("No task ID provided")
		}

		id := domain.ID(args[0])

		return tasks.MarkDone(
			app.Ctx,
			id,
			app.Database,
		)
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)
}

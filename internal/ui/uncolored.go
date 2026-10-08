/*
Package ui provides the UI for TDay.
*/
package ui

import (
	"fmt"

	"github.com/yuriongit/tday/internal/domain"
)

/*
uncoloredOutputNewTask outputs the created task and it's
metadata in a formatted manner.
*/
func uncoloredOutputNewTask(id *domain.ID) {
	fmt.Println("———————————————————————————————————")
	fmt.Printf("✓ Task %q created\n", id)
}

/*
uncoloredOutputAllTasks outputs all the persisted tasks in
a formatted manner.
*/
func uncoloredOutputAllTasks(tasks []domain.Task) {
	tasksAmt := len(tasks)

	fmt.Printf(
		"%d total / %d remaining / %d complete\n",
		tasksAmt,
		tasksAmt,
		0, // TODO: When tasks can be marked as complete
	)
	fmt.Println("—————————————————————————————————————")

	for idx, task := range tasks {
		label := (*task.InputData)[string(domain.LabelField)]
		title := (*task.InputData)[string(domain.TitleField)]
		desc := (*task.InputData)[string(domain.DescriptionField)]
		dueAt := (*task.InputData)[string(domain.DueAtField)]
		completedAt := task.Metadata.CompletedAt
		createdAt := task.Metadata.CreatedAt.Format(domain.TimeLayouts[1])
		uuid := task.Metadata.UUID

		fmt.Printf("{%s} %s: %q\n", uuid, label, title)

		if desc != "" {
			fmt.Printf("  > Desc: %q\n", desc)
		}

		if !completedAt.IsZero() {
			fmt.Printf("  > Completed at: %s (Created at: %s)\n", completedAt.Format(domain.TimeLayouts[1]), createdAt)
		} else {
			fmt.Printf("  > Due at: %s (Created at: %s)\n", dueAt, createdAt)
		}

		if idx != tasksAmt-1 {
			fmt.Println()
		}
	}
}

/*
uncoloredOutputDeletedTask outputs the created task and it's
metadata in a formatted manner.
*/
func uncoloredOutputDeletedTask(id domain.ID) {
	fmt.Printf("✓ Task %q deleted\n", id)
}

func uncoloredOutputUpdatedTask(id domain.ID) {
  fmt.Println(Divider)
	fmt.Printf("✓ Task %q updated\n", id)
}

func uncoloredOutputCompleteTask(id domain.ID) {
	fmt.Printf("✓ Task %q marked as complete\n", id)
}

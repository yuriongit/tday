package ui

import (
	"fmt"

	"github.com/yuriongit/tday/internal/domain"
)

/*
outputNewTask outputs the created task and it's
metadata in a formatted manner.
*/
func noColorOutputNewTask(id *domain.ID) {
	fmt.Println("———————————————————————————————————")
	fmt.Printf("✓ Task %q created successfully.\n", id)
}

/*
outputAllTasks outputs all the persisted tasks in
a formatted manner.
*/
func noColorOutputAllTasks(tasks []domain.Task) {
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
outputDeletedTask outputs the created task and it's
metadata in a formatted manner.
*/
func noColorOutputDeletedTask(id domain.ID) {
	fmt.Printf("✓ Task %q deleted successfully.\n", id)
}

func noColorOutputUpdatedTask(id domain.ID) {
	fmt.Printf("✓ Task %q updated successfully.\n", id)
}

func noColorOutputCompleteTask(id domain.ID) {
	fmt.Printf("✓ Task %q successfully marked as complete\n", id)
}

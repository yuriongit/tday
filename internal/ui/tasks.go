/*
Package ui provides the UI for TDay.
*/
package ui

import (
	"fmt"

	"github.com/yuriongit/tday/internal/domain"
)

// Task defines the structure for task output functions.
type Task struct {
	New      func(id *domain.ID)
	All      func(t []domain.Task)
	Update   func(id domain.ID)
	Complete func(id domain.ID)
	Remove   func(id domain.ID)
}

/*
TaskOutput holds the functions responsible for any task-
related output (besides input collection).
*/
var TaskOutput = Task{
	New:      outputNewTask,
	All:      outputAllTasks,
	Update:   outputUpdatedTask,
	Complete: outputCompleteTask,
	Remove:   outputDeletedTask,
}

/*
outputNewTask outputs the created task and it's
metadata in a formatted manner.
*/
func outputNewTask(id *domain.ID) {
	fmt.Println("———————————————————————————————————")
	fmt.Printf("✓ Task %q created successfully.\n", id)
}

/*
outputAllTasks outputs all the persisted tasks in
a formatted manner.
*/
func outputAllTasks(tasks []domain.Task) {
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
func outputDeletedTask(id domain.ID) {
	fmt.Printf("✓ Task %q deleted successfully.\n", id)
}

func outputUpdatedTask(id domain.ID) {
	fmt.Printf("✓ Task %q updated successfully.\n", id)
}

func outputCompleteTask(id domain.ID) {
	fmt.Printf("✓ Task %q successfully marked as complete\n", id)
}

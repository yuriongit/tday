/*
Package task offers the CRUD functionality for
tasks, currently offers the functionality to
create and read all tasks.
*/
package ui

import (
	"fmt"
	"time"

	"github.com/yuriongit/tday/internal/domain"
)

type Task struct {
	New    func(t *domain.Task)
	All    func(t []domain.Task)
	Remove func(id domain.ID)
}

var TaskOutput = Task{
	New:    outputNewTask,
	All:    outputAllTasks,
	Remove: outputDeletedTask,
}

/*
outputNewTask outputs the created task and it's
metadata in a formatted manner.
*/
func outputNewTask(t *domain.Task) {
	clearTerminal()
	fmt.Println("...")
	fmt.Println("New task created!\nTask details include:")

	fmt.Println("————————————————————————————")

	fmt.Println("Metadata:")
	fmt.Printf("i. UUID:\n  > %s\n", t.Metadata.UUID)
	fmt.Printf(
		"i. Created at:\n  > %s\n  > %s\n",
		time.Now().Format("Jan 2, 2006"),
		time.Now().Format(domain.TimeLayouts[1]),
	)

	fmt.Print("——————————————|")

	fmt.Println("\nData:")
	for _, field := range domain.AllFields {
		v, exists := (*t.InputData)[field.ID]

		if !exists || v == "" {
			continue
		}

		fmt.Printf("• %s:\n", field.Name)

		if field.ID == "due_at" {
			fmt.Printf("  > %v\n", v)
		} else {
			fmt.Printf("  > %q\n", v)
		}
	}

	fmt.Println("————————————————————————————")
}

/*
outputAllTasks outputs all the persisted tasks in
a formatted manner.
*/
func outputAllTasks(tasks []domain.Task) {
	if len(tasks) == 0 {
		fmt.Println("No tasks found")
		return
	}

	tasksAmt := len(tasks)

	fmt.Printf(
		"%d total / %d remaining / %d complete\n",
		tasksAmt,
		tasksAmt,
		0, // TBD: When tasks can be marked as complete
	)
	fmt.Println("—————————————————————————————————————")

	for idx, task := range tasks {
		label := (*task.InputData)[domain.LabelField]
		title := (*task.InputData)[domain.TitleField]
		desc := (*task.InputData)[domain.DescriptionField]
		dueAt := (*task.InputData)[domain.DueAtField]
		completedAt := task.Metadata.CompletedAt
		time := task.Metadata.CreatedAt.Format(domain.TimeLayouts[1])

		fmt.Printf("• {%d} %s: %q\n", (idx + 1), label, title)
		if desc != "" {
			fmt.Printf("   > Desc: %q\n", desc)
		}
		if !completedAt.IsZero() {
			fmt.Printf("   > Completed at: %s\n", completedAt.Format(domain.TimeLayouts[1]))
		}
		fmt.Printf("   > Created at: %s | Due at: %s\n\n", time, dueAt)
	}
}

/*
outputDeletedTask outputs the created task and it's
metadata in a formatted manner.
*/
func outputDeletedTask(id domain.ID) {
	fmt.Printf("✓ Task %q deleted successfully.\n", id)
}

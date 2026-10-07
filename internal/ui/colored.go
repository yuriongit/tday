/*
Package ui provides the UI for TDay.
*/
package ui

import (
	"fmt"

	"github.com/yuriongit/tday/internal/domain"
)

/*
coloredOutputNewTask outputs the created task and its
metadata in a formatted manner.
*/
func coloredOutputNewTask(id *domain.ID) {
  fmt.Println(Divider)
	// Only the ">>" arrows are styled in green; remainder of the text is unstyled
	fmt.Printf("%s Task %s created successfully.\n", GreenArrows, IDStyle.Render(id.String()))
}

/*
coloredOutputAllTasks outputs all the persisted tasks in
a formatted manner.
*/
func coloredOutputAllTasks(tasks []domain.Task) {
	tasksAmt := len(tasks)

	fmt.Printf(
		"%s total / %s remaining / %s complete\n",
		BoldNum.Render(fmt.Sprintf("%d", tasksAmt)),
		DueStyle.Render(fmt.Sprintf("%d", tasksAmt)),
		DoneStyle.Render("0"), // TODO: When tasks can be marked as complete
	)
	fmt.Println(Divider)

	for idx, task := range tasks {
		label := (*task.InputData)[string(domain.LabelField)]
		title := (*task.InputData)[string(domain.TitleField)]
		desc := (*task.InputData)[string(domain.DescriptionField)]
		dueAt := (*task.InputData)[string(domain.DueAtField)]
		completedAt := task.Metadata.CompletedAt
		createdAt := task.Metadata.CreatedAt.Format(domain.TimeLayouts[1])
		uuid := task.Metadata.UUID

		fmt.Printf(
			"%s %s: %s\n",
			UUIDStyle.Render(fmt.Sprintf("{%s}", uuid)),
			LabelStyle.Render(label.(string)),
			TitleStyle.Render(fmt.Sprintf("%q", title)),
		)

		if desc != "" {
			fmt.Printf("  %s Desc: %s\n", MutedStyle.Render(">"), DescStyle.Render(fmt.Sprintf("%q", desc)))
		}

		if !completedAt.IsZero() {
			fmt.Printf(
				"  %s Completed at: %s (%s)\n",
				MutedStyle.Render(">"),
				DoneStyle.Render(completedAt.Format(domain.TimeLayouts[1])),
				MutedStyle.Render("Created at: "+createdAt),
			)
		} else {
			fmt.Printf(
				"  %s Due at: %s (%s)\n",
				MutedStyle.Render(">"),
				DueStyle.Render(dueAt.(string)),
				MutedStyle.Render("Created at: "+createdAt),
			)
		}

		if idx != tasksAmt-1 {
			fmt.Println()
		}
	}
}

/*
coloredOutputDeletedTask outputs the deleted task state with
git-style "<<" prefix and the entire message rendered red.
*/
func coloredOutputDeletedTask(id domain.ID) {
	msg := fmt.Sprintf("%s Task %s deleted successfully.", GreenCheckmark, fmt.Sprintf("%q", IDStyle.Render(id.String())))
	fmt.Println(RedOutput.Render(msg))
}

func coloredOutputUpdatedTask(id domain.ID) {
  fmt.Printf("%s Task %s updated successfully.\n", YellowArrows, fmt.Sprintf("%q", IDStyle.Render(id.String())))
}

func coloredOutputCompleteTask(id domain.ID) {
	fmt.Printf("%s Task %s successfully marked as complete\n", GreenCheckmark, fmt.Sprintf("%q", IDStyle.Render(id.String())))
}

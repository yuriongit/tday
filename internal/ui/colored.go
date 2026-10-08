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
	fmt.Printf("%s Task %s created\n", GreenCheckMark, IDStyle.Render(id.String()))
}

/*
coloredOutputAllTasks outputs all the persisted tasks in
a formatted manner.
*/
func coloredOutputAllTasks(tasks []domain.Task) {
	tasksAmt := len(tasks)

	fmt.Printf(
		"%s / %s / %s\n",
		BoldNum.Render(fmt.Sprintf("%d total", tasksAmt)),
		DueStyle.Render(fmt.Sprintf("%d remaining", tasksAmt)),
		DoneStyle.Render("0 complete"), // TODO: When tasks can be marked as complete
	)
	fmt.Println(Divider)

	for idx, task := range tasks {
		label, _ := (*task.InputData)[string(domain.LabelField)].(string)
		title, _ := (*task.InputData)[string(domain.TitleField)].(string)
		desc, _ := (*task.InputData)[string(domain.DescriptionField)].(string)
		dueAt, _ := (*task.InputData)[string(domain.DueAtField)].(string)
		completedAt := task.Metadata.CompletedAt
		createdAt := task.Metadata.CreatedAt.Format(domain.TimeLayouts[1])
		uuid := task.Metadata.UUID

		isCompleted := !completedAt.IsZero()

		// Choose styles based on completion state
		var renderedUUID, renderedLabel, renderedTitle string

		if isCompleted {
			renderedUUID = DoneStyle.Faint(true).Render("{" + uuid.String() + "}")
			renderedLabel = DoneStyle.Render("[" + label + "]:")
			renderedTitle = DoneStyle.Render(title)
		} else {
			renderedUUID = DueStyle.Faint(true).Render("{" + uuid.String() + "}")
			renderedLabel = DueStyle.Render("[" + label + "]:")
			renderedTitle = TitleStyle.Render(title)
		}

		fmt.Printf(
			"%s %s %s\n",
			renderedUUID,
			renderedLabel,
			renderedTitle,
		)

		// Time (Created at -> Completed at)
		if isCompleted {
			fmt.Printf(
				"%s %s %s %s\n",
				GreenCheckMark,
				DoneStyle.Italic(true).Faint(true).Render("(start)", createdAt),
				DoneStyle.Italic(true).Faint(true).Render(PointerSymbol),
				DoneStyle.Italic(true).Render(completedAt.Format(domain.TimeLayouts[1]), "(done)"),
			)
		} else {
			// Time (Created at -> Due at)
			fmt.Printf(
				"%s %s %s %s\n",
				DueStyle.Render(ProgressSymbol),
				DueStyle.Italic(true).Faint(true).Render("(start)", createdAt),
				DueStyle.Italic(true).Faint(true).Render(PointerSymbol),
				DueStyle.Italic(true).Render(dueAt, "(due)"),
			)
		}

		if desc != "" && !isCompleted {
			fmt.Printf("  %s %s\n", MutedStyle.Render(">"), DescStyle.Render(desc))
		} else if desc != "" {
			fmt.Printf("  %s %s\n", MutedStyle.Render(">"), DescStyle.Render(desc))
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
	msg := fmt.Sprintf("%s Task %s deleted", GreenCheckMark, IDStyle.Render(id.String()))
	fmt.Println(RedOutput.Render(msg))
}

func coloredOutputUpdatedTask(id domain.ID) {
	fmt.Println(Divider)
	fmt.Printf("%s Task %s updated\n", GreenCheckMark, IDStyle.Render(id.String()))
}

func coloredOutputCompleteTask(id domain.ID) {
	fmt.Printf("%s Task %s marked as complete\n", GreenCheckMark, IDStyle.Render(id.String()))
}

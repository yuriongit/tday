package task

import (
	"context"
	"fmt"

	"github.com/yuriongit/tday/internal/domain"
)

func ListAll(rootCtx context.Context, db domain.Database) error {
	tasks, err := db.QueryAllTasks(rootCtx)
	if err != nil {
		return err
	}

	outputAllTasks(tasks)

	return nil
}

func outputAllTasks(tasks []domain.Task) {
	if len(tasks) == 0 {
		fmt.Println("No tasks found")
		return
	}

	fmt.Printf(
  	"%d total tasks, %d remaining, %d complete\n", 
  	len(tasks), 
  	len(tasks),
  	0, // To be determined when tasks can be marked complete
	)
	fmt.Println("——————————————————————————————————————————")
	
	// • 
	
	for idx, task := range tasks {
		label := (*task.InputData)[domain.LabelField]
		title := (*task.InputData)[domain.TitleField]
		dueAt := (*task.InputData)[domain.DueAtField]
		createdAt := task.Metadata.CreatedAt.Format("Jan 2, 2006")
		time := task.Metadata.CreatedAt.Format(domain.TimeLayouts[1])
		fmt.Printf("• [%d] %s: %q\n", idx, label, title)
		fmt.Printf("   > Created: %s - %s\n   > Due: %v\n\n", createdAt, time, dueAt)
	}
}

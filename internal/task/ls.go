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
  	"1%d total / 1%d remaining / %d complete\n", 
  	len(tasks), 
  	len(tasks),
  	0, // To be determined when tasks can be marked complete
	)
	fmt.Println("—————————————————————————————————————")
	
	// • 
	
	for idx, task := range tasks {
		label := (*task.InputData)[domain.LabelField]
		title := (*task.InputData)[domain.TitleField]
		desc := (*task.InputData)[domain.DescriptionField]
		dueAt := (*task.InputData)[domain.DueAtField]
		time := task.Metadata.CreatedAt.Format(domain.TimeLayouts[1])
		fmt.Printf("• {%d} %s: %q\n", idx, label, title)
		if desc != "" {
  		fmt.Printf("   > Desc: %q\n", desc) 
		}
		fmt.Printf("   > Created at: %s | Due at: %s\n\n", time, dueAt)
	}
}

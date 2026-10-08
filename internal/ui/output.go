/*
Package ui provides the UI for TDay.
*/
package ui

import "github.com/yuriongit/tday/internal/domain"

// TaskOutput defines the structure for task output functions.
type TaskOutput struct {
	New      func(id *domain.ID)
	All      func(t []domain.Task)
	Update   func(id domain.ID)
	Complete func(id domain.ID)
	Remove   func(id domain.ID)
}

/*
UncoloredTaskOutput holds the functions responsible for any task-
related output (besides input collection).
*/
var UncoloredTaskOutput = TaskOutput{
	New:      uncoloredOutputNewTask,
	All:      uncoloredOutputAllTasks,
	Update:   uncoloredOutputUpdatedTask,
	Complete: uncoloredOutputCompleteTask,
	// TODO: Add AllCompleted
	Remove: uncoloredOutputDeletedTask,
}

/*
ColoredTaskOutput holds the functions responsible for any task-
related output (besides input collection).
*/
var ColoredTaskOutput = TaskOutput{
	New:      coloredOutputNewTask,
	All:      coloredOutputAllTasks,
	Update:   coloredOutputUpdatedTask,
	Complete: coloredOutputCompleteTask,
	Remove:   coloredOutputDeletedTask,
}

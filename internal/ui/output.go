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
NonColoredTaskOutput holds the functions responsible for any task-
related output (besides input collection).
*/
var NonColoredTaskOutput = TaskOutput{
	New:      noColorOutputNewTask,
	All:      noColorOutputAllTasks,
	Update:   noColorOutputUpdatedTask,
	Complete: noColorOutputCompleteTask,
	// TODO: Add AllCompleted
	Remove:   noColorOutputDeletedTask,
}

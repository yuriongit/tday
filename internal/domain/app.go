/*
Package domain contains the app's core
types and dependencies.
*/
package domain

import (
	"context"
)

/*
App contains the dependencies used
by the application.
*/
type App struct {
	TaskInputHandler *TaskInputHandler
	TaskIDGenerator  *TaskIDGenerator
	Database         *SupabaseDB
	Ctx              context.Context
	Cancel           context.CancelFunc
}

/*
InitApp creates and initializes
the application's dependencies.
*/
func InitApp() (*App, error) {
	db, err := NewSupabaseDB(context.Background())
	if err != nil {
		return nil, err
	}

	rootCtx, rootCancel := context.WithCancel(context.Background())

	return &App{
		TaskInputHandler: NewTaskInputHandler(),
		TaskIDGenerator:  &TaskIDGenerator{},
		Database:         db,
		Ctx:              rootCtx,
		Cancel:           rootCancel,
	}, nil
}

/*
Package domain contains the app's core
types and dependencies.
*/
package domain

import (
	"context"
	"time"
)

/*
App contains the dependencies used
by the application.
*/
type App struct {
	TaskInputHandler *TaskInputHandler
	TaskIDGenerator  *TaskIDGenerator
	Database         *SupabaseDB
}

/*
InitApp creates and initializes
the application's dependencies.
*/
func InitApp() (*App, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)

	db, err := NewSupabaseDB(ctx, cancel)
	if err != nil {
		return nil, err
	}

	return &App{
		TaskInputHandler: NewTaskInputHandler(),
		TaskIDGenerator:  &TaskIDGenerator{},
		Database:         db,
	}, nil
}

/*
Package app contains the app's core
types and dependencies.
*/
package app

/*
App contains the dependencies used
by the application.
*/
type App struct {
	TaskInputHandler *TaskInputHandler
	TaskIDGenerator  *TaskIDGenerator
}

/*
InitApp creates and initializes
the application dependencies.
*/
func InitApp() *App {
	return &App{
		TaskInputHandler: NewInputHandler(),
		TaskIDGenerator:  &TaskIDGenerator{},
	}
}

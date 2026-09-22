package app

type App struct {
	TaskInputHandler *TaskInputHandler
	TaskIDGenerator  *TaskIDGenerator
}

func InitApp() *App {
	return &App{
		TaskInputHandler: NewInputHandler(),
		TaskIDGenerator:  &TaskIDGenerator{},
	}
}

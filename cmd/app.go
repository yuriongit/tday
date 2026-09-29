package cmd

import app "github.com/yuriongit/tday/internal/app"

var application *app.App

func SetApp(a *app.App) {
	application = a
}

func GetApp() *app.App {
	return application
}

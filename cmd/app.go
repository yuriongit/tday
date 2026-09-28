package cmd

import "github.com/yuriongit/tday/internal/domain"

var app *domain.App

func SetApp(a *domain.App) {
	app = a
}

func GetApp() *domain.App {
	return app
}

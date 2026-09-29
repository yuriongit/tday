/*
Package cmd holds all of TDay's commands.

Copyright © 2026 Yuri Okeren <yuri.dev44@outlook.com>.
*/
package cmd

import app "github.com/yuriongit/tday/internal/app"

var application *app.App

func SetApp(a *app.App) {
	application = a
}

func GetApp() *app.App {
	return application
}

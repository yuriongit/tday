/*
Package ui provides the UI for TDay.
*/
package ui

import (
	"os"

	lg "charm.land/lipgloss/v2"
)

// Detect terminal background theme
var (
	HasDarkBG = lg.HasDarkBackground(os.Stdin, os.Stderr)
	lightDark = lg.LightDark(HasDarkBG)
)

// Color definitions supporting dark and light themes dynamically
var (
	blueColor   = lightDark(lg.Color("#3D99FF"), lg.Color("#3D99FF"))
	orangeColor = lightDark(lg.Color("#FF625C"), lg.Color("#FFA39E"))
	greenColor  = lightDark(lg.Color("#00A305"), lg.Color("42"))
	redColor    = lightDark(lg.Color("160"), lg.Color("#FF5B52"))

	labelColor = lightDark(lg.Color("#FF6BB1"), lg.Color("#FF6BB1"))
	descColor  = lightDark(lg.Color(""), lg.Color("#C2C2C2"))
	mutedColor = lightDark(lg.Color("#9E9E9E"), lg.Color("252"))
	titleColor = lightDark(lg.Color("#0D0D0D"), lg.Color("#FFF"))
)

// Lipgloss styles and rendered components
var (
	// Status indicators
	BlueArrow  = lg.NewStyle().Foreground(blueColor).Bold(true).Render(">")
	BlueArrows = lg.NewStyle().Foreground(blueColor).Bold(true).Render(">>")

	// Error indicator
	ErrStyle = lg.NewStyle().Italic(true).Foreground(redColor).Bold(true)

	// Labels & Metadata
	NewTaskFieldLabelStyle = lg.NewStyle().Foreground(blueColor).Bold(true)
	FieldLabelStyle        = lg.NewStyle().Foreground(labelColor).Bold(true)
	OptLabelStyle          = lg.NewStyle().Foreground(mutedColor).Italic(true)

	// ID
	IDStyle = NewTaskFieldLabelStyle.Italic(true)

	// Before / Diff output
	RedOutput = lg.NewStyle().Foreground(redColor)

	// Marker symbols
	GreenCheckMark = lg.NewStyle().Foreground(greenColor).Bold(true).Render(checkMark)
	RedXSymbol     = lg.NewStyle().Foreground(redColor).Bold(true).Render(xSymbol)

	// Task list styles
	TitleStyle = lg.NewStyle().Foreground(titleColor).Bold(true)
	MutedStyle = lg.NewStyle().Foreground(mutedColor)
	DescStyle  = lg.NewStyle().Foreground(descColor)
	DueStyle   = lg.NewStyle().Bold(true).Foreground(orangeColor)
	DoneStyle  = lg.NewStyle().Bold(true).Foreground(greenColor)
	BoldNum    = lg.NewStyle().Foreground(titleColor).Bold(true)
)

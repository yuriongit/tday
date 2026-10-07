/*
Package ui provides the UI for TDay.
*/
package ui

import (
	lg "charm.land/lipgloss/v2"
)

// Lipgloss styles
var (
	// Status indicators
	BlueArrows   = lg.NewStyle().Foreground(lg.Blue).Bold(true).Render(">>")
	YellowArrows = lg.NewStyle().Foreground(lg.Color("214")).Bold(true).Render("~>")
	GreenArrows  = lg.NewStyle().Foreground(lg.Color("42")).Bold(true).Render(">>")

	// Error indicator
	ErrStyle = lg.NewStyle().Foreground(lg.Color("196")).Bold(true)
	
	// Labels & Metadata
	NewTaskFieldLabelStyle = lg.NewStyle().Foreground(lg.Blue).Bold(true)
	UpdTaskFieldLabelStyle = lg.NewStyle().Foreground(lg.Color("42")).Bold(true)
	FieldLabelStyle = lg.NewStyle().Foreground(lg.Color("212")).Bold(true)
	OptLabelStyle   = lg.NewStyle().Foreground(lg.Color("243")).Italic(true)
	
	// ID
	IDStyle = NewTaskFieldLabelStyle.Italic(true)

	// Before output
	RedOutput = lg.NewStyle().Foreground(lg.Color("196"))

	// Marker symbols
	GreenCheckmark = lg.NewStyle().Foreground(lg.Color("42")).Bold(true).Render(checkmark)
	RedErrormark   = lg.NewStyle().Foreground(lg.Color("196")).Bold(true).Render(errormark)

	// Task list styles
	UUIDStyle  = lg.NewStyle().Foreground(lg.Color("63"))
	LabelStyle = lg.NewStyle().Foreground(lg.Color("212")).Bold(true)
	TitleStyle = lg.NewStyle().Bold(true)
	MutedStyle = lg.NewStyle().Foreground(lg.Color("243"))
	DescStyle  = lg.NewStyle().Foreground(lg.Color("252"))
	DueStyle   = lg.NewStyle().Foreground(lg.Color("214"))
	DoneStyle  = lg.NewStyle().Foreground(lg.Color("42"))
	BoldNum    = lg.NewStyle().Bold(true)
)

/*
Package ui provides the UI for TDay.
*/
package ui

import "charm.land/lipgloss/v2"

// Unicode symbols
const (
  checkMark = "✓"
  _divider = "—————————————————————————————————————"
  xSymbol = "✗"
  ProgressSymbol = "◌"
  PointerSymbol = "-—+>"
)

// Styled components
var (
  Divider = lipgloss.NewStyle().Faint(true).Render(_divider)
)
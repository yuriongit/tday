/*
Package domain is responsible for business
logic, pure data representation, and
constants.
*/
package domain

/*
Indent is the amount of spaces for an
indent constant.
*/
var Indent = "  "

// QuoteSymbol for wrapping quoted text.
var QuoteSymbol = `"`

/*
TimeLayouts defines the tolerated time
layouts for displaying time and accepting
a value for a task's "due_at" field.
*/
var TimeLayouts = []string{
	"3PM",
	"3:04PM",
}

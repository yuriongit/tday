/*
Package domain is responsible for business
logic, pure data representation, and
constants.
*/
package domain

/*
Indent is the amount of spaces for an indent
presented in any output.
*/
var Indent = "  "

/*
TimeLayouts defines the tolerated time
layouts for displaying time and accepting
a value for a task's "due_at" field.
*/
var TimeLayouts = []string{
	"3pm",
	"3:04pm",
}

/*
DBConnStringVarName is the name of the connection
string environment variable for the database.
*/
var DBConnStringVarName = "POSTGRES_DIRECT_URI"

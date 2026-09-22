/*
Package app contains the app's core
types and dependencies.
*/
package app

import (
	"uuid"
)

/*
ID is branded type which is of type string,
serving the purpose of being a uniquely
identified type for fields.
*/
type ID string

/*
String is the method for casting an ID to
a string.
*/
func (id *ID) String() string {
	return string(*id)
}

/*
IDGenerator is a generic interface for ID
generators; Includes a Generate method for
IDGenerators.
*/
type IDGenerator interface {
	Generate() ID
}

/*
TaskIDGenerator is an IDGenerator for
tasks.
*/
type TaskIDGenerator struct{}

/*
Generate implements IDGenerator's Generate
method for TaskIDGenerator.
*/
func (*TaskIDGenerator) Generate() ID {
	return ID(uuid.New().String()[:5])
}

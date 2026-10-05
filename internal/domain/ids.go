/*
Package domain contains the domain's core
types and dependencies.
*/
package domain

/*
ID is branded type which is of type string,
serving the purpose of being a uniquely
identified type for fields.
*/
type ID string

// TaskIDLen is the length of generated UUIDs.
var TaskIDLen = 5

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

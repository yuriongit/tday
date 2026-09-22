package app

import (
	"uuid"
)

// ID is a branded string.
type ID string

// String casts ID to a string.
func (id *ID) String() string {
	return string(*id)
}

/*
IDGenerator is a generic interface for ID generators;
Includes a Generate method for IDGenerators.
*/
type IDGenerator interface {
	Generate() ID
}

// TaskIDGenerator is an IDGenerator for Tasks.
type TaskIDGenerator struct{}

// Implements IDGenerator's Generate method for TaskIDGenerator.
func (_ *TaskIDGenerator) Generate() ID {
	return ID(uuid.New().String()[:5])
}

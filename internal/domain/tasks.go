/*
Package domain is responsible for business
logic, pure data representation, and
constants.
*/
package domain

import (
	"fmt"
	"time"
)

/*
TaskFieldType identifies the type of
field a task can have.
*/
type TaskFieldType string

const (
	/*
	  LabelField identifies the label field	used
	  to categorize a task.
	*/
	LabelField TaskFieldType = "label"

	/*
	  TitleField identifies the title field
	 	used to name a task.
	*/
	TitleField TaskFieldType = "title"

	/*
	  DescriptionField identifies the
	  description field used to provide task
	  details.
	*/
	DescriptionField TaskFieldType = "description"

	/*
		DueAtField identifies the due date field
		used to set when a task is due.
	*/
	DueAtField TaskFieldType = "due_at"
)

/*
FieldDefinition describes a field and how
its value should be validated.
*/
type FieldDefinition struct {
	ID       TaskFieldType
	Name     string
	Type     string
	Required bool
	Validate func(value string, isUpdate bool) error
}

/*
AllFields contains the fields that make up
a task.
*/
var AllFields = []FieldDefinition{
	{
		ID:       LabelField,
		Name:     "Label",
		Type:     "string",
		Required: true,
		Validate: validateLabel,
	},
	{
		ID:       TitleField,
		Name:     "Title",
		Type:     "string",
		Required: true,
		Validate: validateTitle,
	},
	{
		ID:       DescriptionField,
		Name:     "Description",
		Type:     "string",
		Required: false,
		Validate: validateDescription,
	},
	{
		ID:       DueAtField,
		Name:     "Due at",
		Type:     "uint8",
		Required: true,
		Validate: validateDueAt,
	},
}

/*
TaskMetadata contains information
*/
type TaskMetadata struct {
	UUID        ID
	CreatedAt   time.Time
	CompletedAt time.Time
}

/*
TaskInputData contains the values provided
for a task's fields.
*/
// type TaskInputData map[TaskFieldType]any
type TaskInputData map[string]any

/*
Task represents a task and its
associated data.
*/
type Task struct {
	Metadata  TaskMetadata
	InputData *TaskInputData
}

/*
GetFieldDef returns the definition
for a field by its ID.
*/
func GetFieldDef(id TaskFieldType) *FieldDefinition {
	for i := range AllFields {
		if AllFields[i].ID == id {
			return &AllFields[i]
		}
	}
	return nil
}

func validateLabel(value string, update bool) error {
	if update && value == "" {
		return nil
	}
	if value == "" {
		return fmt.Errorf("Label cannot be empty")
	}
	if len(value) > 10 {
		return fmt.Errorf("Label cannot exceed 10 characters")
	}
	return nil
}

func validateTitle(value string, update bool) error {
	if update && value == "" {
		return nil
	}
	if value == "" {
		return fmt.Errorf("Title cannot be empty")
	}
	if len(value) > 100 {
		return fmt.Errorf("Title cannot exceed 100 characters")
	}
	return nil
}

func validateDescription(value string, update bool) error {
	if update && value == "" {
		return nil
	}
	if len(value) > 500 {
		return fmt.Errorf("Description cannot exceed 500 characters")
	}
	return nil
}

func validateDueAt(value string, update bool) error {
	if update && value == "" {
		return nil
	}
	if value == "" {
		return fmt.Errorf("Due at cannot be empty")
	}

	for _, layout := range TimeLayouts {
		if _, err := time.Parse(layout, value); err == nil {
			return nil
		}
	}

	return fmt.Errorf("Invalid format (e.g. %q or %q)", TimeLayouts[0], TimeLayouts[1])
}

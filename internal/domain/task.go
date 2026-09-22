/*
Package domain is responsible for business
logic, pure data representation, and
constants.
*/
package domain

import (
	"fmt"
	"time"

	"github.com/yuriongit/tday/internal/app"
)

/*
ID represents a unique identifier used by
the domain.
*/
type ID string

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
FieldDefinition describes a field and
how its value should be validated.
*/
type FieldDefinition struct {
	ID       TaskFieldType
	Name     string
	Type     string
	Required bool
	Validate func(string) error
}

/*
AllFields contains the fields that can
be used when creating a task.
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
Metadata contains information shared
across domain objects.
*/
type Metadata struct {
	UUID      app.ID
	CreatedAt time.Time
}

/*
TaskInputData contains the values
provided for a task's fields.
*/
type TaskInputData map[TaskFieldType]any

/*
Task represents a task and its
associated data.
*/
type Task struct {
	Metadata
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

func validateLabel(s string) error {
	if len(s) == 0 {
		return fmt.Errorf("label cannot be empty")
	}
	if len(s) > 50 {
		return fmt.Errorf("label cannot exceed 50 characters")
	}
	return nil
}

func validateTitle(s string) error {
	if len(s) == 0 {
		return fmt.Errorf("title cannot be empty")
	}
	return nil
}

func validateDescription(s string) error {
	if len(s) > 500 {
		return fmt.Errorf("description cannot exceed 500 characters")
	}
	return nil
}

func validateDueAt(s string) error {
	for _, layout := range TimeLayouts {
		if _, err := time.Parse(layout, s); err == nil {
			return nil
		}
	}

	return fmt.Errorf(
		"invalid time format. \nFormats include: '%s' and '%s'",
		TimeLayouts[0],
		TimeLayouts[1],
	)
}

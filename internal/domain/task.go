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

type ID string

type TaskFieldType string

const (
  LabelField TaskFieldType = "label"
  TitleField TaskFieldType = "title"
  DescriptionField TaskFieldType = "description"
  DueAtField TaskFieldType = "due_at"
)

type FieldDefinition struct {
  ID  TaskFieldType
  Name  string
  Type string
  Required bool
  Validate func(string) (error)
}
var AllFields = []FieldDefinition{
  {
    ID: LabelField,
    Name: "Label",
    Type: "string",
    Required: true,
    Validate: validateLabel,
  },
  {
    ID: TitleField,
    Name: "Title",
    Type: "string",
    Required: true,
    Validate: validateTitle,
  },
  {
    ID: DescriptionField,
    Name: "Description",
    Type: "string",
    Required: false,
    Validate: validateDescription,
  },
  {
    ID: DueAtField,
    Name: "Due at",
    Type: "uint8",
    Required: true,
    Validate: validateDueAt,
  },
}

type Metadata struct {
  UUID app.ID 
  CreatedAt time.Time
}

type TaskInputData map[TaskFieldType]any

type Task struct {
  Metadata
  InputData *TaskInputData
}

// Helper to get a field def by ID
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

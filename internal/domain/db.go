package domain

// Column describes a database column used by the tasks table.
//
// Name is the PostgreSQL column name.
// Type describes the expected application-level type. It is currently
// informational and can later be used for validation or conversion.
type Column struct {
	Name string
	Type string
}

// TasksModel describes the structure of the tasks table.
type TasksModel struct {
	Columns []Column
}

/*
TasksSchema is the single source of truth for the tasks table columns.

The first three columns are database-managed metadata:

	uuid
	created_at
	completed_at

The remaining columns are populated from Task.InputData.
*/
var TasksSchema = TasksModel{
	Columns: []Column{
		{Name: "uuid", Type: "ID"},
		{Name: "created_at", Type: "time.Time"},
		{Name: "completed_at", Type: "time.Time"},
		{Name: "label", Type: "string"},
		{Name: "title", Type: "string"},
		{Name: "description", Type: "string"},
		{Name: "due_at", Type: "string"},
	},
}

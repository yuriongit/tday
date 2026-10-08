/*
Package app contains the application's core types
and dependencies.
*/
package app

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/yuriongit/tday/internal/cnf"
	"github.com/yuriongit/tday/internal/domain"
	"github.com/yuriongit/tday/internal/ui"
)

const (
	// metadataColumnCount is the number of columns managed by the application
	// rather than by Task.InputData.
	metadataColumnCount = 3

	// databaseOperationTimeout is the maximum time allowed for a database
	// operation before its context is cancelled.
	databaseOperationTimeout = 5 * time.Second
)

/*
Database defines connectivity and managerial
methods for tasks.
*/
type Database interface {
	// Connectivity methods.
	Ping(rootCtx context.Context) error

	// CRUD methods.
	InsertTask(rootCtx context.Context, task *domain.Task) error
	QueryTask(rootCtx context.Context, id domain.ID) (*domain.Task, error)
	QueryAllTasks(rootCtx context.Context) ([]domain.Task, error)
	UpdateTask(rootCtx context.Context, updatedTaskFields *domain.TaskInputData, id domain.ID) error
	DeleteTask(rootCtx context.Context, id domain.ID) error

	// Remaining CRUD methods.
	/*
		DeleteSetOfTask()
		CompleteTask()
	*/
}

// SupabaseDB implements the Database interface.
type SupabaseDB struct {
	Pool *pgxpool.Pool
}

/*
NewSupabaseDB instantiates a new SupabaseDB
struct and opens a PostgreSQL connection pool.
*/
func NewSupabaseDB(tempCtx context.Context) (*SupabaseDB, error) {
	db := &SupabaseDB{}

	if err := db.newPool(tempCtx); err != nil {
		return nil, err
	}

	return db, nil
}

// --------------------
// Database helpers
// --------------------

/*
quoteIdentifier quotes a PostgreSQL identifier.

Column names cannot be supplied as normal PostgreSQL parameters such as $1.
They must be part of the SQL string itself. These identifiers are safe here
because they come from the trusted, compiled-in TasksSchema definition.

The escaping also protects against accidentally including a double quote in
a future schema name.
*/
func quoteIdentifier(identifier string) string {
	return `"` + strings.ReplaceAll(identifier, `"`, `""`) + `"`
}

/*
schemaColumns returns all task table columns in the order declared by
domain.TasksSchema.

The returned names are quoted PostgreSQL identifiers.
*/
func schemaColumns() []string {
	columns := make([]string, 0, len(domain.TasksSchema.Columns))

	for _, column := range domain.TasksSchema.Columns {
		columns = append(columns, quoteIdentifier(column.Name))
	}

	return columns
}

/*
inputSchemaColumns returns the task-input columns.

The first three columns in TasksSchema are metadata columns:

	uuid
	created_at
	completed_at

All remaining columns are read from or written to Task.InputData.
*/
func inputSchemaColumns() []domain.Column {
	if len(domain.TasksSchema.Columns) <= metadataColumnCount {
		return nil
	}

	return domain.TasksSchema.Columns[metadataColumnCount:]
}

/*
validateTasksSchema verifies the assumptions used by the database methods.

The database code relies on the first three schema columns being the metadata
columns. Failing early here produces a clearer error than allowing INSERT and
SELECT values to become misaligned.
*/
func validateTasksSchema() error {
	if len(domain.TasksSchema.Columns) < metadataColumnCount {
		return fmt.Errorf(
			"Tasks schema must contain at least %d columns",
			metadataColumnCount,
		)
	}

	expectedMetadataColumns := []string{
		"uuid",
		"created_at",
		"completed_at",
	}

	for index, expectedName := range expectedMetadataColumns {
		actualName := domain.TasksSchema.Columns[index].Name

		if actualName != expectedName {
			return fmt.Errorf(
				"Invalid tasks schema: column %d must be %q, got %q",
				index,
				expectedName,
				actualName,
			)
		}
	}

	for index, column := range domain.TasksSchema.Columns {
		if strings.TrimSpace(column.Name) == "" {
			return fmt.Errorf(
				"Invalid tasks schema: column %d has an empty name",
				index,
			)
		}
	}

	return nil
}

// --------------------
// Connectivity
// --------------------

// newPool creates the PostgreSQL connection pool.
func (db *SupabaseDB) newPool(tempCtx context.Context) error {
	if err := validateTasksSchema(); err != nil {
		return err
	}

	// Change into the application's configuration directory.
	if err := cnf.ChdirToConfigDir(); err != nil {
		return err
	}

	// Load environment variables from the .env file.
	if err := godotenv.Load(".env"); err != nil {
		return fmt.Errorf("Environment variables load error: %w", err)
	}

	// Retrieve the database connection string.
	dbURI := os.Getenv(domain.DBConnStringVarName)
	if dbURI == "" {
		return fmt.Errorf("Missing %q", domain.DBConnStringVarName)
	}

	// Parse the PostgreSQL connection configuration.
	poolConfig, err := pgxpool.ParseConfig(dbURI)
	if err != nil {
		return fmt.Errorf("Configuration error: %w", err)
	}

	// Use the simple protocol instead of pgx's prepared-statement cache.
	// This is appropriate for a short-lived CLI application.
	poolConfig.ConnConfig.DefaultQueryExecMode = pgx.QueryExecModeSimpleProtocol

	// Keep the pool small because this is a sequential CLI application.
	poolConfig.MaxConns = 5
	poolConfig.MinConns = 0
	poolConfig.MaxConnIdleTime = 10 * time.Second

	// Use a bounded context while opening and validating the pool.
	ctx, cancel := context.WithTimeout(
		tempCtx,
		databaseOperationTimeout,
	)
	defer cancel()

	// Create the connection pool.
	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		return fmt.Errorf("PostgreSQL connection error: %w", err)
	}

	// Store the pool on the database object.
	db.Pool = pool

	// Verify that the database is reachable.
	if err := db.Ping(ctx); err != nil {
		pool.Close()
		db.Pool = nil

		return fmt.Errorf("Database ping error: %w", err)
	}

	return nil
}

// Ping verifies that the database is reachable.
func (db *SupabaseDB) Ping(rootCtx context.Context) error {
	ctx, cancel := context.WithTimeout(
		rootCtx,
		databaseOperationTimeout,
	)
	defer cancel()

	if db.Pool == nil {
		return fmt.Errorf("Database pool is nil")
	}

	if err := db.Pool.Ping(ctx); err != nil {
		return err
	}

	return nil
}

// --------------------
// CRUD operations
// --------------------

/*
InsertTask inserts a new task into the database.

The SQL column list is generated from domain.TasksSchema. Runtime values are
passed separately as PostgreSQL parameters, which prevents SQL injection.
*/
func (db *SupabaseDB) InsertTask(
	rootCtx context.Context,
	task *domain.Task,
) error {
	ctx, cancel := context.WithTimeout(
		rootCtx,
		databaseOperationTimeout,
	)
	defer cancel()

	if task == nil {
		return fmt.Errorf("Cannot insert a nil task")
	}

	if task.InputData == nil {
		return fmt.Errorf("Cannot insert a task with nil input data")
	}

	query, args, err := buildInsertTaskQuery(task)
	if err != nil {
		return err
	}

	cmdTag, err := db.Pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("Insert execution error: %w", err)
	}

	if cmdTag.RowsAffected() != 1 {
		return fmt.Errorf(
			"Insert failed: expected 1 affected row, got %d",
			cmdTag.RowsAffected(),
		)
	}

	return nil
}

/*
buildInsertTaskQuery builds the INSERT statement and its arguments.

The columns come from the trusted TasksSchema definition. The values are
passed as PostgreSQL parameters using $1, $2, and so on.
*/
func buildInsertTaskQuery(task *domain.Task) (string, []any, error) {
	if err := validateTasksSchema(); err != nil {
		return "", nil, err
	}

	inputData := *task.InputData

	// All schema columns are inserted in schema order.
	columns := schemaColumns()

	// The first three values are database metadata.
	args := []any{
		task.Metadata.UUID,
		task.Metadata.CreatedAt,
		task.Metadata.CompletedAt,
	}

	// The remaining values come from Task.InputData.
	for _, column := range inputSchemaColumns() {
		value, ok := inputData[column.Name]
		if !ok {
			return "", nil, fmt.Errorf(
				"Missing task input value for column %q",
				column.Name,
			)
		}

		args = append(args, value)
	}

	// Generate $1, $2, $3, etc. for every argument.
	placeholders := make([]string, len(args))

	for index := range args {
		placeholders[index] = fmt.Sprintf("$%d", index+1)
	}

	query := fmt.Sprintf(
		`INSERT INTO "tasks" (%s) VALUES (%s)`,
		strings.Join(columns, ", "),
		strings.Join(placeholders, ", "),
	)

	return query, args, nil
}

// QueryTask queries a tasks from the database.
func (db *SupabaseDB) QueryTask(
	rootCtx context.Context,
	id domain.ID,
) (*domain.Task, error) {
	ctx, cancel := context.WithTimeout(
		rootCtx,
		databaseOperationTimeout,
	)
	defer cancel()

	columns := schemaColumns()

	query := fmt.Sprintf(
		`SELECT %s
		 FROM "tasks"
		 WHERE uuid = $1`,
		strings.Join(columns, ", "),
	)

	rows, err := db.Pool.Query(ctx, query, id)
	if err != nil {
		return nil, fmt.Errorf(
			"Failed to query task: %w",
			err,
		)
	}
	defer rows.Close()

	tasks := make([]domain.Task, 0)

	for rows.Next() {
		var task domain.Task

		// Create one scan destination for each schema column.
		scanArgs := make([]any, len(domain.TasksSchema.Columns))

		// Scan the three metadata columns into their strongly typed fields.
		scanArgs[0] = &task.Metadata.UUID
		scanArgs[1] = &task.Metadata.CreatedAt
		scanArgs[2] = &task.Metadata.CompletedAt

		// Scan dynamic task-input columns into temporary values.
		inputValues := make([]any, len(inputSchemaColumns()))

		for index := range inputValues {
			scanArgs[index+metadataColumnCount] = &inputValues[index]
		}

		if err := rows.Scan(scanArgs...); err != nil {
			return nil, fmt.Errorf(
				"Failed to scan task: %w",
				err,
			)
		}

		// Rebuild Task.InputData using the schema column names.
		inputData := make(domain.TaskInputData, len(inputValues))

		for index, column := range inputSchemaColumns() {
			inputData[column.Name] = inputValues[index]
		}

		task.InputData = &inputData
		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf(
			"Error iterating tasks: %w",
			err,
		)
	}

	if len(tasks) == 0 {
		return nil, fmt.Errorf(
			"Task %s %s",
			ui.IDStyle.Render(id.String()),
			ui.ErrStyle.Render("does not exist"),
		)
	}

	return &tasks[0], nil
}

/*
QueryAllTasks queries all tasks from the database.

The SELECT column list and scan order are both derived from TasksSchema.
Because pgx scans values positionally, keeping these orders identical is
important.
*/
func (db *SupabaseDB) QueryAllTasks(
	rootCtx context.Context,
) ([]domain.Task, error) {
	ctx, cancel := context.WithTimeout(
		rootCtx,
		databaseOperationTimeout,
	)
	defer cancel()

	columns := schemaColumns()

	// completed_at is the third column in TasksSchema.
	orderColumn := quoteIdentifier(
		domain.TasksSchema.Columns[2].Name,
	)

	query := fmt.Sprintf(
		`SELECT %s
		 FROM "tasks"
		 ORDER BY %s ASC`,
		strings.Join(columns, ", "),
		orderColumn,
	)

	rows, err := db.Pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf(
			"Failed to query all tasks: %w",
			err,
		)
	}
	defer rows.Close()

	tasks := make([]domain.Task, 0)

	for rows.Next() {
		var task domain.Task

		// Create one scan destination for each schema column.
		scanArgs := make([]any, len(domain.TasksSchema.Columns))

		// Scan the three metadata columns into their strongly typed fields.
		scanArgs[0] = &task.Metadata.UUID
		scanArgs[1] = &task.Metadata.CreatedAt
		scanArgs[2] = &task.Metadata.CompletedAt

		// Scan dynamic task-input columns into temporary values.
		inputValues := make([]any, len(inputSchemaColumns()))

		for index := range inputValues {
			scanArgs[index+metadataColumnCount] = &inputValues[index]
		}

		if err := rows.Scan(scanArgs...); err != nil {
			return nil, fmt.Errorf(
				"Failed to scan task: %w",
				err,
			)
		}

		// Rebuild Task.InputData using the schema column names.
		inputData := make(domain.TaskInputData, len(inputValues))

		for index, column := range inputSchemaColumns() {
			inputData[column.Name] = inputValues[index]
		}

		task.InputData = &inputData
		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf(
			"Error iterating tasks: %w",
			err,
		)
	}

	if len(tasks) == 0 {
		return nil, fmt.Errorf("No tasks found")
	}

	return tasks, nil
}

/*
DeleteTask deletes a task from the database using its UUID.
*/
func (db *SupabaseDB) DeleteTask(
	rootCtx context.Context,
	id domain.ID,
) error {
	ctx, cancel := context.WithTimeout(
		rootCtx,
		databaseOperationTimeout,
	)
	defer cancel()

	cmdTag, err := db.Pool.Exec(
		ctx,
		`DELETE FROM "tasks" WHERE "uuid" = $1`,
		id,
	)
	if err != nil {
		return fmt.Errorf(
			"Failed to delete task: %w",
			err,
		)
	}

	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf(
			"Task %s %s",
			ui.IDStyle.Render(id.String()),
			ui.ErrStyle.Render("does not exist"),
		)
	}

	return nil
}

/*
buildUpdateTaskQuery dynamically constructs the UPDATE statement and argument slice.

The column identifiers are safely quoted, and values are passed sequentially
as $2, $3, etc. Parameter $1 is always reserved for the task UUID in the WHERE clause.
*/
func buildUpdateTaskQuery(
	updatedTaskFields *domain.TaskInputData,
	id domain.ID,
) (string, []any, error) {
	if updatedTaskFields == nil || len(*updatedTaskFields) == 0 {
		return "", nil, fmt.Errorf("no fields provided to update")
	}

	setClauses := make([]string, 0, len(*updatedTaskFields))
	args := make([]any, 0, len(*updatedTaskFields)+1)

	// $1 is reserved for the task ID in the WHERE clause.
	args = append(args, id)

	paramIdx := 2
	for key, val := range *updatedTaskFields {
		setClauses = append(
			setClauses,
			fmt.Sprintf("%s = $%d", quoteIdentifier(key), paramIdx),
		)
		args = append(args, val)
		paramIdx++
	}

	query := fmt.Sprintf(
		`UPDATE "tasks" SET %s WHERE "uuid" = $1`,
		strings.Join(setClauses, ", "),
	)

	return query, args, nil
}

/*
UpdateTask modifies an existing task in the database using its UUID.

It utilizes buildUpdateTaskQuery to construct the dynamic SQL statement
and parameters, executes the update within a timeout context, and verifies
that the target task exists.
*/
func (db *SupabaseDB) UpdateTask(
	rootCtx context.Context,
	updatedTaskFields *domain.TaskInputData,
	id domain.ID,
) error {
	ctx, cancel := context.WithTimeout(
		rootCtx,
		databaseOperationTimeout,
	)
	defer cancel()

	query, args, err := buildUpdateTaskQuery(updatedTaskFields, id)
	if err != nil {
		return err
	}

	cmdTag, err := db.Pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("Failed to update task: %w", err)
	}

	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf(
			"Task %s %s",
			ui.IDStyle.Render(id.String()),
			ui.ErrStyle.Render("does not exist"),
		)
	}

	return nil
}

/*
buildCompleteTaskQuery constructs the UPDATE statement and argument slice for completing a task.

It sets the completed_at timestamp to the current time, reserving $1 for the task UUID in the WHERE clause.
*/
func buildCompleteTaskQuery(id domain.ID) (string, []any) {
	args := []any{id, time.Now()}

	query := fmt.Sprintf(
		`UPDATE "tasks" SET %s = $2 WHERE "uuid" = $1`,
		quoteIdentifier("completed_at"),
	)

	return query, args
}

/*
CompleteTask modifies an existing task in the database by setting its completion timestamp.

It utilizes buildCompleteTaskQuery to construct the SQL statement, executes the update within a timeout context, and
verifies that the target task exists.
*/
func (db *SupabaseDB) CompleteTask(
	rootCtx context.Context,
	id domain.ID,
) error {
	ctx, cancel := context.WithTimeout(
		rootCtx,
		databaseOperationTimeout,
	)
	defer cancel()

	query, args := buildCompleteTaskQuery(id)

	cmdTag, err := db.Pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("Failed to mark task as complete: %w", err)
	}

	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf(
			"Task %s %s",
			ui.IDStyle.Render(id.String()),
			ui.ErrStyle.Render("does not exist"),
		)
	}

	return nil
}

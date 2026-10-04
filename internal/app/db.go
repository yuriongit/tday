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

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/yuriongit/tday/internal/config"
	"github.com/yuriongit/tday/internal/domain"
)

/*
Database defines connectivity and managerial
methods for tasks.
*/
type Database interface {
	// Connectivity methods
	Ping(rootCtx context.Context) error

	// CRUD methods
	InsertTask(rootCtx context.Context, task *domain.Task) error
	QueryAllTasks(rootCtx context.Context) ([]domain.Task, error)
	DeleteTask(rootCtx context.Context, id domain.ID) error

	// Remaining CRUD methods
	/* QueryTask()
	DeleteTask()
	DeleteSetOfTask()
	UpdateTask()
	CompleteTask() */
}

// SupabaseDB implements the Database interface.
type SupabaseDB struct {
	Pool *pgxpool.Pool
}

/*
NewSupabaseDB instantiates a new SupabaseDB
struct.
*/
func NewSupabaseDB(tempCtx context.Context) (*SupabaseDB, error) {
	db := &SupabaseDB{}

	if err := db.newPool(tempCtx); err != nil {
		return nil, err
	}

	return db, nil
}

// ---------------

// Connectivity operations

// newPool creates the connection pool.
func (db *SupabaseDB) newPool(tempCtx context.Context) error {
	// Change into config directory
	if err := config.ChdirToConfigDir(); err != nil {
		return err
	}

	// Load environment variables from .env file.
	if err := godotenv.Load(".env"); err != nil {
		return fmt.Errorf("Environment variables load error: %w", err)
	}

	// Retrieves DB_URI environment variable.
	dbURI := os.Getenv(domain.DBConnStringVarName)
	if dbURI == "" {
		return fmt.Errorf("Missing '%s'", domain.DBConnStringVarName)
	}

	// Creates a Postgres config.
	config, err := pgxpool.ParseConfig(dbURI)
	if err != nil {
		return fmt.Errorf("Configuration Error: %w", err)
	}

	ctx, cancel := context.WithTimeout(tempCtx, 5*time.Second)
	defer cancel()

	// Creates a connection pool.
	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return fmt.Errorf("Supabase Connection Error: %w", err)
	}

	/*
		Assign the connection pool directly to the
		SupabaseDB struct.
	*/
	db.Pool = pool

	// Checks health of database via a ping.
	if err := db.Ping(ctx); err != nil {
		return err
	}

	return nil
}

// Ping is the health check for the database.
func (db *SupabaseDB) Ping(rootCtx context.Context) error {
	ctx, cancel := context.WithTimeout(rootCtx, 5*time.Second)
	defer cancel()

	if err := db.Pool.Ping(ctx); err != nil {
		return err
	}

	return nil
}

// Connectivity operations

// ---------------

// CRUD operations:

/*
InsertTask inserts a new task into the
database.
*/
func (db *SupabaseDB) InsertTask(rootCtx context.Context, task *domain.Task) error {
	ctx, cancel := context.WithTimeout(rootCtx, 5*time.Second)
	defer cancel()

	query, args := db.buildInsertTaskQuery(task)

	cmdTag, err := db.Pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("Insert Execution Error: %w", err)
	}

	// Check for insertion failure; if 0 rows were
	// affected
	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("Insert Failed: No rows affected")
	}

	return nil
}

func (db *SupabaseDB) buildInsertTaskQuery(task *domain.Task) (string, []any) {
	inputMap := *task.InputData

	columns := []string{"uuid", "created_at", "completed_at"}
	args := []any{task.Metadata.UUID, task.Metadata.CreatedAt, task.Metadata.CompletedAt}

	for _, field := range domain.AllFields {
		columns = append(columns, fmt.Sprintf(`"%s"`, field.ID))
		args = append(args, inputMap[field.ID])
	}

	placeholders := make([]string, len(args))
	for i := range args {
		placeholders[i] = fmt.Sprintf("$%d", i+1)
	}

	query := fmt.Sprintf(
		`INSERT INTO "tasks" (%s) VALUES (%s)`,
		strings.Join(columns, ", "),
		strings.Join(placeholders, ", "),
	)

	return query, args
}

// QueryAllTasks queries all tasks from the database.
func (db *SupabaseDB) QueryAllTasks(rootCtx context.Context) ([]domain.Task, error) {
	ctx, cancel := context.WithTimeout(rootCtx, 5*time.Second)
	defer cancel()

	// Build dynamic SELECT statement from AllFields
	columnNames := make([]string, 0, len(domain.AllFields))
	for _, field := range domain.AllFields {
		columnNames = append(columnNames, string(field.ID))
	}

	query := fmt.Sprintf("SELECT uuid, created_at, completed_at, %s FROM tasks ORDER BY completed_at ASC", strings.Join(columnNames, ", "))

	rows, err := db.Pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("Failed to query all tasks: %w", err)
	}
	defer rows.Close()

	var tasks []domain.Task

	for rows.Next() {
		var task domain.Task

		// Create scan args: UUID, CreatedAt, then one for each field
		scanArgs := make([]any, len(domain.AllFields)+3)
		scanArgs[0] = &task.Metadata.UUID
		scanArgs[1] = &task.Metadata.CreatedAt
		scanArgs[2] = &task.Metadata.CompletedAt

		values := make([]any, len(domain.AllFields))
		for i := range domain.AllFields {
			scanArgs[i+3] = &values[i]
		}

		if err := rows.Scan(scanArgs...); err != nil {
			return nil, fmt.Errorf("Failed to scan task: %w", err)
		}

		// Populate InputData map dynamically from AllFields
		inputData := domain.TaskInputData{}
		for i, field := range domain.AllFields {
			inputData[field.ID] = values[i]
		}
		task.InputData = &inputData

		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Error iterating rows: %w", err)
	}

	if len(tasks) == 0 {
		return nil, fmt.Errorf("No tasks found")
	}
	
	return tasks, nil
}

// DeleteTask deletes a task from the database.
func (db *SupabaseDB) DeleteTask(rootCtx context.Context, id domain.ID) error {
  if len(id) != domain.IDLen {
    return fmt.Errorf("Invalid task ID; Task ID must be 5 characters.")
  }
  
	cmd, err := db.Pool.Exec(rootCtx, "DELETE FROM tasks WHERE uuid = $1", id)
	
	if cmd.RowsAffected() == 0 {
		return fmt.Errorf("Task %q does not exist.", id)
	}
	
	if err != nil {
		return fmt.Errorf("Failed to delete task: %w", err)
	}
	return nil
}

// ---------------

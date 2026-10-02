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
		return fmt.Errorf("env load error: %w", err)
	}

	// Retrieves DB_URI environment variable.
	dbURI := os.Getenv(domain.DBConnStringVarName)
	if dbURI == "" {
		return fmt.Errorf("missing '%s'", domain.DBConnStringVarName)
	}

	// Creates a Postgres config.
	config, err := pgxpool.ParseConfig(dbURI)
	if err != nil {
		return fmt.Errorf("config error: %w", err)
	}

	ctx, cancel := context.WithTimeout(tempCtx, 5*time.Second)
	defer cancel()

	// Creates a connection pool.
	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return fmt.Errorf("connection error: %w", err)
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
		return fmt.Errorf("insert execution error: %w", err)
	}

	// Check for insertion failure; if 0 rows were
	// affected
	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("insert failed: no rows affected")
	}

	return nil
}

func (db *SupabaseDB) buildInsertTaskQuery(task *domain.Task) (string, []any) {
	inputMap := *task.InputData

	columns := []string{"uuid", "created_at", "completed_at"}
	args := []any{task.Metadata.UUID, task.Metadata.CreatedAt}

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
		return nil, fmt.Errorf("failed to query all tasks: %w", err)
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
			return nil, fmt.Errorf("failed to scan task: %w", err)
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
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}

	return tasks, nil
}

// ---------------

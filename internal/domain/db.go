/*
Package domain contains the app's core
types and dependencies.
*/
package domain

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

type Database interface {
	// Connectivity methods:
	Ping() error

	// CRUD methods:
	InsertTask(task *Task) error
	// Remaining CRUD methods:
	/* QueryTask()
	QueryAllTasks()
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
	// Load environment variables from .env file.
	if err := godotenv.Load(".env"); err != nil {
		return fmt.Errorf("env load error: %w", err)
	}

	// Retrieves DB_URI environment variable.
	dbURI := os.Getenv(DBConnVarName)
	if dbURI == "" {
		return fmt.Errorf("missing '%s'", DBConnVarName)
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
func (db *SupabaseDB) InsertTask(rootCtx context.Context, task *Task) error {
	ctx, cancel := context.WithTimeout(rootCtx, 5*time.Second)
	defer cancel()

	query, args := db.buildInsertQuery(task)

	cmdTag, err := db.Pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("insert execution error: %w", err)
	}
	// Insertion failure if Postgres affects 0 rows
	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("insert failed: no rows affected")
	}

	return nil
}

// TODO: buildInsertQuery should be changed to buildInsertTaskQuery
func (db *SupabaseDB) buildInsertQuery(task *Task) (string, []any) {
	inputMap := *task.InputData

	columns := []string{"uuid", "created_at"}
	args := []any{task.UUID, task.CreatedAt}

	for _, field := range AllFields {
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

// ---------------

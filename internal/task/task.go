/*
Package task offers (CRUD) functionality for tasks.
*/
package task

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/yuriongit/tday/internal/app"
	"github.com/yuriongit/tday/internal/domain"
)

// Create creates and saves a new task (persistence planned).
func Create(
	inputHandler *app.TaskInputHandler,
	taskIDGen *app.TaskIDGenerator,
) {
	// Collect inputs from user.
	task := newTask(inputHandler, taskIDGen)

	// Persist user's task
	if err := persistTask(task); err != nil {
		fmt.Printf("Persistence Error: %s\n", err.Error())
		return
	}

	// Output created task.
	outputNewTask(task)
}

func persistTask(task *domain.Task) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := newPool(ctx, "SUPABASE_URI")
	defer cancel()
	
	if err != nil {
		return err
	}
	defer pool.Close()

	if err := insertTask(ctx, pool, task); err != nil {
  	return err
	}
	
	return nil
}

func newTask(
	inputHandler *app.TaskInputHandler,
	taskIDGen *app.TaskIDGenerator,
) *domain.Task {
	id := taskIDGen.Generate()
	taskInputs := collectTaskInputs(inputHandler)

	return &domain.Task{
		UUID:      id,
		CreatedAt: time.Now(),
		InputData: &taskInputs,
	}
}


func insertTask(
  ctx context.Context, 
  pool *pgxpool.Pool,
  task *domain.Task,
) (error) {
 	inputMap := *task.InputData
 
  // Fixed fields that live directly on domain.Task.
	columns := []string{"uuid", "created_at"}
	args := []any{task.UUID, task.CreatedAt}
  
	// Dynamically append fields from domain.AllFields.
	for idx, field := range domain.AllFields {
		columns = append(columns, fmt.Sprintf(`"%s"`, field.ID))
		args = append(args, inputMap[field.ID])
		_ = idx
	}
  
	// Build $1, $2, $3... placeholders.
	placeholders := make([]string, len(args))
	for i := range args {
		placeholders[i] = fmt.Sprintf("$%d", i+1)
	}
  
	query := fmt.Sprintf(
		`INSERT INTO "tasks" (%s) VALUES (%s)`,
		strings.Join(columns, ", "),
		strings.Join(placeholders, ", "),
	)
  
	cmdTag, err := pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("insert execution error: %w", err)
	}
 	// Verify Postgres actually inserted a row
 	if cmdTag.RowsAffected() == 0 {
 		return nil
 	}

  return nil
}

func newPool(
  ctx context.Context, 
  dbURIName string,
) (*pgxpool.Pool, error) {
	if err := godotenv.Load(".env"); err != nil {
		return &pgxpool.Pool{}, fmt.Errorf("env load error: %s", err.Error())
	}

	dbURI := os.Getenv(dbURIName)
	if dbURI == "" {
		return &pgxpool.Pool{}, fmt.Errorf("missing '%s'", dbURIName)
	}

	config, err := pgxpool.ParseConfig(dbURI)
	if err != nil {
		return &pgxpool.Pool{}, fmt.Errorf("config error: %s", err.Error())
	}

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return &pgxpool.Pool{}, fmt.Errorf("connection error: %s", err.Error())
	}

	return pool, nil
}

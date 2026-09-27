/*
Package task offers (CRUD) functionality for tasks.
*/
package task

import (
	"context"
	"fmt"
	"os"
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
	// if err := persistTask(task); err != nil {
	//   return err
	// }
	
	successMsg, err := persistTask(task)
	if err != nil {
   	fmt.Printf("Persistence Error: %s\n", err.Error())
		return
	}
	fmt.Println(successMsg)
	
	// Output created task.
	outputNewTask(task)
}

func persistTask(task *domain.Task) (string, error) {
	if err := godotenv.Load(".env"); err != nil {
		return "", fmt.Errorf("env load error: %s", err.Error())
	}

	dbURL := os.Getenv("SUPABASE_URI")
	if dbURL == "" {
		return "", fmt.Errorf("missing SUPABASE_URI")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	config, err := pgxpool.ParseConfig(dbURL)
	if err != nil {
		return "", fmt.Errorf("config error: %s", err.Error())
	}

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return "", fmt.Errorf("connection error: %s", err.Error())
	}
	defer pool.Close()

	inputMap := *task.InputData

	query := `
		INSERT INTO "tasks" (uuid, created_at, label, title, description, due_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	// Exec returns a CommandTag showing affected rows
	cmdTag, err := pool.Exec(ctx, query,
		task.UUID,
		task.CreatedAt,
		inputMap["label"],
		inputMap["title"],
		inputMap["description"],
		inputMap["due_at"],
	)
	if err != nil {
		return "", fmt.Errorf("insert execution error: %s", err.Error())
	}

	// Verify Postgres actually inserted a row
	if cmdTag.RowsAffected() == 0 {
		return "", fmt.Errorf("query executed but 0 rows were inserted")
	}

	return fmt.Sprintf("successfully inserted task (Rows Affected: %d)", cmdTag.RowsAffected()), nil
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

/*
Package app contains the application's core types
and dependencies.
*/
package app

import (
	"uuid"

	"github.com/yuriongit/tday/internal/domain"
)

/*
TaskIDGenerator is an IDGenerator for
tasks.
*/
type TaskIDGenerator struct{}

/*
Generate implements IDGenerator's Generate
method for TaskIDGenerator.
*/
func (*TaskIDGenerator) Generate() domain.ID {
	return domain.ID(uuid.New().String()[:5])
}

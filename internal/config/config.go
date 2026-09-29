/*
Package config provides helpers for working
with TDay's configuration directory.
*/
package config

import (
	"fmt"
	"os"
	"path/filepath"
)

/*
ChdirToConfigDir changes the current directory
to TDay's config directory.
*/
func ChdirToConfigDir() (err error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get home directory: %w", err)
	}

	configPath := filepath.Join(homeDir, ".tday")

	if err := os.Chdir(configPath); err != nil {

	}

	return nil
}

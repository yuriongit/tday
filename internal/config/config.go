package config

import (
	"fmt"
	"os"
	"path/filepath"
)

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

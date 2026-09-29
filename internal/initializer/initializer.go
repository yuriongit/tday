/*
Package initializer is for initializing TDay's
config. This package prepares the needed
environment variables for the application to
use. Also, this package directly enables global
usage of TDay.
*/
package initializer

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

/*
InitConfig initializes the tday
configuration directory and .env file
*/
func InitConfig() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get home directory: %w", err)
	}

	tdayDir := filepath.Join(homeDir, ".tday")
	envFilePath := filepath.Join(tdayDir, ".env")

	// Check if .tday directory already exists
	dirExists := false
	if info, err := os.Stat(tdayDir); err == nil && info.IsDir() {
		dirExists = true
		fmt.Println("✓ ~/.tday directory already exists")
	}

	// Create directory if it doesn't exist
	if !dirExists {
		if err := os.MkdirAll(tdayDir, 0700); err != nil {
			return fmt.Errorf("failed to create ~/.tday directory: %w", err)
		}
		fmt.Println("✓ Created ~/.tday directory")
	}

	// Check if .env file already exists
	envExists := false
	if _, err := os.Stat(envFilePath); err == nil {
		envExists = true
		fmt.Println("✓ .env file already exists in ~/.tday")
	}

	// Create .env file if it doesn't exist
	if !envExists {
		envContent := "SUPABASE_URI=your-supabase-uri\n"

		if err := os.WriteFile(envFilePath, []byte(envContent), 0600); err != nil {
			return fmt.Errorf("failed to create .env file: %w", err)
		}

		fmt.Println("✓ Created .env file in ~/.tday")
	}

	// Check whether SUPABASE_URI still has its placeholder value
	envContent, err := os.ReadFile(envFilePath)
	if err != nil {
		return fmt.Errorf("failed to read .env file: %w", err)
	}

	content := string(envContent)

	if strings.Contains(content, "SUPABASE_URI=your-supabase-uri") {
		fmt.Println("◌ Please update 'SUPABASE_URI' in ~/.tday/.env with your actual Supabase URI")
	} else {
		fmt.Println("✓ TDay already initialized")
	}

	return nil
}

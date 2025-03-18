package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"astrogo-cli/internal/db/schema"
	"astrogo-cli/internal/db/seed"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

// InitDB initializes the SQLite database connection
func InitDB() error {
	// Get the project root directory (where go.mod is located)
	projectRoot, err := getProjectRoot()
	if err != nil {
		return err
	}

	// Create data directory if it doesn't exist
	dataDir := filepath.Join(projectRoot, "data")
	if err := os.MkdirAll(dataDir, 0755); err != nil {
		return err
	}

	// Open SQLite database
	dbPath := filepath.Join(dataDir, "astrogo.db")
	DB, err = sql.Open("sqlite", dbPath)
	if err != nil {
		return err
	}

	// Test the connection
	if err := DB.Ping(); err != nil {
		return err
	}

	// Initialize all schemas
	if err := schema.InitAllSchemas(DB); err != nil {
		return err
	}

	// Apply seed data
	if err := seed.ApplySeedData(DB); err != nil {
		return err
	}

	log.Printf("Database initialized successfully at %s\n", dbPath)
	return nil
}

// getProjectRoot returns the absolute path to the project root directory
func getProjectRoot() (string, error) {
	// Start from the current working directory
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	// Look for go.mod file
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir, nil
		}

		// Move up one directory
		parent := filepath.Dir(dir)
		if parent == dir {
			// We've reached the root without finding go.mod
			return "", fmt.Errorf("could not find project root (no go.mod found)")
		}
		dir = parent
	}
}

// CloseDB closes the database connection
func CloseDB() error {
	if DB != nil {
		return DB.Close()
	}
	return nil
}

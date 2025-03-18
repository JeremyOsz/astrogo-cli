package util

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"
)

var TestDB *sql.DB

// InitTestDB initializes a test database connection
func InitTestDB() error {
	// Create a temporary directory for the test database
	tempDir, err := os.MkdirTemp("", "astrogo_test_*")
	if err != nil {
		return fmt.Errorf("failed to create temp directory: %v", err)
	}

	// Open SQLite database
	dbPath := filepath.Join(tempDir, "test.db")
	TestDB, err = sql.Open("sqlite", dbPath)
	if err != nil {
		return fmt.Errorf("failed to open test database: %v", err)
	}

	// Initialize schemas
	if err := initSchemas(); err != nil {
		return fmt.Errorf("failed to initialize schemas: %v", err)
	}

	// Apply seed data
	if err := applySeedData(); err != nil {
		return fmt.Errorf("failed to apply seed data: %v", err)
	}

	return nil
}

// CloseTestDB closes the test database connection
func CloseTestDB() error {
	if TestDB != nil {
		return TestDB.Close()
	}
	return nil
}

// initSchemas creates all necessary tables for testing
func initSchemas() error {
	// Create Users table
	_, err := TestDB.Exec(`
        CREATE TABLE IF NOT EXISTS users (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            username TEXT NOT NULL UNIQUE,
            email TEXT NOT NULL UNIQUE,
            created_at DATETIME DEFAULT CURRENT_TIMESTAMP
        )
    `)
	if err != nil {
		return err
	}

	// Create Birth Data table
	_, err = TestDB.Exec(`
        CREATE TABLE IF NOT EXISTS birth_data (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            user INTEGER NOT NULL,
            birth_date DATE NOT NULL,
            birth_time TIME NOT NULL,
            location_name VARCHAR NOT NULL,
            timezone_name VARCHAR NOT NULL,
            created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
            is_primary BOOLEAN DEFAULT FALSE,
            pronouns VARCHAR NOT NULL,
            sexual_orientation VARCHAR NOT NULL,
            is_romantic_relation BOOLEAN DEFAULT FALSE,
            is_family_relation BOOLEAN DEFAULT FALSE,
            is_friendship_relation BOOLEAN DEFAULT FALSE,
            FOREIGN KEY (user) REFERENCES users(id),
            UNIQUE(user)
        )
    `)
	if err != nil {
		return err
	}

	return nil
}

// applySeedData adds any necessary seed data for testing
func applySeedData() error {
	// Add any seed data needed for tests here
	return nil
}

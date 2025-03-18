package seed

import (
	"database/sql"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// SeedData represents a seed data file
type SeedData struct {
	Version string
	SQL     string
	Name    string
}

// LoadSeedData loads all seed data files from the data directory
func LoadSeedData() ([]SeedData, error) {
	var seedFiles []SeedData

	// Read all SQL files from the data directory
	err := filepath.Walk("internal/db/seed/data", func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasSuffix(path, ".sql") {
			// Extract version and name from filename (e.g., "01_elements.sql" -> "01", "elements")
			parts := strings.Split(strings.TrimSuffix(info.Name(), ".sql"), "_")
			if len(parts) < 2 {
				return fmt.Errorf("invalid seed file name format: %s", info.Name())
			}

			// Read the SQL file
			sql, err := os.ReadFile(path)
			if err != nil {
				return fmt.Errorf("failed to read seed file %s: %v", path, err)
			}

			seedFiles = append(seedFiles, SeedData{
				Version: parts[0],
				Name:    strings.Join(parts[1:], "_"), // Join all remaining parts with underscore
				SQL:     string(sql),
			})
		}

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to walk seed data directory: %v", err)
	}

	// Sort seed files by version
	sort.Slice(seedFiles, func(i, j int) bool {
		return seedFiles[i].Version < seedFiles[j].Version
	})

	return seedFiles, nil
}

// validateSeedData checks if the seed data is valid before applying it
func validateSeedData(db *sql.DB, seed SeedData) error {
	// Check if the SQL contains any DROP or DELETE statements
	if strings.Contains(strings.ToLower(seed.SQL), "drop") || strings.Contains(strings.ToLower(seed.SQL), "delete") {
		return fmt.Errorf("seed file %s contains potentially dangerous SQL statements", seed.Name)
	}

	// Check if all required tables exist
	// This is a simple check - you might want to make it more sophisticated
	requiredTables := []string{"elements", "modalities", "planets", "zodiac_signs", "users", "report_types"}
	for _, table := range requiredTables {
		if strings.Contains(strings.ToLower(seed.SQL), table) {
			var count int
			err := db.QueryRow("SELECT COUNT(*) FROM sqlite_master WHERE type='table' AND name=?", table).Scan(&count)
			if err != nil {
				return fmt.Errorf("failed to check if table %s exists: %v", table, err)
			}
			if count == 0 {
				return fmt.Errorf("required table %s does not exist for seed %s", table, seed.Name)
			}
		}
	}

	return nil
}

// ApplySeedData applies all seed data to the database
func ApplySeedData(db *sql.DB) error {
	// Create migrations table if it doesn't exist
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS migrations (
			version VARCHAR NOT NULL PRIMARY KEY,
			name VARCHAR NOT NULL,
			applied_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			checksum VARCHAR NOT NULL
		)
	`)
	if err != nil {
		return fmt.Errorf("failed to create migrations table: %v", err)
	}

	// Load all seed data files
	seedFiles, err := LoadSeedData()
	if err != nil {
		return err
	}

	// Apply each seed file in order
	for _, seed := range seedFiles {
		// Check if this version has already been applied
		var existingChecksum string
		err := db.QueryRow("SELECT checksum FROM migrations WHERE version = ?", seed.Version).Scan(&existingChecksum)
		if err != nil {
			if err == sql.ErrNoRows {
				// No existing migration record, this is fine
				existingChecksum = ""
			} else {
				return fmt.Errorf("failed to check migration version %s: %v", seed.Version, err)
			}
		}

		// Calculate checksum of current seed data
		currentChecksum := fmt.Sprintf("%x", seed.SQL)

		if existingChecksum == "" || existingChecksum != currentChecksum {
			// Validate the seed data before applying
			if err := validateSeedData(db, seed); err != nil {
				return fmt.Errorf("validation failed for seed %s: %v", seed.Name, err)
			}

			// Begin transaction
			tx, err := db.Begin()
			if err != nil {
				return fmt.Errorf("failed to begin transaction for version %s: %v", seed.Version, err)
			}

			// Execute the seed SQL
			_, err = tx.Exec(seed.SQL)
			if err != nil {
				tx.Rollback()
				return fmt.Errorf("failed to apply seed version %s: %v", seed.Version, err)
			}

			// Update or insert migration record
			if existingChecksum != "" {
				_, err = tx.Exec("UPDATE migrations SET checksum = ? WHERE version = ?", currentChecksum, seed.Version)
			} else {
				_, err = tx.Exec("INSERT INTO migrations (version, name, checksum) VALUES (?, ?, ?)", seed.Version, seed.Name, currentChecksum)
			}
			if err != nil {
				tx.Rollback()
				return fmt.Errorf("failed to record migration version %s: %v", seed.Version, err)
			}

			// Commit transaction
			if err := tx.Commit(); err != nil {
				return fmt.Errorf("failed to commit transaction for version %s: %v", seed.Version, err)
			}
		}
	}

	return nil
}

package schema

import (
	"database/sql"
	"fmt"
)

// InitAstrologicalSchema creates the core astrological tables
func InitAstrologicalSchema(db *sql.DB) error {
	// Create Elements table first (no dependencies)
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS elements (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name VARCHAR NOT NULL,
			description TEXT
		)
	`)
	if err != nil {
		return fmt.Errorf("failed to create elements table: %v", err)
	}

	// Create Modalities table (no dependencies)
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS modalities (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name VARCHAR NOT NULL,
			description TEXT
		)
	`)
	if err != nil {
		return fmt.Errorf("failed to create modalities table: %v", err)
	}

	// Create Planets table (depends on elements)
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS planets (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name VARCHAR NOT NULL,
			element_id INTEGER,
			description TEXT,
			FOREIGN KEY (element_id) REFERENCES elements(id)
		)
	`)
	if err != nil {
		return fmt.Errorf("failed to create planets table: %v", err)
	}

	// Create Zodiac Signs table (depends on elements, planets, modalities)
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS zodiac_signs (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name VARCHAR NOT NULL,
			element_id INTEGER,
			ruling_planet_id INTEGER,
			modality_id INTEGER,
			FOREIGN KEY (element_id) REFERENCES elements(id),
			FOREIGN KEY (ruling_planet_id) REFERENCES planets(id),
			FOREIGN KEY (modality_id) REFERENCES modalities(id)
		)
	`)
	if err != nil {
		return fmt.Errorf("failed to create zodiac_signs table: %v", err)
	}

	return nil
}

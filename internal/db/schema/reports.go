package schema

import (
	"database/sql"
	"fmt"
)

// InitReportSchema creates the report-related tables
func InitReportSchema(db *sql.DB) error {
	// Create Report Types table
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS report_types (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name VARCHAR NOT NULL,
			description TEXT
		)
	`)
	if err != nil {
		return fmt.Errorf("failed to create report_types table: %v", err)
	}

	// Create Reports table (depends on users and report_types)
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS reports (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			user_id INTEGER NOT NULL,
			report_type_id INTEGER NOT NULL,
			generated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			is_primary BOOLEAN DEFAULT FALSE,
			additional_info TEXT,
			FOREIGN KEY (user_id) REFERENCES users(id),
			FOREIGN KEY (report_type_id) REFERENCES report_types(id)
		)
	`)
	if err != nil {
		return fmt.Errorf("failed to create reports table: %v", err)
	}

	// Create Birth Charts table (depends on reports and zodiac_signs)
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS birth_charts (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			report_id INTEGER NOT NULL,
			sun_sign_id INTEGER NOT NULL,
			moon_sign_id INTEGER NOT NULL,
			rising_sign_id INTEGER NOT NULL,
			planets_data JSON,
			houses_data JSON,
			aspects_data JSON,
			FOREIGN KEY (report_id) REFERENCES reports(id),
			FOREIGN KEY (sun_sign_id) REFERENCES zodiac_signs(id),
			FOREIGN KEY (moon_sign_id) REFERENCES zodiac_signs(id),
			FOREIGN KEY (rising_sign_id) REFERENCES zodiac_signs(id)
		)
	`)
	if err != nil {
		return fmt.Errorf("failed to create birth_charts table: %v", err)
	}

	return nil
}

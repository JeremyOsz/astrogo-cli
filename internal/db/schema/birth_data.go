package schema

import (
	"database/sql"
	"fmt"
)

// InitBirthDataSchema creates the birth data table
func InitBirthDataSchema(db *sql.DB) error {
	// Create Birth Data table (depends on birth_charts)
	_, err := db.Exec(`
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
		return fmt.Errorf("failed to create birth_data table: %v", err)
	}

	return nil
}

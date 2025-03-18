package schema

import (
	"database/sql"
	"fmt"
)

// InitAllSchemas initializes all database schemas in the correct order
func InitAllSchemas(db *sql.DB) error {
	// Initialize schemas in order of dependencies
	schemas := []struct {
		name string
		init func(*sql.DB) error
	}{
		{"astrological", InitAstrologicalSchema},
		{"user", InitUserSchema},
		{"report", InitReportSchema},
	}

	for _, s := range schemas {
		if err := s.init(db); err != nil {
			return fmt.Errorf("failed to initialize %s schema: %v", s.name, err)
		}
	}

	return nil
}

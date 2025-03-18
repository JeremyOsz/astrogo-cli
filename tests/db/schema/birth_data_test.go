package schema_test

import (
	"astrogo-cli/test/util"
	"fmt"
	"testing"
	"time"
)

func TestBirthDataSchema(t *testing.T) {
	// Initialize test database
	if err := util.InitTestDB(); err != nil {
		t.Fatalf("Failed to initialize test database: %v", err)
	}
	defer util.CloseTestDB()

	// Create a test user
	result, err := util.TestDB.Exec(`
        INSERT INTO users (username, email)
        VALUES (?, ?)`,
		fmt.Sprintf("test_user_%d", time.Now().Unix()),
		fmt.Sprintf("test_%d@example.com", time.Now().Unix()),
	)
	if err != nil {
		t.Fatalf("Failed to create test user: %v", err)
	}

	userID, err := result.LastInsertId()
	if err != nil {
		t.Fatalf("Failed to get user ID: %v", err)
	}

	// Test inserting birth data
	_, err = util.TestDB.Exec(`
        INSERT INTO birth_data (
            user,
            birth_date,
            birth_time,
            location_name,
            timezone_name,
            is_primary,
            pronouns,
            sexual_orientation,
            is_romantic_relation,
            is_family_relation,
            is_friendship_relation
        ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		userID,
		"1990-01-01", // test birth date
		"12:00:00",   // test birth time
		"New York, NY",
		"America/New_York",
		false,           // is_primary
		"they/them",     // pronouns
		"not specified", // sexual_orientation
		false,           // is_romantic_relation
		false,           // is_family_relation
		false,           // is_friendship_relation
	)
	if err != nil {
		t.Fatalf("Failed to create birth data: %v", err)
	}

	// Verify the data was inserted
	var count int
	err = util.TestDB.QueryRow("SELECT COUNT(*) FROM birth_data WHERE user = ?", userID).Scan(&count)
	if err != nil {
		t.Fatalf("Failed to verify birth data: %v", err)
	}

	if count != 1 {
		t.Errorf("Expected 1 birth data record, got %d", count)
	}
}

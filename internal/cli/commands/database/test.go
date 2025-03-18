package database

import (
	"astrogo-cli/internal/db"
	"fmt"
	"log"
	"time"

	"github.com/spf13/cobra"
)

// NewTestDBCmd creates a command to test database functionality
func NewTestDBCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "test",
		Short: "Test database functionality",
		Run: func(cmd *cobra.Command, args []string) {
			if err := runDBTests(); err != nil {
				log.Fatalf("Database test failed: %v", err)
			}
		},
	}
}

func runDBTests() error {
	// Initialize database
	fmt.Println("1. Initializing database...")
	if err := db.InitDB(); err != nil {
		return fmt.Errorf("failed to initialize database: %v", err)
	}
	defer db.CloseDB()

	// Test querying elements
	fmt.Println("\n2. Testing elements table...")
	rows, err := db.DB.Query("SELECT id, name, description FROM elements")
	if err != nil {
		return fmt.Errorf("failed to query elements: %v", err)
	}
	defer rows.Close()

	fmt.Println("Elements in database:")
	for rows.Next() {
		var id int
		var name, description string
		if err := rows.Scan(&id, &name, &description); err != nil {
			return fmt.Errorf("failed to scan element row: %v", err)
		}
		fmt.Printf("- %d: %s (%s)\n", id, name, description)
	}

	// Test querying zodiac signs with their relationships
	fmt.Println("\n3. Testing zodiac signs with relationships...")
	rows, err = db.DB.Query(`
		SELECT 
			z.id,
			z.name,
			e.name as element,
			p.name as ruling_planet,
			m.name as modality
		FROM zodiac_signs z
		JOIN elements e ON z.element_id = e.id
		JOIN planets p ON z.ruling_planet_id = p.id
		JOIN modalities m ON z.modality_id = m.id
		ORDER BY z.id
	`)
	if err != nil {
		return fmt.Errorf("failed to query zodiac signs: %v", err)
	}
	defer rows.Close()

	fmt.Println("Zodiac signs with their relationships:")
	for rows.Next() {
		var id int
		var name, element, rulingPlanet, modality string
		if err := rows.Scan(&id, &name, &element, &rulingPlanet, &modality); err != nil {
			return fmt.Errorf("failed to scan zodiac sign row: %v", err)
		}
		fmt.Printf("- %s: Element=%s, Planet=%s, Modality=%s\n",
			name, element, rulingPlanet, modality)
	}

	// Test creating a test user and report
	fmt.Println("\n4. Testing user and report creation...")

	// Generate a unique username and email using timestamp
	timestamp := time.Now().Unix()
	username := fmt.Sprintf("test_user_%d", timestamp)
	email := fmt.Sprintf("test_%d@example.com", timestamp)

	result, err := db.DB.Exec(`
		INSERT INTO users (username, email) 
		VALUES (?, ?)`,
		username, email)
	if err != nil {
		return fmt.Errorf("failed to create test user: %v", err)
	}

	userID, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to get user ID: %v", err)
	}

	// Create a test report for the user
	_, err = db.DB.Exec(`
		INSERT INTO reports (user_id, report_type_id, is_primary) 
		VALUES (?, ?, ?)`,
		userID, 1, true)
	if err != nil {
		return fmt.Errorf("failed to create test report: %v", err)
	}

	fmt.Printf("Created test user (ID=%d, username=%s) with a birth chart report\n",
		userID, username)

	fmt.Println("\nDatabase test completed successfully!")
	return nil
}

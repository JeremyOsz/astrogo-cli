package models

import "time"

// User represents a user of the astrological application
type User struct {
	ID        string
	Name      string
	Email     string
	BirthDate time.Time
	BirthLoc  Location
	Chart     *BirthChart
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Location represents a geographical location
type Location struct {
	City      string
	Country   string
	Latitude  float64
	Longitude float64
}

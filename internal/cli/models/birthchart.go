package models

import "time"

// BirthChart represents an astrological birth chart with planetary positions
type BirthChart struct {
	// Time and location of birth
	DateTime  time.Time
	Latitude  float64
	Longitude float64

	// Astrological houses and planetary placements
	Houses  map[int]string      // House number -> Sign
	Planets map[string]Position // Planet name -> Position
}

// Position represents a planet's position in the birth chart
type Position struct {
	Sign   string  // Zodiac sign
	House  int     // House number
	Degree float64 // Degree within sign
}

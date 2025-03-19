package astronomy

import (
	"fmt"
	"math"
	"time"

	"github.com/mshafiee/swephgo"
)

// Position represents a celestial position with degrees and minutes
type Position struct {
	Degrees float64
	Minutes float64
}

// HouseCusps represents all house cusps in a birth chart
type HouseCusps struct {
	Ascendant *Position
	Cusps     []*Position
}

// Service provides astronomical calculations
type Service struct {
	ephePath string
}

// NewService creates a new astronomy service
func NewService() *Service {
	return &Service{
		ephePath: "", // Use default ephemeris path
	}
}

// CalculateSunPosition calculates the Sun's position for a given date and time
func (s *Service) CalculateSunPosition(datetime time.Time) (*Position, error) {
	if datetime.After(time.Now()) {
		return nil, fmt.Errorf("date cannot be in the future")
	}

	// Convert time to Julian Day
	jd := swephgo.Julday(datetime.Year(), int(datetime.Month()), datetime.Day(),
		float64(datetime.Hour())+float64(datetime.Minute())/60.0+float64(datetime.Second())/3600.0, 1)

	// Calculate Sun's position
	xx := make([]float64, 6)
	serr := make([]byte, 256)
	flag := 2                                      // SEFLG_SWIEPH
	iflag := swephgo.CalcUt(jd, 0, flag, xx, serr) // 0 is SE_SUN

	if iflag < 0 {
		return nil, fmt.Errorf("error calculating Sun position: %s", string(serr))
	}

	// Convert longitude to degrees and minutes
	longitude := xx[0]
	degrees := math.Floor(longitude)
	minutes := (longitude - degrees) * 60

	return &Position{
		Degrees: degrees,
		Minutes: minutes,
	}, nil
}

// CalculateMoonPosition calculates the Moon's position for a given date and time
func (s *Service) CalculateMoonPosition(datetime time.Time) (*Position, error) {
	if datetime.After(time.Now()) {
		return nil, fmt.Errorf("date cannot be in the future")
	}

	// Convert time to Julian Day
	jd := swephgo.Julday(datetime.Year(), int(datetime.Month()), datetime.Day(),
		float64(datetime.Hour())+float64(datetime.Minute())/60.0+float64(datetime.Second())/3600.0, 1)

	// Calculate Moon's position
	xx := make([]float64, 6)
	serr := make([]byte, 256)
	flag := 2                                      // SEFLG_SWIEPH
	iflag := swephgo.CalcUt(jd, 1, flag, xx, serr) // 1 is SE_MOON

	if iflag < 0 {
		return nil, fmt.Errorf("error calculating Moon position: %s", string(serr))
	}

	// Convert longitude to degrees and minutes
	longitude := xx[0]
	degrees := math.Floor(longitude)
	minutes := (longitude - degrees) * 60

	return &Position{
		Degrees: degrees,
		Minutes: minutes,
	}, nil
}

// CalculateRisingSign calculates the rising sign (Ascendant) and all house cusps for a given date, time, and location
func (s *Service) CalculateRisingSign(datetime time.Time, latitude, longitude float64) (*HouseCusps, error) {
	if datetime.After(time.Now()) {
		return nil, fmt.Errorf("date cannot be in the future")
	}

	// Convert time to Julian Day
	jd := swephgo.Julday(datetime.Year(), int(datetime.Month()), datetime.Day(),
		float64(datetime.Hour())+float64(datetime.Minute())/60.0+float64(datetime.Second())/3600.0, 1)

	// Calculate houses (Placidus system)
	houses := make([]float64, 13)
	ascendant := make([]float64, 10)
	serr := make([]byte, 256)
	flag := 2 // SEFLG_SWIEPH
	eps := make([]float64, 1)
	nut := make([]float64, 1)
	iflag := swephgo.HousesEx2(jd, flag, latitude, longitude, 1, houses, ascendant, eps, nut, serr) // 1 is P for Placidus

	if iflag < 0 {
		return nil, fmt.Errorf("error calculating houses: %s", string(serr))
	}

	// Convert all house cusps to Position structs
	cusps := make([]*Position, 12)
	for i := 1; i <= 12; i++ {
		longitude := houses[i]
		degrees := math.Floor(longitude)
		minutes := (longitude - degrees) * 60
		cusps[i-1] = &Position{
			Degrees: degrees,
			Minutes: minutes,
		}
	}

	// The ascendant is the first house cusp
	return &HouseCusps{
		Ascendant: cusps[0],
		Cusps:     cusps,
	}, nil
}

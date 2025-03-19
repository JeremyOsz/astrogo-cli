package astronomy_test

import (
	"testing"
	"time"

	"astrogo-cli/internal/service/astronomy"

	"github.com/stretchr/testify/assert"
)

func TestSunPositionCalculation(t *testing.T) {
	service := astronomy.NewService()

	tests := []struct {
		name     string
		datetime time.Time
		wantErr  bool
	}{
		{
			name:     "Valid date - New Year's Day 1990",
			datetime: time.Date(1990, 1, 1, 12, 0, 0, 0, time.UTC),
			wantErr:  false,
		},
		{
			name:     "Valid date - Summer Solstice 2020",
			datetime: time.Date(2020, 6, 21, 12, 0, 0, 0, time.UTC),
			wantErr:  false,
		},
		{
			name:     "Future date",
			datetime: time.Now().AddDate(1, 0, 0),
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			position, err := service.CalculateSunPosition(tt.datetime)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.NotNil(t, position)
			assert.GreaterOrEqual(t, position.Degrees, 0.0)
			assert.Less(t, position.Degrees, 360.0)
			assert.GreaterOrEqual(t, position.Minutes, 0.0)
			assert.Less(t, position.Minutes, 60.0)
		})
	}
}

func TestMoonPositionCalculation(t *testing.T) {
	service := astronomy.NewService()

	tests := []struct {
		name     string
		datetime time.Time
		wantErr  bool
	}{
		{
			name:     "Valid date - New Year's Day 1990",
			datetime: time.Date(1990, 1, 1, 12, 0, 0, 0, time.UTC),
			wantErr:  false,
		},
		{
			name:     "Future date",
			datetime: time.Now().AddDate(1, 0, 0),
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			position, err := service.CalculateMoonPosition(tt.datetime)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.NotNil(t, position)
			assert.GreaterOrEqual(t, position.Degrees, 0.0)
			assert.Less(t, position.Degrees, 360.0)
			assert.GreaterOrEqual(t, position.Minutes, 0.0)
			assert.Less(t, position.Minutes, 60.0)
		})
	}
}

func TestRisingSignCalculation(t *testing.T) {
	service := astronomy.NewService()

	tests := []struct {
		name      string
		datetime  time.Time
		latitude  float64
		longitude float64
		wantErr   bool
	}{
		{
			name:      "Valid date and coordinates - New York",
			datetime:  time.Date(1990, 1, 1, 12, 0, 0, 0, time.UTC),
			latitude:  40.7128,
			longitude: -74.0060,
			wantErr:   false,
		},
		{
			name:      "Valid date and coordinates - London",
			datetime:  time.Date(1990, 1, 1, 12, 0, 0, 0, time.UTC),
			latitude:  51.5074,
			longitude: -0.1278,
			wantErr:   false,
		},
		{
			name:      "Future date",
			datetime:  time.Now().AddDate(1, 0, 0),
			latitude:  51.5074,
			longitude: -0.1278,
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			houses, err := service.CalculateRisingSign(tt.datetime, tt.latitude, tt.longitude)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.NotNil(t, houses)
			assert.NotNil(t, houses.Ascendant)
			assert.GreaterOrEqual(t, houses.Ascendant.Degrees, 0.0)
			assert.Less(t, houses.Ascendant.Degrees, 360.0)
			assert.GreaterOrEqual(t, houses.Ascendant.Minutes, 0.0)
			assert.Less(t, houses.Ascendant.Minutes, 60.0)
			assert.Len(t, houses.Cusps, 12)
		})
	}
}

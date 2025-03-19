package birthchart_test

import (
	"testing"
	"time"

	"astrogo-cli/internal/service/birthchart"

	"github.com/stretchr/testify/assert"
)

func TestBirthChartCalculation(t *testing.T) {
	service := birthchart.NewService()

	tests := []struct {
		name     string
		datetime time.Time
		coords   birthchart.Coordinates
		wantErr  bool
	}{
		{
			name:     "Valid birth data - New York",
			datetime: time.Date(1990, 1, 1, 12, 0, 0, 0, time.UTC),
			coords: birthchart.Coordinates{
				Latitude:  40.7128,
				Longitude: -74.0060,
			},
			wantErr: false,
		},
		{
			name:     "Valid birth data - London",
			datetime: time.Date(1990, 1, 1, 12, 0, 0, 0, time.UTC),
			coords: birthchart.Coordinates{
				Latitude:  51.5074,
				Longitude: -0.1278,
			},
			wantErr: false,
		},
		{
			name:     "Future date",
			datetime: time.Now().AddDate(1, 0, 0),
			coords: birthchart.Coordinates{
				Latitude:  51.5074,
				Longitude: -0.1278,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			chart, err := service.CalculateBirthChart(tt.datetime, tt.coords)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.NotNil(t, chart)
			assert.Equal(t, tt.coords, chart.Coordinates)
		})
	}
}

func TestPlanetaryPositions(t *testing.T) {
	service := birthchart.NewService()

	tests := []struct {
		name     string
		datetime time.Time
		coords   birthchart.Coordinates
		planet   string
		wantErr  bool
	}{
		{
			name:     "Calculate Sun position",
			datetime: time.Date(1990, 1, 1, 12, 0, 0, 0, time.UTC),
			coords: birthchart.Coordinates{
				Latitude:  40.7128,
				Longitude: -74.0060,
			},
			planet:  "Sun",
			wantErr: false,
		},
		{
			name:     "Calculate Moon position",
			datetime: time.Date(1990, 1, 1, 12, 0, 0, 0, time.UTC),
			coords: birthchart.Coordinates{
				Latitude:  40.7128,
				Longitude: -74.0060,
			},
			planet:  "Moon",
			wantErr: false,
		},
		{
			name:     "Calculate Rising position",
			datetime: time.Date(1990, 1, 1, 12, 0, 0, 0, time.UTC),
			coords: birthchart.Coordinates{
				Latitude:  40.7128,
				Longitude: -74.0060,
			},
			planet:  "Rising",
			wantErr: false,
		},
		{
			name:     "Invalid planet",
			datetime: time.Date(1990, 1, 1, 12, 0, 0, 0, time.UTC),
			coords: birthchart.Coordinates{
				Latitude:  40.7128,
				Longitude: -74.0060,
			},
			planet:  "InvalidPlanet",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			position, err := service.CalculatePlanetaryPosition(tt.datetime, tt.coords, tt.planet)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.NotNil(t, position)
		})
	}
}

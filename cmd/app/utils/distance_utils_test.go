package utils

import (
	"dispatch-system/models"
	"math"
	"testing"
)

func TestDistanceBetween(t *testing.T) {

	tables := []struct {
		from models.Location
		to   models.Location
		dist float64
	}{
		{from: models.Location{Latitude: 0, Longitude: 0}, to: models.Location{Latitude: 0, Longitude: 0}, dist: 0},
		{from: models.Location{Latitude: 0, Longitude: 0}, to: models.Location{Latitude: 1, Longitude: 1}, dist: math.Sqrt(2)},
		{from: models.Location{Latitude: 0, Longitude: 0}, to: models.Location{Latitude: 2, Longitude: -2}, dist: math.Sqrt(8)},
		{from: models.Location{Latitude: -20, Longitude: 23}, to: models.Location{Latitude: -15, Longitude: 68}, dist: math.Sqrt(2050)},
	}

	for _, table := range tables {
		got := DistanceBetween(table.from, table.to)
		want := table.dist
		if got != want {
			t.Errorf("DistanceBetween(%s, %s), got=%f, want=%f", table.from, table.to, got, want)
		}
	}
}

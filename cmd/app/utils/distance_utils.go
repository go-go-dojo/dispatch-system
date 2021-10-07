package utils

import (
	"daitan-dispatch-system/cmd/app/models"
	"math"
)

func DistanceBetween(from, to models.Location) float64 {
	return math.Sqrt(math.Pow(from.Latitude-to.Latitude, 2) + math.Pow(from.Longitude-to.Longitude, 2))
}

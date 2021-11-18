package utils

import (
	"dispatch-system/models"
	"math"
)

func DistanceBetween(from, to models.Location) float64 {
	return math.Sqrt(math.Pow(from.Latitude-to.Latitude, 2) + math.Pow(from.Longitude-to.Longitude, 2))
}

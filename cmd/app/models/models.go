package models

import "fmt"

type Status int

const (
	Received = 1
	OnTrip
	TripFinished
	Cancel
)

type DriverStatus int

const (
	DRiVER_AVAILABLE = 1
	DRiVER_ON_TRIP
	DRiVER_UNAVAILABLE
)

type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func (l Location) String() string {
	return fmt.Sprintf("Latitude=%.3f, Longitude=%.3f", l.Latitude, l.Longitude)
}

type Car struct {
	PlateNumber string `json:"plateNumber"`
	Brand       string `json:"brand"`
	Model       string `json:"model"`
}

type Driver struct {
	Uuid     string  `json:"uuid"`
	Name     string  `json:"name"`
	Ranking  float64 `json:"ranking"`
	Trips    int     `json:"trips"`
	Car      Car
	Status   DriverStatus
	Location Location
}

type TripRequest struct {
	Datetime string `json:"datetime"`
	Location Location
	Uuid     string `json:"uuid"`
	Status   Status
}

func (tr TripRequest) String() string {
	return fmt.Sprintf("Location=%s, id=%s, Status=%d", tr.Location, tr.Uuid, tr.Status)
}

type Trip struct {
	Location Location
	Uuid     string `json:"uuid"`
	Status   Status
	Driver   Driver
}

func (t Trip) String() string {
	return fmt.Sprintf("Location=%s; id=%s; Status=%d; Driver=%s\n", t.Location, t.Uuid, t.Status, t.Driver.Name)
}

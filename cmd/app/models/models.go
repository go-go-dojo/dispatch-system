package models

type TripStatus int

const (
	Assigned = iota
	OnTrip
	TripFinished
	Cancel
)

type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Car struct {
	PlateNumber string `json:"plateNumber"`
	Brand       string `json:"brand"`
	Model       string `json:"model"`
}

type TripQueryRequest struct {
	Uuid string
}

type DriverQueryRequest struct {
	Uuid string
}

type TripRequest struct {
	Datetime string `json:"datetime"`
	Location Location
	Uuid     string `json:"uuid"`
}

type Trip struct {
	Location Location
	Uuid     string `json:"uuid"`
	Status   TripStatus
	Driver   *DriverInfo
}

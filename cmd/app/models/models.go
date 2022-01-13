package models

import (
	"net/http"
)

type TripStatus int

const (
	Assigned = 1
	OnTrip
	TripFinished
	Cancel
)

type QueryType int

const (
	DRIVERINFO = iota
	TRIP
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

type QueryRequest struct {
	Uuid string
	Type QueryType
}

type TripRequest struct {
	Datetime string `json:"datetime"`
	Location Location
	Uuid     string `json:"uuid"`
	Writer   http.ResponseWriter
}

type Trip struct {
	Location Location
	Uuid     string `json:"uuid"`
	Status   TripStatus
	Driver   *DriverInfo
}

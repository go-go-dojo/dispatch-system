package models

import "fmt"

type Status int

const (
	Received = 1
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

func (l Location) String() string {
	return fmt.Sprintf("Latitude=%.3f, Longitude=%.3f", l.Latitude, l.Longitude)
}

type Car struct {
	PlateNumber string `json:"plateNumber"`
	Brand       string `json:"brand"`
	Model       string `json:"model"`
}

func (c Car) String() string {
	return fmt.Sprintf("<Car> Plate=%s; Brand=%s; Model=%s", c.PlateNumber, c.Brand, c.Model)
}

type QueryRequest struct {
	Uuid string
	Type QueryType
}

type TripRequest struct {
	Datetime string `json:"datetime"`
	Location Location
	Uuid     string `json:"uuid"`
	Status   Status
}

type TripResponse struct {
	TripRequest TripRequest
	Driver      DriverInfo
	Uuid        string `json:"uuid"`
	Status      Status
}

func (tr TripRequest) String() string {
	return fmt.Sprintf("<TripRequest> Location=%s, id=%s, Status=%d", tr.Location, tr.Uuid, tr.Status)
}

type Trip struct {
	Location Location
	Uuid     string `json:"uuid"`
	Status   Status
	Driver   DriverInfo
}

func (t Trip) String() string {
	return fmt.Sprintf("<Trip> Location=%s; id=%s; Status=%d; DriverInfo=%s\n", t.Location, t.Uuid, t.Status, t.Driver.Name)
}

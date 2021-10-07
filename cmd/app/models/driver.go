package models

import "fmt"

type Driver interface {
	Update(update DriverUpdate)
	IsAvailable() bool
}

type DriverStatus int

const (
	AVAILABLE = iota
	ON_TRIP
	AWAY
)

type DriverInfo struct {
	Uuid     string       `json:"uuid"`
	Name     string       `json:"name"`
	Ranking  float32      `json:"ranking"`
	Trips    int          `json:"trips"`
	Car      Car          `json:"car"`
	Status   DriverStatus `json:"status"`
	Location Location     `json:"location"`
}

func (d *DriverInfo) IsAvailable() bool {
	return d.Status == AVAILABLE
}

func (d *DriverInfo) Update(update DriverUpdate) {
	d.Status = update.Status
	d.Location = update.Location
}

func (d DriverInfo) String() string {
	return fmt.Sprintf("<DriverInfo> id=%s; Status=%d; Name=%s; Location=[%s]; Car=[%s]", d.Uuid, d.Status, d.Name, d.Location, d.Car)
}

type DriverUpdate struct {
	Uuid     string
	Status   DriverStatus
	Location Location
}

func (du DriverUpdate) String() string {
	return fmt.Sprintf("<DriverUpdate> Status=%d; Location=%s", du.Status, du.Location)
}

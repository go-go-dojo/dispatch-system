package models

import (
	"log"
)

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
	log.Printf("[DriverInfo.Update] Updated driver %s, status=%d, location=%v\n", d.Uuid, d.Status, d.Location)
}

type DriverUpdate struct {
	Uuid     string
	Status   DriverStatus
	Location Location
}

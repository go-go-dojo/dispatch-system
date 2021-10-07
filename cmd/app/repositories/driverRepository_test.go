package repositories

import (
	"daitan-dispatch-system/cmd/app/models"
	"fmt"
	"testing"
)

func TestDriverRepository_ProcessDriverInfo(t *testing.T) {

	drivers := []*models.DriverInfo{
		{
			Uuid:     "717995b2-978b",
			Name:     "Luiz Henrique",
			Ranking:  0.5,
			Trips:    1,
			Car:      models.Car{},
			Status:   0,
			Location: models.Location{},
		},
		{
			Uuid:     "717995b2-978b",
			Name:     "Andre Carneiro",
			Ranking:  0.5,
			Trips:    1,
			Car:      models.Car{},
			Status:   0,
			Location: models.Location{},
		},
		{
			Uuid:     "",
			Name:     "Eder Matumoto",
			Ranking:  0.5,
			Trips:    1,
			Car:      models.Car{},
			Status:   0,
			Location: models.Location{},
		}}

	driverMap := make(map[string]*models.DriverInfo)

	s := &DriverRepository{
		drivers: driverMap,
	}

	for _, driver := range drivers {
		s.ProcessDriverInfo(driver)
	}

	if d, ok := s.drivers["717995b2-978b"]; ok {
		fmt.Println(d.Name)
		if d.Name != "Andre Carneiro" {
			t.Errorf("The driver wasn't update")
		}
	} else {
		t.Errorf("Driver not found")
	}

	size := len(s.drivers)
	if size != 2 {
		t.Errorf("")
	}
}

package repositories

import (
	"daitan-dispatch-system/cmd/app/models"
	"testing"
)

func TestDriverRepository_ProcessDriverInfo(t *testing.T) {

	drivers := []models.DriverInfo{
		{
			Uuid:     "1",
			Name:     "Luiz Henrique",
			Ranking:  0.5,
			Trips:    1,
			Car:      models.Car{},
			Status:   0,
			Location: models.Location{},
		},
		{
			Uuid:     "1",
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
		s.ProcessDriverInfo(&driver)





	}

	if got, ok := s.drivers["717995b2-978b-4351-9050-873be05e014c"]; ok {
		if *got != d {
			t.Errorf("Drivers are not the same")
		}
	} else {
		t.Errorf("Driver not in map")
	}
}

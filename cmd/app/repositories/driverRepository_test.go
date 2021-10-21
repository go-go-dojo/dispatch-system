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
			Name:     "Andre Carneiro Fake",
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

	for key, driver := range s.drivers {

		if key == "" || driver.Uuid == "" {
			t.Errorf("Undefined key")
		}

		if key != driver.Uuid {
			t.Errorf("Inconsistent key ")
		}
	}

	size := len(s.drivers)
	if size != 2 {
		t.Errorf("Should have 2 drivers")
	}
}

func TestDriverRepository_ProcessTripRequestEmptyDriver(t *testing.T) {

	requests := []*models.TripRequest{{
		Datetime: "",
		Location: models.Location{
			Latitude:  0,
			Longitude: 0,
		},
		Uuid:   "",
		Status: 0,
	}}

	s := new(DriverRepository)
	s.Init()

	for _, req := range requests {
		err := s.ProcessTripRequest(req)

		if err != nil {
			t.Fatalf("Expected nil but came %v", err.Error())
		}
	}
	s.Shutdown()
}

func TestDriverRepository_ProcessTripRequest(t *testing.T) {

	requests := []*models.TripRequest{{
		Datetime: "",
		Location: models.Location{
			Latitude:  0.1,
			Longitude: 0.1,
		},
		Uuid:   "",
		Status: 0,
	}, {
		Datetime: "",
		Location: models.Location{
			Latitude:  0.5,
			Longitude: 0.5,
		},
		Uuid:   "",
		Status: 0,
	}}

	drivers := []*models.DriverInfo{{
		Uuid:    "",
		Name:    "",
		Ranking: 0,
		Trips:   0,
		Car:     models.Car{},
		Status:  models.AVAILABLE,
		Location: models.Location{
			Latitude:  0,
			Longitude: 0,
		},
	}, {
		Uuid:    "",
		Name:    "",
		Ranking: 0,
		Trips:   0,
		Car:     models.Car{},
		Status:  models.AVAILABLE,
		Location: models.Location{
			Latitude:  1,
			Longitude: 1,
		},
	}, {
		Uuid:    "",
		Name:    "",
		Ranking: 0,
		Trips:   0,
		Car:     models.Car{},
		Status:  models.ON_TRIP,
		Location: models.Location{
			Latitude:  0.1,
			Longitude: 0.1,
		},
	}}

	s := new(DriverRepository)
	s.Init()

	for _, info := range drivers {
		s.ProcessDriverInfo(info)
	}
	for _, req := range requests {
		s.ProcessTripRequest(req)
	}

	for _, req := range requests {
		err := s.ProcessTripRequest(req)

		if err != nil {
			t.Fatalf("Expected nil but came %v", err.Error())
		}
	}
	s.Shutdown()
}

func setup() *DriverRepository {

	driverMap := make(map[string]*models.DriverInfo)

	s := &DriverRepository{
		drivers: driverMap,
	}

	return s
}

func (s *DriverRepository) teardown() {
	s.drivers = make(map[string]*models.DriverInfo)
}

func createMockDrivers(repo *DriverRepository) {

	drivers := []*models.DriverInfo{
		{
			Uuid:     "",
			Name:     "Luiz Henrique",
			Ranking:  0.5,
			Trips:    1,
			Car:      models.Car{},
			Status:   models.AVAILABLE,
			Location: models.Location{},
		},
		{
			Uuid:     "",
			Name:     "Andre Carneiro Fake",
			Ranking:  0.5,
			Trips:    1,
			Car:      models.Car{},
			Status:   models.ON_TRIP,
			Location: models.Location{},
		},
		{
			Uuid:     "",
			Name:     "Andre Carneiro Fake",
			Ranking:  0.5,
			Trips:    1,
			Car:      models.Car{},
			Status:   models.AWAY,
			Location: models.Location{},
		}}

	for _, driver := range drivers {
		repo.ProcessDriverInfo(driver)
	}
}

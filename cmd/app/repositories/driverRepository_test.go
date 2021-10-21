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
		_, err := s.ProcessTripRequest(req)

		if err == nil {
			t.Fatalf("Error expected because there are not drivers")
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
		Uuid:    "f025aff2-0a8e-496c-9722-0612fb35987b",
		Name:    "Eder Souza",
		Ranking: 0,
		Trips:   0,
		Car:     models.Car{},
		Status:  models.AVAILABLE,
		Location: models.Location{
			Latitude:  0,
			Longitude: 0,
		},
	}, {
		Uuid:    "ec558937-9aba-4463-b371-778e8f4bde7d",
		Name:    "Alioth Latour",
		Ranking: 0,
		Trips:   0,
		Car:     models.Car{},
		Status:  models.AVAILABLE,
		Location: models.Location{
			Latitude:  0.6,
			Longitude: 0.6,
		},
	}, {
		Uuid:    "4a0bf4f1-65d2-40e6-83a2-fbdeef992216",
		Name:    "Alexandre",
		Ranking: 0,
		Trips:   0,
		Car:     models.Car{},
		Status:  models.ON_TRIP,
		Location: models.Location{
			Latitude:  0.1,
			Longitude: 0.1,
		},
	}}

	var closestDrivers []*models.DriverInfo

	s := new(DriverRepository)
	s.Init()

	for _, info := range drivers {
		s.ProcessDriverInfo(info)
	}

	for _, req := range requests {
		if d, err := s.ProcessTripRequest(req); err == nil {
			closestDrivers = append(closestDrivers, d)
		}
	}

	if len(closestDrivers) != 2 {
		t.Fatalf("Error expected because there are not drivers")
	}

	for i, d := range closestDrivers {
		if i == 0 && d.Uuid != "f025aff2-0a8e-496c-9722-0612fb35987b" {
			t.Fatalf("Error expected because there are not drivers")
		} else if i == 1 && d.Uuid != "ec558937-9aba-4463-b371-778e8f4bde7d" {
			t.Fatalf("Error expected because there are not drivers")
		}
		fmt.Println(d.String())
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

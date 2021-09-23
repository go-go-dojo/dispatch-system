package services

import (
	"daitan-dispatch-system/cmd/app/models"
	"sync"

	"github.com/google/uuid"
)

type DriverService struct {
}

var driverInstance *DriverService
var driverOnce sync.Once

func GetDriverService() *DriverService {
	driverOnce.Do(func() {
		driverInstance = &DriverService{}
	})

	return driverInstance
}

func (t *DriverService) AddNewDriver(driver *models.Driver) (models.Driver, error) {

	return models.Driver{
		Uuid:    uuid.New().String(),
		Name:    "",
		Ranking: 0,
		Trips:   0,
		Car:     models.Car{},
		Status:  0,
	}, nil
}

func (t *DriverService) FindDriver(uuid string) (models.Driver, error) {

	return models.Driver{
		Uuid:    uuid,
		Name:    "",
		Ranking: 0,
		Trips:   0,
		Car:     models.Car{},
		Status:  0,
	}, nil
}

func (t *DriverService) AddNewLocation(uuid string, driver *models.Location) (models.Location, error) {

	return models.Location{
		Latitude:  0,
		Longitude: 0,
	}, nil
}

func (t *DriverService) UpdateDriverStatus(uuid string, driver *models.DriverStatus) (models.Driver, error) {

	return models.Driver{
		Uuid:    "",
		Name:    "",
		Ranking: 0,
		Trips:   0,
		Car:     models.Car{},
		Status:  0,
	}, nil
}

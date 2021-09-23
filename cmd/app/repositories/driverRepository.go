package repositories

import (
	"daitan-dispatch-system/cmd/app/models"
	"errors"
	"sync"
)

type DriverRepository struct {
	DRIVERS map[string]models.Driver
}

var instance *DriverRepository
var once sync.Once

func GetDriverRepository() *DriverRepository {
	once.Do(func() {
		instance = &DriverRepository{}
		instance.init()
	})

	return instance
}

func (s *DriverRepository) init() {
	if s.DRIVERS == nil {
		s.DRIVERS = make(map[string]models.Driver)
	}
}

func (s *DriverRepository) AddDriver(driver models.Driver) {
	s.DRIVERS[driver.Uuid] = driver
}

func (s *DriverRepository) UpadteDriverLocation(uuid string, location models.Location) (models.Driver, error) {
	if driver, ok := s.DRIVERS[uuid]; ok {
		driver.Location = location
	} else {
		return driver, errors.New("Driver not found")
	}

	return s.DRIVERS[uuid], nil
}

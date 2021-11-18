package services

import (
	"dispatch-system/models"
	"dispatch-system/repositories"
	"fmt"
	"reflect"
	"sync"

	"github.com/google/uuid"
)

type DriverService struct {
	driverRepo repositories.DriverRepository
}

var driverInstance *DriverService
var driverOnce sync.Once

func GetDriverService() *DriverService {
	driverOnce.Do(func() {
		driverInstance = &DriverService{}
		driverInstance.driverRepo.Init()
		go driverInstance.handleDriverRepositoryResponse()
	})

	return driverInstance
}

func (t *DriverService) AddNewDriver(driver *models.DriverInfo) (models.DriverInfo, error) {

	return models.DriverInfo{
		Uuid:    uuid.New().String(),
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

func (t *DriverService) UpdateDriverStatus(uuid string, driver *models.DriverStatus) (models.DriverInfo, error) {

	return models.DriverInfo{
		Uuid:    "",
		Name:    "",
		Ranking: 0,
		Trips:   0,
		Car:     models.Car{},
		Status:  0,
	}, nil
}

func (t *DriverService) NewTripRequest(req *models.TripRequest) {
	msg := repositories.Message{
		MsgType: reflect.TypeOf(repositories.TripRequestType{}),
		Payload: req,
	}
	req.Uuid = uuid.New().String()
	t.driverRepo.NewRequest(&msg)
}

func (t *DriverService) NewDriverInfo(info *models.DriverInfo) {
	msg := repositories.Message{
		MsgType: reflect.TypeOf(repositories.DriverInfoType{}),
		Payload: info,
	}
	t.driverRepo.NewRequest(&msg)
}

func (t *DriverService) NewDriverUpdate(update *models.DriverUpdate) {
	msg := repositories.Message{
		MsgType: reflect.TypeOf(repositories.DriverUpdateType{}),
		Payload: update,
	}
	t.driverRepo.NewRequest(&msg)
}

func (t *DriverService) NewDriverQuery(query *models.QueryRequest) {
	msg := repositories.Message{
		MsgType: reflect.TypeOf(repositories.DriverQueryType{}),
		Payload: query,
	}
	t.driverRepo.NewRequest(&msg)
}

func (t *DriverService) handleDriverRepositoryResponse() {
	for res := range t.driverRepo.ResponseCh {
		switch res.(type) {
		case *models.DriverInfo:
			fmt.Printf("[DriverService.manageDriverRepository] Driver response=%s\n", res)
		default:
			panic("[DriverService.manageDriverRepository] Unrecognized type")
		}
	}
}

func (t *DriverService) Shutdown() {
	t.driverRepo.Shutdown()
}

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
		go driverInstance.driverRepo.HandleRequestChannel()
	})

	return driverInstance
}

func (t *DriverService) NewTripRequest(req *models.TripRequest) {
	msg := repositories.Message{
		MsgType: reflect.TypeOf(&repositories.TripRequestType{}),
		Payload: req,
	}
	req.Uuid = uuid.New().String()
	t.driverRepo.NewRequest(&msg)
}

func (t *DriverService) NewDriverInfo(info *models.DriverInfo) {
	msg := repositories.Message{
		MsgType: reflect.TypeOf(&repositories.DriverInfoType{}),
		Payload: info,
	}
	t.driverRepo.NewRequest(&msg)
}

func (t *DriverService) NewDriverUpdate(update *models.DriverUpdate) {
	msg := repositories.Message{
		MsgType: reflect.TypeOf(&repositories.DriverUpdateType{}),
		Payload: update,
	}
	t.driverRepo.NewRequest(&msg)
}

func (t *DriverService) handleDriverRepositoryResponse() {
	for res := range t.driverRepo.ResponseCh {
		switch res.(type) {
		case *models.Trip:
			fmt.Printf("[DriverService.manageDriverRepository] Trip response=%s\n", res)
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

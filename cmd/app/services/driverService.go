package services

import (
	"dispatch-system/models"
	"dispatch-system/repositories"
	"fmt"
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

func (t *DriverService) NewTripRequest(req *models.TripRequest) (interface{}, error) {
	req.Uuid = uuid.New().String()
	return t.driverRepo.NewRequest(req)
}

// TODO: Refactor these functions which do the same thing
func (t *DriverService) NewTripStatusRequest(req *models.TripQueryRequest) (interface{}, error) {
	return t.driverRepo.NewRequest(req)
}

func (t *DriverService) NewDriverInfo(info *models.DriverInfo) {
	t.driverRepo.NewRequest(info)
}

func (t *DriverService) NewDriverUpdate(update *models.DriverUpdate) {
	t.driverRepo.NewRequest(update)
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

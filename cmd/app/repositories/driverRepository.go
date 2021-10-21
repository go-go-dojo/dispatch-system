package repositories

import (
	"daitan-dispatch-system/cmd/app/models"
	"daitan-dispatch-system/cmd/app/utils"
	"errors"
	"fmt"
	"github.com/google/uuid"
)

type DriverRepository struct {
	drivers map[string]*models.DriverInfo

	requestCh  chan interface{}
	ResponseCh chan interface{}
}

func (s *DriverRepository) Init() {
	if s.drivers == nil {
		s.drivers = make(map[string]*models.DriverInfo)
		s.requestCh = make(chan interface{})
		s.ResponseCh = make(chan interface{})
		go s.handleRequestChannel()
	}
}

func (s *DriverRepository) Shutdown() {
	// Close request and response channel
	fmt.Println("[DriverRepository.Shutdown] Begin")
	close(s.requestCh)
	close(s.ResponseCh)
}

func (s *DriverRepository) NewRequest(obj interface{}) {
	s.requestCh <- obj
}

func (s *DriverRepository) handleRequestChannel() {
	for req := range s.requestCh {
		switch req.(type) {
		case *models.TripRequest:
			// Find trip
			s.ProcessTripRequest(req.(*models.TripRequest))
			// Take result and insert into channel
		case *models.DriverInfo:
			// Update or add driver
			s.ProcessDriverInfo(req.(*models.DriverInfo))
		case *models.DriverUpdate:
			// Periodic driver location/status update
			s.ProcessDriverUpdate(req.(*models.DriverUpdate))
		case *models.QueryRequest:
			// Request driver info
			s.ProcessDriverQuery(req.(*models.QueryRequest))
		default:
			panic("[DriverRepository.handleRequestChannel] Unrecognized type")
		}
	}
}

func (s *DriverRepository) ProcessDriverQuery(query *models.QueryRequest) {

	if driver, ok := s.drivers[query.Uuid]; ok {
		fmt.Printf("[DriverRepository.ProcessDriverQuery] Driver info=%s\n", driver)
	} else {
		fmt.Printf("[DriverRepository.ProcessDriverQuery] Could not find driver info for id=%s\n", query.Uuid)
	}
}

func (s *DriverRepository) ProcessDriverUpdate(update *models.DriverUpdate) {

	if driver, ok := s.drivers[update.Uuid]; ok {
		driver.Update(*update)
	} else {
		fmt.Printf("[DriverRepository.ProcessDriverUpdate] Unknwon driver id=%s\n", update.Uuid)
	}
}

func (s *DriverRepository) ProcessDriverInfo(newDriver *models.DriverInfo) {

	if driver, ok := s.drivers[newDriver.Uuid]; ok && (newDriver.Uuid != "") {
		// Due to this simplification, new drivers are able to override other driver's information
		fmt.Printf("[DriverRepository.ProcessDriverInfo] DriverInfo updated, %s\n", *driver)
		s.drivers[newDriver.Uuid] = newDriver
	} else {
		// New driver, generate unique id
		if newDriver.Uuid == "" {
			newDriver.Uuid = uuid.New().String()
		}
		s.drivers[newDriver.Uuid] = newDriver
		fmt.Printf("[DriverRepository.ProcessDriverInfo] New driver registered, %s\n", *s.drivers[newDriver.Uuid])
	}
}

func (s *DriverRepository) ProcessTripRequest(req *models.TripRequest) error {

	var dist, lowestDist float64
	var closestDriver *models.DriverInfo

	// TODO: Use a parallel for loop to find the closest driver
	first := true
	for _, driver := range s.drivers {
		if driver.IsAvailable() {
			dist = utils.DistanceBetween(req.Location, driver.Location)
			fmt.Printf("[DriverRepo.ProcessTripRequest] Assessing distance=%.3f, driver=%s\n", dist, *driver)
			if dist < lowestDist || first {
				first = false
				lowestDist = dist
				closestDriver = driver
			}
		}
	}

	if closestDriver != nil {
		fmt.Printf("[DriverRepository.ProcessTripRequest] Found driver=%s at distance=%.2f for request=%s\n", *closestDriver, lowestDist, *req)
		closestDriver.Status = models.ON_TRIP
		s.ResponseCh <- closestDriver
		return nil
	} else {
		return errors.New("drivers unavailable")
	}
}

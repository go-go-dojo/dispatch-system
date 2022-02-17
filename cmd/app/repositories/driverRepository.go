package repositories

import (
	"dispatch-system/models"
	"dispatch-system/utils"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"log"
	"reflect"
	"sync"
)

func (s *DriverRepository) ProcessTripRequest(tripRequest *models.TripRequest) (*models.Trip, error) {

	driver, err := s.FindClosestDriver(tripRequest)
	if err != nil {
		log.Printf("[TripRequest.ProcessPayload] Error=%s\n", err.Error())
		return &models.Trip{}, errors.New("could not find closes driver for request")
	}

	if driver == nil {
		return &models.Trip{}, errors.New("no driver available")
	}

	driver.Status = models.ON_TRIP

	t := &models.Trip{
		Location: tripRequest.Location,
		Uuid:     uuid.New().String(),
		Status:   models.Assigned,
		Driver:   driver,
	}
	s.trips[t.Uuid] = t
	return t, nil
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
		fmt.Printf("[DriverRepository.ProcessDriverInfo] DriverInfo updated, %v\n", *driver)
		s.drivers[newDriver.Uuid] = newDriver
	} else {
		// New driver, generate unique id
		if newDriver.Uuid == "" {
			newDriver.Uuid = uuid.New().String()
		}
		s.drivers[newDriver.Uuid] = newDriver
		fmt.Printf("[DriverRepository.ProcessDriverInfo] New driver registered, %v\n", *s.drivers[newDriver.Uuid])
	}
}

func (s *DriverRepository) ProcessTripQueryRequest(query *models.TripQueryRequest) (*models.Trip, error) {
	if trip, ok := s.trips[query.Uuid]; ok {
		fmt.Printf("[DriverRepository.ProcessTripQueryRequest] Tryp info=%v\n", trip)
		return trip, nil
	} else {
		fmt.Printf("[DriverRepository.ProcessTripQueryRequest] Could not find driver info for id=%s\n", query.Uuid)
		return &models.Trip{}, errors.New(fmt.Sprintf("could not find the trip %s for uuid", query.Uuid))
	}
}

type DriverRepository struct {
	drivers map[string]*models.DriverInfo
	trips   map[string]*models.Trip

	ResponseCh chan interface{}

	repositoryMutex sync.Mutex
}

func (s *DriverRepository) Init() {
	if s.drivers == nil {
		s.drivers = make(map[string]*models.DriverInfo)
		s.trips = make(map[string]*models.Trip)
		s.ResponseCh = make(chan interface{})
	}
}

func (s *DriverRepository) Shutdown() {
	// Close request and response channel
	fmt.Println("[DriverRepository.Shutdown] Begin")
	close(s.ResponseCh)
}

func (s *DriverRepository) NewRequest(msg interface{}) (interface{}, error) {
	s.repositoryMutex.Lock()
	res, err := s.handleRequest(msg)
	s.repositoryMutex.Unlock()
	return res, err
}

func (s *DriverRepository) handleRequest(req interface{}) (interface{}, error) {

	switch req.(type) {
	case *models.TripRequest:
		return s.ProcessTripRequest(req.(*models.TripRequest))
	case *models.DriverUpdate:
		s.ProcessDriverUpdate(req.(*models.DriverUpdate))
	case *models.DriverInfo:
		s.ProcessDriverInfo(req.(*models.DriverInfo))
	case *models.DriverQueryRequest:
		s.ProcessDriverQuery(req.(*models.DriverQueryRequest))
	case *models.TripQueryRequest:
		return s.ProcessTripQueryRequest(req.(*models.TripQueryRequest))
	default:
		log.Printf("[DriverRepository.handleRequest] invalid struct type %s\n", reflect.TypeOf(req))
	}

	return nil, nil
}

func (s *DriverRepository) ProcessDriverQuery(query *models.DriverQueryRequest) {
	if driver, ok := s.drivers[query.Uuid]; ok {
		fmt.Printf("[DriverRepository.ProcessDriverQuery] Driver info=%v\n", driver)
	} else {
		fmt.Printf("[DriverRepository.ProcessDriverQuery] Could not find driver info for id=%s\n", query.Uuid)
	}
}

func (s *DriverRepository) FindClosestDriver(req *models.TripRequest) (*models.DriverInfo, error) {
	var dist, lowestDist float64
	var closestDriver *models.DriverInfo

	// TODO: Use a parallel for loop to find the closest driver
	first := true
	for _, driver := range s.drivers {
		if driver.IsAvailable() {
			dist = utils.DistanceBetween(req.Location, driver.Location)
			fmt.Printf("[DriverRepo.FindClosestDriver] Assessing distance=%.3f, driver=%v\n", dist, *driver)
			if dist < lowestDist || first {
				first = false
				lowestDist = dist
				closestDriver = driver
			}
		}
	}

	if closestDriver != nil {
		fmt.Printf("[DriverRepository.FindClosestDriver] Found driver=%v at distance=%.2f for request=%v\n", *closestDriver, lowestDist, *req)
		return closestDriver, nil
	} else {
		return nil, errors.New("drivers unavailable")
	}
}

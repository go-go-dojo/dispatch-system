package repositories

import (
	"dispatch-system/models"
	"dispatch-system/utils"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"log"
	"reflect"
)

type Message struct {
	MsgType reflect.Type
	Payload interface{}
}

type IService interface {
	ProcessPayload(payload interface{}, s *DriverRepository)
}

type TripRequestType struct {
}

type DriverUpdateType struct {
}

type DriverInfoType struct {
}

type DriverQueryType struct {
}

func (t *TripRequestType) ProcessPayload(payload interface{}, s *DriverRepository) {

	tripRequest, ok := payload.(*models.TripRequest)

	if !ok {
		log.Printf("[TripRequestType.ProcessPayload] Wrong struct type!")
		// TODO: Add error to the response channel
		return
	}

	driver, err := s.FindClosestDriver(tripRequest)
	if err != nil {
		log.Printf("[TripRequest.ProcessPayload] Error=%s\n", err.Error())
	}
	if driver != nil {
		t := &models.Trip{
			Location: tripRequest.Location,
			Uuid:     uuid.New().String(),
			Status:   models.Assigned,
			Driver:   driver,
		}

		s.ResponseCh <- t
		s.trips[t.Uuid] = t

		// TODO: This is not ideal, the response should not be handled by driverRepository
		body, err := json.Marshal(*t)
		if err != nil {
			log.Printf("[TripRequest.ProcessPayload] Error marshalling trip")
		}
		tripRequest.Writer.Header().Set("Content-Type", "application/json")
		//tripRequest.Writer.WriteHeader(http.StatusOK)
		tripRequest.Writer.Write(body)
	}
}

func (t *DriverUpdateType) ProcessPayload(payload interface{}, s *DriverRepository) {
	update := payload.(*models.DriverUpdate)
	if driver, ok := s.drivers[update.Uuid]; ok {
		driver.Update(*update)
	} else {
		fmt.Printf("[DriverRepository.ProcessDriverUpdate] Unknwon driver id=%s\n", update.Uuid)
	}
}

func (t *DriverInfoType) ProcessPayload(payload interface{}, s *DriverRepository) {
	newDriver := payload.(*models.DriverInfo)
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

func (t *DriverQueryType) ProcessPayload(payload interface{}, s *DriverRepository) {
	query := payload.(models.QueryRequest)
	if driver, ok := s.drivers[query.Uuid]; ok {
		fmt.Printf("[DriverRepository.ProcessDriverQuery] Driver info=%s\n", driver)
	} else {
		fmt.Printf("[DriverRepository.ProcessDriverQuery] Could not find driver info for id=%s\n", query.Uuid)
	}
}

type DriverRepository struct {
	drivers map[string]*models.DriverInfo
	trips   map[string]*models.Trip

	handles map[string]IService

	requestCh  chan *Message
	ResponseCh chan interface{}
}

func (s *DriverRepository) Init() {
	if s.drivers == nil {
		s.drivers = make(map[string]*models.DriverInfo)
		s.trips = make(map[string]*models.Trip)
		s.requestCh = make(chan *Message)
		s.ResponseCh = make(chan interface{})
		s.handles = make(map[string]IService)
	}

	tripRequestType := &TripRequestType{}
	if err := s.RegisterService(tripRequestType); err != nil {
		fmt.Printf("Error at TripRequestType: %s", err.Error())
	}

	driverUpdateType := &DriverUpdateType{}
	if err := s.RegisterService(driverUpdateType); err != nil {
		fmt.Printf("Error at DriverUpdateType: %s", err.Error())
	}

	driverInfoType := &DriverInfoType{}
	if err := s.RegisterService(driverInfoType); err != nil {
		fmt.Printf("Error at DriverInfoType: %s", err.Error())
	}

	driverQueryType := &DriverQueryType{}
	if err := s.RegisterService(driverQueryType); err != nil {
		fmt.Printf("Error at DriverQueryType: %s", err.Error())
	}

}

func (s *DriverRepository) HandleRequestChannel() {
	for req := range s.requestCh {
		s.handleRequest(req)
	}
}

func (s *DriverRepository) Shutdown() {
	// Close request and response channel
	fmt.Println("[DriverRepository.Shutdown] Begin")
	close(s.requestCh)
	close(s.ResponseCh)
}

func (s *DriverRepository) RegisterService(service IService) error {

	name := reflect.TypeOf(service).String()
	if _, exists := s.handles[name]; exists {
		return fmt.Errorf("service already exists %s", name)
	}
	s.handles[name] = service
	return nil
}

func (s *DriverRepository) NewRequest(msg *Message) {
	s.requestCh <- msg
}

func (s *DriverRepository) handleRequest(req *Message) {

	if svc, ok := s.handles[req.MsgType.String()]; ok {
		svc.ProcessPayload(req.Payload, s)
	} else {
		fmt.Printf("Error, MsgType=%s not found\n", req.MsgType)
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
			fmt.Printf("[DriverRepo.FindClosestDriver] Assessing distance=%.3f, driver=%s\n", dist, *driver)
			if dist < lowestDist || first {
				first = false
				lowestDist = dist
				closestDriver = driver
			}
		}
	}

	if closestDriver != nil {
		fmt.Printf("[DriverRepository.FindClosestDriver] Found driver=%s at distance=%.2f for request=%s\n", *closestDriver, lowestDist, *req)
		return closestDriver, nil
	} else {
		return nil, errors.New("drivers unavailable")
	}
}

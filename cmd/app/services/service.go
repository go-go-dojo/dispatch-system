package services

import (
	"fmt"
	"log"
	"math"
	"sync"

	"github.com/google/uuid"

	"daitan-dispatch-system/cmd/app/models"
)

type Service struct {
	models.TripRequest
	models.Trip
	models.Location

	tripRequestCh chan *models.TripRequest
}

var instance *Service
var once sync.Once

func GetInstance() *Service {
	once.Do(func() {
		instance = &Service{}
		instance.startTripFinder()
	})

	return instance
}

func (t *Service) NewTripRequest(request *models.TripRequest) error {

	t.tripRequestCh <- request
	fmt.Printf("[NewTripRequest] Trip request added to channel, %s\n", *request)
	return nil
}

func (t *Service) RequestTrip(request *models.TripRequest) (models.Trip, error) {

	userLocation := request.Location
	_uuid := uuid.New().String()
	driver, _ := t.FindMockDriver(&userLocation)
	location := models.Location{
		Latitude:  -15.838399,
		Longitude: -48.010947,
	}

	return models.Trip{
		Uuid:     _uuid,
		Driver:   driver,
		Status:   2,
		Location: location,
	}, nil

}

func (trip *Service) FindTrip(uuid string) (models.Trip, error) {

	driver := models.Driver{
		Name:    "Agatha Ayla",
		Ranking: 4.9,
		Trips:   1212,
		Car: models.Car{
			PlateNumber: "GJX-2343",
			Brand:       "Lamborghini",
			Model:       "aventador-svj",
		},
	}

	location := models.Location{
		Latitude:  -15.838399,
		Longitude: -48.010947,
	}

	return models.Trip{
		Uuid:     uuid,
		Driver:   driver,
		Status:   2,
		Location: location,
	}, nil
}

func (t *Service) FindMockDriver(location *models.Location) (models.Driver, error) {

	if location == nil {

	}

	return models.Driver{
		Uuid:    "ad716768-3a58-4a5a-aa40-cccda2d5b718",
		Name:    "Agatha Ayla",
		Ranking: 4.9,
		Trips:   1212,
		Car: models.Car{
			PlateNumber: "GJX-2343",
			Brand:       "Lamborghini",
			Model:       "aventador-svj",
		},
	}, nil
}

func add(from, to models.Location) float64 {
	const PI float64 = 3.141592653589793

	radlat1 := float64(PI * from.Latitude / 180)
	radlat2 := float64(PI * to.Latitude / 180)

	theta := float64(from.Longitude - to.Longitude)
	radtheta := float64(PI * theta / 180)

	dist := math.Sin(radlat1)*math.Sin(radlat2) + math.Cos(radlat1)*math.Cos(radlat2)*math.Cos(radtheta)

	if dist > 1 {
		dist = 1
	}

	dist = math.Acos(dist)
	dist = dist * 180 / PI
	dist = dist * 60 * 1.1515
	dist = dist * 1.609344 * 1000

	return dist
}

func (t *Service) startTripFinder() {

	log.Println("[startTripFinder] Started")
	t.tripRequestCh = make(chan *models.TripRequest, 4)
	go func() {
		for req := range t.tripRequestCh {
			newTrip, err := t.RequestTrip(req)
			if err != nil {
				fmt.Printf("[go startTripFinder] Error requesting trip for req=%s\n", *req)
				// TODO: Notify someone
				// TODO: Retry mechanism
				continue
			}

			fmt.Printf("[go startTripFinder] Found trip=%s for request=%s\n", newTrip, *req)
			// TODO: Do something with the trip (?)
		}
		fmt.Println("[startTripFinder] goroutine end")
	}()
}

func (t *Service) Shutdown() {
	close(t.tripRequestCh)
	fmt.Println("Service shutdown, trip request channel closed")
}

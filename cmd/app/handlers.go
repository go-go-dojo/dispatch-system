package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"daitan-dispatch-system/cmd/app/models"
	"daitan-dispatch-system/cmd/app/services"

	"github.com/gorilla/mux"
)

func (app *application) AddNewDriver(w http.ResponseWriter, r *http.Request) {
	var d *models.Driver
	err := json.NewDecoder(r.Body).Decode(&d)
	if err != nil {
		app.serverError(w, err)
	}

	driver, err := services.GetDriverService().AddNewDriver(d)
	if err != nil {
		app.serverError(w, err)
	}

	body, err := json.Marshal(driver)
	if err != nil {
		app.serverError(w, err)
	}

	app.info.Println("Driver included")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (app *application) AddNewLocation(w http.ResponseWriter, r *http.Request) {
	var l *models.Location
	err := json.NewDecoder(r.Body).Decode(&l)
	if err != nil {
		app.serverError(w, err)
	}

	vars := mux.Vars(r)
	uuid := vars["uuid"]
	location, err := services.GetDriverService().AddNewLocation(uuid, l)
	if err != nil {
		app.serverError(w, err)
	}

	body, err := json.Marshal(location)
	if err != nil {
		app.serverError(w, err)
	}

	app.info.Println("Location created")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (app *application) NewDriverStatus(w http.ResponseWriter, r *http.Request) {
	var d *models.Driver
	err := json.NewDecoder(r.Body).Decode(&d)
	if err != nil {
		app.serverError(w, err)
	}

	vars := mux.Vars(r)
	uuid := vars["uuid"]
	driver, err := services.GetDriverService().UpdateDriverStatus(uuid, &d.Status)
	if err != nil {
		app.serverError(w, err)
	}

	body, err := json.Marshal(driver)
	if err != nil {
		app.serverError(w, err)
	}

	app.info.Println("Location created")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (app *application) FindDriver(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uuid := vars["uuid"]
	driver, err := services.GetDriverService().FindDriver(uuid)
	if err != nil {
		app.serverError(w, err)
	}

	body, err := json.Marshal(driver)
	if err != nil {
		app.serverError(w, err)
	}

	app.info.Println("Location created")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}
func (app *application) requestTrip(w http.ResponseWriter, r *http.Request) {

	var u *models.TripRequest
	// TODO: Decode does not return errror when the parameters in the json do not correspond to the TripRequest struct
	err := json.NewDecoder(r.Body).Decode(&u)

	err = services.GetInstance().NewTripRequest(u)
	if err != nil {
		app.serverError(w, err)
	}

	// body, err := json.Marshal(trip)
	// if err != nil {
	// 	app.serverError(w, err)
	// }

	// app.info.Println("Trip created")

	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusCreated)
	// w.Write(body)
}

func (app *application) findTrip(w http.ResponseWriter, r *http.Request) {

	var u models.Trip
	vars := mux.Vars(r)
	uuid := vars["uuid"]
	fmt.Printf("[findTrip] Read uuid=%s\n", uuid)
	trip, err := services.GetInstance().FindTrip(uuid)
	if err != nil {
		app.serverError(w, err)
	}
	err = json.NewDecoder(r.Body).Decode(&u)
	body, err := json.Marshal(trip)
	if err != nil {
		app.serverError(w, err)
	}

	app.info.Println("Trip was found")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

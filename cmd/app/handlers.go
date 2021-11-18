package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"dispatch-system/models"
	"dispatch-system/services"

	"github.com/gorilla/mux"
)

func (app *application) updateDriverInfo(w http.ResponseWriter, r *http.Request) {

	var driver models.DriverInfo

	if err := json.NewDecoder(r.Body).Decode(&driver); err != nil {
		app.serverError(w, errors.New("error decoding driver info update"))
		return
	}

	services.GetDriverService().NewDriverInfo(&driver)
}

func (app *application) updateDriver(w http.ResponseWriter, r *http.Request) {

	var update models.DriverUpdate

	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		app.serverError(w, errors.New("error decoding driver location/status update"))
		return
	}

	services.GetDriverService().NewDriverUpdate(&update)
}

func (app *application) AddNewDriver(w http.ResponseWriter, r *http.Request) {
	var d *models.DriverInfo
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

	app.info.Println("DriverInfo included")

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
	var d *models.DriverInfo
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

func (app *application) findDriver(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uuid := vars["uuid"]
	query := models.QueryRequest{Uuid: uuid, Type: models.DRIVERINFO}
	services.GetDriverService().NewDriverQuery(&query)

	//w.Header().Set("Content-Type", "application/json")
	//w.WriteHeader(http.StatusOK)
	//w.Write(body)
}
func (app *application) requestTrip(w http.ResponseWriter, r *http.Request) {

	var req models.TripRequest
	// TODO: Decode does not return error when the parameters in the json do not correspond to the TripRequest struct
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		app.serverError(w, err)
		return
	}

	fmt.Println("[application.requestTrip] Created new request ", req)
	services.GetDriverService().NewTripRequest(&req)
}

func (app *application) findTrip(w http.ResponseWriter, r *http.Request) {

	panic("[application.findTrip] Not yet implemented")
	// This method might not be used because trips status updates will be automatically and periodically sent by the server to the clients
	/*
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
	*/
}

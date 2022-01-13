package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"dispatch-system/models"
	"dispatch-system/services"
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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	//w.Write()
}

func (app *application) requestTrip(w http.ResponseWriter, r *http.Request) {

	var req models.TripRequest
	// TODO: Decode does not return error when the parameters in the json do not correspond to the TripRequest struct
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		app.serverError(w, err)
		return
	}

	req.Writer = w
	fmt.Println("[application.requestTrip] Created new request ", req)
	services.GetDriverService().NewTripRequest(&req)
}

func (app *application) findTrip(w http.ResponseWriter, r *http.Request) {

	panic("[application.findTrip] Not yet implemented")
}

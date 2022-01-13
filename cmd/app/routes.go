package main

import (
	"github.com/gorilla/mux"
)

func (app *application) Routes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/trips", app.requestTrip).Methods("POST")
	r.HandleFunc("/api/driver/update", app.updateDriver).Methods("POST")
	r.HandleFunc("/api/driver/updateInfo", app.updateDriverInfo).Methods("POST")
	r.HandleFunc("/api/trips/{uuid}", app.findTrip).Methods("GET")

	return r
}

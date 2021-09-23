package main

import (
	"github.com/gorilla/mux"
)

func (app *application) Routes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/trips/", app.requestTrip).Methods("POST")
	r.HandleFunc("/api/trips/{uuid}", app.findTrip).Methods("GET")

	r.HandleFunc("/api/driver", app.AddNewDriver).Methods("POST")
	r.HandleFunc("/api/driver/{uuid}", app.FindDriver).Methods("GET")
	r.HandleFunc("/api/driver/{uuid}/status", app.NewDriverStatus).Methods("PATCH")
	r.HandleFunc("/api/driver/{uuid}/locations", app.AddNewLocation).Methods("POST")
	return r
}

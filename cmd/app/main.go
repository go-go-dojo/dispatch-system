package main

import (
	"dispatch-system/services"
	"flag"
	"github.com/labstack/echo/v4"
	"log"
)

func main() {

	port := flag.String("port", "8089", "HTTP server network port")

	flag.Parse()

	e := echo.New()
	e.POST("/api/trips", requestTrip)
	e.POST("/api/driver/update", updateDriver)
	e.POST("/api/driver/setInfo", setDriverInfo)
	e.GET("/api/trips/:uuid", findTrip)

	log.Printf("Starting server on %s", *port)

	e.Logger.Fatal(e.Start("localhost:8089"))
	services.GetDriverService().Shutdown()
}

package main

import (
	"dispatch-system/models"
	"dispatch-system/services"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
)

func updateDriverInfo(c echo.Context) error {

	var driver models.DriverInfo

	if err := json.NewDecoder(c.Request().Body).Decode(&driver); err != nil {
		return errors.New("error decoding driver info update")
	}

	services.GetDriverService().NewDriverInfo(&driver)
	return nil
}

func updateDriver(c echo.Context) error {

	var update models.DriverUpdate

	if err := json.NewDecoder(c.Request().Body).Decode(&update); err != nil {
		return errors.New("error decoding driver location/status update")
	}

	services.GetDriverService().NewDriverUpdate(&update)
	return nil
}

func requestTrip(c echo.Context) error {

	var req models.TripRequest
	// TODO: Decode does not return error when the parameters in the json do not correspond to the TripRequest struct
	if err := json.NewDecoder(c.Request().Body).Decode(&req); err != nil {
		return err
	}
	req.Context = c

	fmt.Println("[application.requestTrip] Created new request ", req)
	services.GetDriverService().NewTripRequest(&req)
	return nil
}

func findTrip(c echo.Context) error {

	panic("[application.findTrip] Not yet implemented")
	return nil
}

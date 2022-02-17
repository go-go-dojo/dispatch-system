package main

import (
	"dispatch-system/models"
	"dispatch-system/services"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func setDriverInfo(c echo.Context) error {

	var driver models.DriverInfo

	if err := json.NewDecoder(c.Request().Body).Decode(&driver); err != nil {
		return errors.New("error decoding driver info update")
	}

	updatedDriver, err := services.GetDriverService().NewDriverInfo(&driver)
	if err != nil {
		return c.String(http.StatusOK, err.Error())
	}

	return c.JSON(http.StatusOK, updatedDriver)
}

func updateDriver(c echo.Context) error {

	var update models.DriverUpdate

	if err := json.NewDecoder(c.Request().Body).Decode(&update); err != nil {
		return errors.New("error decoding driver location/status update")
	}

	updatedDriver, err := services.GetDriverService().NewDriverUpdate(&update)
	if err != nil {
		return c.String(http.StatusOK, err.Error())
	}

	return c.JSON(http.StatusOK, updatedDriver)
}

func requestTrip(c echo.Context) error {

	var req models.TripRequest
	// TODO: Decode does not return error when the parameters in the json do not correspond to the TripRequest struct
	if err := json.NewDecoder(c.Request().Body).Decode(&req); err != nil {
		return err
	}

	fmt.Println("[application.requestTrip] Created new request ", req)
	res, err := services.GetDriverService().NewTripRequest(&req)

	if err != nil {
		return c.String(http.StatusOK, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}

func findTrip(c echo.Context) error {

	tripId := c.Param("uuid")
	log.Printf("[findTrip] uuid: %s\n", tripId)

	tripQuery := &models.TripQueryRequest{Uuid: tripId}

	res, err := services.GetDriverService().NewTripStatusRequest(tripQuery)

	if err != nil {
		return c.String(http.StatusOK, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}

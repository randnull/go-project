package handlers

import (
	"net/http"
	"project/internal/driver/service"
)

type DriverHandler struct {
	driver *service.Driver
}

func NewHandler(driver *service.Driver) *DriverHandler {
	return &DriverHandler{
		driver: driver,
	}
}

func (d *DriverHandler) GetAllTripHandler(w http.ResponseWriter, r *http.Request) {

}

package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"project/internal/location/errors"
	"project/internal/location/repository"
	"strconv"
)

type LatLngLiteral struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type LocationHandler struct {
	location *repository.LocationRepository
}

func NewHandler(location *repository.LocationRepository) *LocationHandler {
	return &LocationHandler{
		location: location,
	}
}

func (lhandler *LocationHandler) GetDriversHandler(w http.ResponseWriter, r *http.Request) {
	drivers, err := lhandler.location.GetDrivers()
	if err != nil {
		http.Error(w, errors.DriversNotFound.Error(), 404)
		return
	}
	json.NewEncoder(w).Encode(drivers)
}

func (lhandler *LocationHandler) UpdateDriverLocationHandler(w http.ResponseWriter, r *http.Request) {
	driverID, ok := mux.Vars(r)["driver_id"]
	if !ok {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	intDriverID, err := strconv.Atoi(driverID)
	if err != nil {
		http.Error(w, "Bad request - driverID is not an integer", http.StatusBadRequest)
		return
	}
	var locReq LatLngLiteral
	err = json.NewDecoder(r.Body).Decode(&locReq)
	if err != nil {
		http.Error(w, "errors.InvalidLocation.Error()", 400)
		return
	}
	driver, err := lhandler.location.UpdateDriverPosition(intDriverID, locReq.Lat, locReq.Lng)
	if err != nil {
		http.Error(w, "Invalid server error", 500)
		return
	}
	json.NewEncoder(w).Encode(driver)
}

package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"math"
	"net/http"
	"project/internal/location/errors"
	"project/internal/location/service"
	"project/modals"
)

type LatLngLiteral struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type LocationRequest struct {
	Lat    float64 `json:"lat"`
	Lng    float64 `json:"lng"`
	Radius float64 `json:"radius"`
}

type LocationHandler struct {
	location *service.Location
}

func NewHandler(location *service.Location) *LocationHandler {
	return &LocationHandler{
		location: location,
	}
}

//func (lhandler *LocationHandler) GetDriversHandler(w http.ResponseWriter, r *http.Request) {
//	drivers, err := lhandler.location.GetAllDrivers()
//	if err != nil {
//		http.Error(w, errors.DriversNotFound.Error(), 404)
//		return
//	}
//	json.NewEncoder(w).Encode(drivers)
//	w.Header().Set("Content-Type", "application/json")
//	w.WriteHeader(http.StatusOK)
//}

func calculateDistance(lat1 float64, lon1 float64, lat2 float64, lon2 float64) float64 {
	return math.Sqrt(math.Pow((lat2-lat1), 2) + math.Pow((lon2-lon1), 2))
}

func (lhandler *LocationHandler) GetDriversHandler(w http.ResponseWriter, r *http.Request) {
	var locReq LocationRequest

	err := json.NewDecoder(r.Body).Decode(&locReq)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	drivers, err := lhandler.location.GetAllDrivers()
	if err != nil {
		http.Error(w, errors.DriversNotFound.Error(), 404)
		return
	}

	var driversInRadius []modals.Driver
	for _, driver := range drivers {
		distance := calculateDistance(locReq.Lat, locReq.Lng, driver.Lat, driver.Lng)
		if distance <= locReq.Radius {
			driversInRadius = append(driversInRadius, driver)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(driversInRadius)
	w.WriteHeader(http.StatusOK)
}

func (lhandler *LocationHandler) UpdateDriverLocationHandler(w http.ResponseWriter, r *http.Request) {
	driverID, ok := mux.Vars(r)["driver_id"]
	if !ok {
		http.Error(w, errors.InvalidDriverId.Error(), http.StatusBadRequest)
		return
	}

	var locReq LatLngLiteral
	err := json.NewDecoder(r.Body).Decode(&locReq)
	if err != nil {
		http.Error(w, errors.InvalidLocation.Error(), 400)
		return
	}
	driver, err := lhandler.location.UpdateDriverPosition(driverID, locReq.Lat, locReq.Lng)
	if err != nil {
		http.Error(w, errors.FailedToUpdatePosition.Error(), 500)
		return
	}
	json.NewEncoder(w).Encode(driver)
	w.WriteHeader(http.StatusOK)
}

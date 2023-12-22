package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/opentracing/opentracing-go"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"log"
	"math"
	"net/http"
	"project/internal/location/errors"
	"project/internal/location/service"
	"project/modals"
)

var (
	getDrivers_allRequests = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "counters_getDrivers", Name: "allRequests", Help: "GetDrivers all requests counter",
	},
	)
	getDrivers_successfulRequests = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "counters_getDrivers", Name: "successfulRequest", Help: "GetDrivers successful requests counter",
	},
	)
	updateDriverLocation_allRequests = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "counters_updateDriverLocation", Name: "allRequests", Help: "UpdateDriverLocation all requests counter",
	},
	)
	updateDriverLocation_successfulRequests = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "counters_updateDriverLocation", Name: "successfulRequest", Help: "UpdateDriverLocation successful requests counter",
	},
	)
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

func calculateDistance(lat1 float64, lon1 float64, lat2 float64, lon2 float64) float64 {
	return math.Sqrt(math.Pow((lat2-lat1), 2) + math.Pow((lon2-lon1), 2))
}

func (lhandler *LocationHandler) GetDriversHandler(w http.ResponseWriter, r *http.Request) {
	span := opentracing.StartSpan("GetDriversHandler")
	defer span.Finish()

	getDrivers_allRequests.Inc()

	var locReq LocationRequest

	err := json.NewDecoder(r.Body).Decode(&locReq)
	if err != nil {
		http.Error(w, errors.InvalidRequestBody.Error(), http.StatusBadRequest)
		return
	}

	drivers, err := lhandler.location.GetAllDrivers()
	if err != nil {
		http.Error(w, errors.DriversNotFound.Error(), http.StatusNotFound)
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
	getDrivers_successfulRequests.Inc()
	log.Println("GetDriversHandler - success")
}

func (lhandler *LocationHandler) UpdateDriverLocationHandler(w http.ResponseWriter, r *http.Request) {
	span := opentracing.StartSpan("UpdateDriverLocationHandler")
	defer span.Finish()

	updateDriverLocation_allRequests.Inc()

	driverID, ok := mux.Vars(r)["driver_id"]
	if !ok {
		http.Error(w, errors.InvalidDriverId.Error(), http.StatusBadRequest)
		fmt.Println(errors.InvalidDriverId.Error())
		return
	}

	var locReq LatLngLiteral
	err := json.NewDecoder(r.Body).Decode(&locReq)
	if err != nil {
		http.Error(w, errors.InvalidLocation.Error(), 400)
		fmt.Println(errors.InvalidLocation.Error())
		return
	}
	driver, err := lhandler.location.UpdateDriverPosition(driverID, locReq.Lat, locReq.Lng)
	if err != nil {
		http.Error(w, errors.FailedToUpdatePosition.Error(), 500)
		fmt.Println(errors.FailedToUpdatePosition.Error())
		return
	}
	json.NewEncoder(w).Encode(driver)
	w.WriteHeader(http.StatusOK)
	updateDriverLocation_successfulRequests.Inc()
	log.Println("UpdateDriverLocationHandler - success")
}

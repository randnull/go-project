package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"project/internal/driver/errors"
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

func (dhandler *DriverHandler) GetAllTripHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("user_id")

	trips, err := dhandler.driver.GetAllTrips(userID)
	if err != nil {
		http.Error(w, errors.FailedToGetTripsList.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(trips)
	w.WriteHeader(http.StatusOK)
}

func (dhandler *DriverHandler) GetTripByIdHandler(w http.ResponseWriter, r *http.Request) {
	tripID, ok := mux.Vars(r)["trip_id"]
	if !ok {
		http.Error(w, errors.InvalidTripID.Error(), http.StatusBadRequest)
		return
	}

	userID := r.Header.Get("user_id")

	trip, err := dhandler.driver.GetIdTrip(tripID, userID)
	if err != nil {
		http.Error(w, errors.InvalidTripID.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(trip)
	w.WriteHeader(http.StatusOK)
}

func (dhandler *DriverHandler) AcceptTripHandler(w http.ResponseWriter, r *http.Request) {
	tripID, ok := mux.Vars(r)["trip_id"]
	if !ok {
		http.Error(w, errors.InvalidTripID.Error(), http.StatusBadRequest)
		return
	}

	userID := r.Header.Get("user_id")

	if err := dhandler.driver.Accept(tripID, userID); err != nil {
		http.Error(w, "Failed to accept trip", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (dhandler *DriverHandler) StartTripHandler(w http.ResponseWriter, r *http.Request) {
	tripID, ok := mux.Vars(r)["trip_id"]
	if !ok {
		http.Error(w, "Invalid trip id", http.StatusBadRequest)
		return
	}

	userID := r.Header.Get("user_id")

	if err := dhandler.driver.Start(tripID, userID); err != nil {
		http.Error(w, "Failed to start trip", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (dhandler *DriverHandler) CancelTripHandler(w http.ResponseWriter, r *http.Request) {
	tripID, ok := mux.Vars(r)["trip_id"]
	if !ok {
		http.Error(w, "Invalid trip id", http.StatusBadRequest)
		return
	}

	userID := r.Header.Get("user_id")
	reason := r.URL.Query().Get("reason")

	if err := dhandler.driver.Cancel(tripID, userID, reason); err != nil {
		http.Error(w, "Failed to cancel trip", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (dhandler *DriverHandler) EndTripHandler(w http.ResponseWriter, r *http.Request) {
	tripID, ok := mux.Vars(r)["trip_id"]
	if !ok {
		http.Error(w, "Invalid trip id", http.StatusBadRequest)
		return
	}

	userID := r.Header.Get("user_id")

	if err := dhandler.driver.End(tripID, userID); err != nil {
		http.Error(w, "Failed to end trip", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

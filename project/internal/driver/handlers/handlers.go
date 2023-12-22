package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"project/internal/driver/errors"
	"project/internal/driver/service"
	"project/modals"
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
	fmt.Println("handler get")
	tripID, ok := mux.Vars(r)["trip_id"]
	if !ok {
		http.Error(w, errors.InvalidTripID.Error(), http.StatusBadRequest)
		return
	}

	userID := r.Header.Get("user_id")
	fmt.Println(tripID, userID)
	trip, err := dhandler.driver.GetIdTrip(userID, tripID)
	fmt.Println("handler ask")
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

	if err := dhandler.driver.Accept(userID, tripID); err != nil {
		http.Error(w, "Failed to accept trip", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (dhandler *DriverHandler) StartTripHandler(w http.ResponseWriter, r *http.Request) {
	tripID, ok := mux.Vars(r)["trip_id"]
	fmt.Println("start trip")
	if !ok {
		http.Error(w, "Invalid trip id", http.StatusBadRequest)
		return
	}

	userID := r.Header.Get("user_id")

	if err := dhandler.driver.Start(userID, tripID); err != nil {
		http.Error(w, "Failed to start trip", http.StatusInternalServerError)
		return
	}
	fmt.Println("all ok")
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

	if err := dhandler.driver.Cancel(userID, tripID, reason); err != nil {
		http.Error(w, "Failed to cancel trip", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// /6584149f0b80e9fe6de31ceb/cancel
func (dhandler *DriverHandler) EndTripHandler(w http.ResponseWriter, r *http.Request) {
	tripID, ok := mux.Vars(r)["trip_id"]
	if !ok {
		http.Error(w, "Invalid trip id", http.StatusBadRequest)
		return
	}

	userID := r.Header.Get("user_id")

	if err := dhandler.driver.End(userID, tripID); err != nil {
		http.Error(w, "Failed to end trip", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (dhandler *DriverHandler) PutNewTripHandler(w http.ResponseWriter, r *http.Request) {
	var trip modals.Trip

	err := json.NewDecoder(r.Body).Decode(&trip)
	if err != nil {
		http.Error(w, "Invalid trip", http.StatusBadRequest)
		return
	}

	if err := dhandler.driver.PutNewTrip(trip); err != nil {
		http.Error(w, "Failed to put new trip", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

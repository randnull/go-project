package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"log"
	"net/http"
	"project/internal/driver/errors"
	"project/internal/driver/kafka_producer"
	"project/internal/driver/service"
	"project/modals"
)

var (
	getAllTrip_allRequests = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "counters_getAllTrip", Name: "allRequests", Help: "getAllTrip all requests counter",
	},
	)
	getAllTrip_successfulRequests = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "counters_getAllTrip", Name: "successfulRequest", Help: "getAllTrip successful requests counter",
	},
	)
	getTripById_allRequests = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "counters_getTripById", Name: "allRequests", Help: "getTripById all requests counter",
	},
	)
	getTripById_successfulRequests = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "counters_getTripById", Name: "successfulRequest", Help: "getTripById successful requests counter",
	},
	)
	start_allRequests = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "counters_start", Name: "allRequests", Help: "start all requests counter",
	},
	)
	start_successfulRequests = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "counters_start", Name: "successfulRequest", Help: "start successful requests counter",
	},
	)
	accept_allRequests = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "counters_accept", Name: "allRequests", Help: "accept all requests counter",
	},
	)
	accept_successfulRequests = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "counters_accept", Name: "successfulRequest", Help: "accept successful requests counter",
	},
	)
	cancel_allRequests = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "counters_cancel", Name: "allRequests", Help: "cancel all requests counter",
	},
	)
	cancel_successfulRequests = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "counters_cancel", Name: "successfulRequest", Help: "cancel successful requests counter",
	},
	)
	end_allRequests = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "counters_end", Name: "allRequests", Help: "end all requests counter",
	},
	)
	end_successfulRequests = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "counters_end", Name: "successfulRequest", Help: "end successful requests counter",
	},
	)
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
	getAllTrip_allRequests.Inc()
	userID := r.Header.Get("user_id")

	trips, err := dhandler.driver.GetAllTrips(userID)
	if err != nil {
		http.Error(w, errors.FailedToGetTripsList.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(trips)
	w.WriteHeader(http.StatusOK)
	getAllTrip_successfulRequests.Inc()
	log.Println("GetAllTripHandler - success")
}

func (dhandler *DriverHandler) GetTripByIdHandler(w http.ResponseWriter, r *http.Request) {
	getTripById_allRequests.Inc()
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
	getTripById_successfulRequests.Inc()
	log.Println("GetTripByIdHandler - success")
}

func (dhandler *DriverHandler) AcceptTripHandler(w http.ResponseWriter, r *http.Request) {
	accept_allRequests.Inc()
	tripID, ok := mux.Vars(r)["trip_id"]
	if !ok {
		http.Error(w, errors.InvalidTripID.Error(), http.StatusBadRequest)
		return
	}

	userID := r.Header.Get("user_id")

	if err := dhandler.driver.Accept(userID, tripID); err != nil {
		http.Error(w, errors.FailedToAcceptTrip.Error(), http.StatusInternalServerError)
		return
	}

	tripdata, err := dhandler.driver.GetIdTrip(" ", tripID)

	if err != nil {
		log.Fatal(err)
	}

	kafka_producer.Produce_data(tripID, "accept", tripdata)
	err = dhandler.driver.AcceptDriver(tripID)
	if err != nil {
		fmt.Println(err)
	}
	w.WriteHeader(http.StatusOK)
	accept_successfulRequests.Inc()
	log.Println("AcceptTripHandler - success")
}

func (dhandler *DriverHandler) StartTripHandler(w http.ResponseWriter, r *http.Request) {
	start_allRequests.Inc()
	tripID, ok := mux.Vars(r)["trip_id"]
	fmt.Println("start trip")
	if !ok {
		http.Error(w, errors.InvalidTripID.Error(), http.StatusBadRequest)
		return
	}

	userID := r.Header.Get("user_id")

	if err := dhandler.driver.Start(userID, tripID); err != nil {
		http.Error(w, errors.FailedToStartTrip.Error(), http.StatusInternalServerError)
		return
	}

	tripdata, err := dhandler.driver.GetIdTrip(" ", tripID)

	if err != nil {
		log.Fatal(err)
	}
	go func() {
		kafka_producer.Produce_data(tripID, "started", tripdata)
	}()

	fmt.Println("all ok")
	w.WriteHeader(http.StatusOK)
	start_successfulRequests.Inc()
	log.Println("StartTripHandler - success")
}

func (dhandler *DriverHandler) CancelTripHandler(w http.ResponseWriter, r *http.Request) {
	cancel_allRequests.Inc()
	tripID, ok := mux.Vars(r)["trip_id"]
	if !ok {
		http.Error(w, errors.InvalidTripID.Error(), http.StatusBadRequest)
		return
	}

	userID := r.Header.Get("user_id")
	reason := r.URL.Query().Get("reason")

	if err := dhandler.driver.Cancel(userID, tripID, reason); err != nil {
		http.Error(w, errors.FailedToCancelTrip.Error(), http.StatusInternalServerError)
		return
	}

	tripdata, err := dhandler.driver.GetIdTrip(" ", tripID)

	if err != nil {
		log.Fatal(err)
	}

	kafka_producer.Produce_data(tripID, "canceled", tripdata)

	w.WriteHeader(http.StatusOK)
	cancel_successfulRequests.Inc()
	log.Println("CancelTripHandler - success")
}

func (dhandler *DriverHandler) EndTripHandler(w http.ResponseWriter, r *http.Request) {
	end_allRequests.Inc()
	tripID, ok := mux.Vars(r)["trip_id"]
	if !ok {
		http.Error(w, errors.InvalidTripID.Error(), http.StatusBadRequest)
		return
	}

	userID := r.Header.Get("user_id")

	if err := dhandler.driver.End(userID, tripID); err != nil {
		http.Error(w, errors.FailedToEndTrip.Error(), http.StatusInternalServerError)
		return
	}

	tripdata, err := dhandler.driver.GetIdTrip(" ", tripID)

	if err != nil {
		log.Fatal(err)
	}

	kafka_producer.Produce_data(tripID, "ended", tripdata)

	w.WriteHeader(http.StatusOK)
	end_successfulRequests.Inc()
	log.Println("EndTripHandler - success")
}

func (dhandler *DriverHandler) PutNewTripHandler(w http.ResponseWriter, r *http.Request) {
	var trip modals.Trip

	err := json.NewDecoder(r.Body).Decode(&trip)
	if err != nil {
		http.Error(w, errors.InvalidTripID.Error(), http.StatusBadRequest)
		return
	}

	if err := dhandler.driver.PutNewTrip(trip); err != nil {
		http.Error(w, errors.FailedToPutNewTrip.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	log.Println("PutNewTripHandler - success")
}

package driver_app

import (
	"github.com/gorilla/mux"
	"net/http"
	"project/internal/driver/handlers"
	"project/internal/driver/repository"
	"project/internal/driver/service"
)

type App struct {
	repo   *repository.DriverRepository
	server *handlers.DriverHandler
	driver *service.Driver
}

func NewApp() *App {
	repo := repository.NewDriverRepository()
	driv := service.NewDriverService(repo)
	server := handlers.NewHandler(driv)

	apl := &App{
		repo:   repo,
		server: server,
		driver: driv,
	}
	return apl
}

func (a *App) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/trips/", a.server.GetAllTripHandler).Methods("GET")
	router.HandleFunc("/trips/{trip_id}/", a.server.GetTripByIdHandler).Methods("GET")

	router.HandleFunc("/trips/{trip_id}/cancel", a.server.CancelTripHandler).Methods("POST")
	router.HandleFunc("/trips/{trip_id}/accept", a.server.AcceptTripHandler).Methods("POST")
	router.HandleFunc("/trips/{trip_id}/start", a.server.StartTripHandler).Methods("POST")
	router.HandleFunc("/trips/{trip_id}/end", a.server.EndTripHandler).Methods("POST")

	http.ListenAndServe(":1542", router)
}

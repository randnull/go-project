package location_app

import (
	"github.com/gorilla/mux"
	"net/http"
	"project/internal/location/handlers"
	"project/internal/location/repository"
	"project/internal/location/service"
)

type App struct {
	repo        *repository.LocationRepository
	server      *handlers.LocationHandler
	BankService *service.Location
}

func NewApp() *App {
	repo := repository.NewLocationRepository()
	serv := service.NewDriverService(repo)
	server := handlers.NewHandler(serv)

	apl := &App{
		repo:        repo,
		server:      server,
		BankService: serv,
	}
	return apl
}

func (a *App) Run() {
	router := mux.NewRouter()
	router.HandleFunc("/drivers", a.server.GetDriversHandler).Methods("GET")
	router.HandleFunc("/drivers/{driver_id}/location", a.server.UpdateDriverLocationHandler).Methods("POST")

	http.ListenAndServe(":1542", router)
}

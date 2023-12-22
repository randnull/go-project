package location_app

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"project/internal/location/handlers"
	"project/internal/location/repository"
	"project/internal/location/service"
)

type App struct {
	repo     *repository.LocationRepository
	server   *handlers.LocationHandler
	location *service.Location
}

func NewApp() *App {
	repo := repository.NewLocationRepository()
	serv := service.NewLocationService(repo)
	server := handlers.NewHandler(serv)

	apl := &App{
		repo:     repo,
		server:   server,
		location: serv,
	}
	return apl
}

func (a *App) Run() {
	router := mux.NewRouter()
	router.HandleFunc("/drivers", a.server.GetDriversHandler).Methods("GET")
	router.HandleFunc("/drivers/{driver_id}/location", a.server.UpdateDriverLocationHandler).Methods("POST")

	addr := ":1544"
	fmt.Printf("listen %s", addr)
	http.ListenAndServe(addr, router)
}

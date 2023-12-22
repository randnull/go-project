package location_app

import (
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"project/internal/location/handlers"
	"project/internal/location/repository"
	"project/internal/location/service"
)

type App struct {
	repo     *repository.LocationRepository
	server   *handlers.LocationHandler
	location *service.Location
	//closer   io.Closer
}

func NewApp() *App {
	//cfg, _ := config.FromEnv()
	//tracer, closer, _ := cfg.NewTracer(config.Logger(jaeger.StdLogger))
	//opentracing.SetGlobalTracer(tracer)

	repo := repository.NewLocationRepository()
	serv := service.NewLocationService(repo)
	server := handlers.NewHandler(serv)

	apl := &App{
		repo:     repo,
		server:   server,
		location: serv,
		//closer:   closer,
	}
	return apl
}

func (a *App) Run() {
	router := mux.NewRouter()
	router.HandleFunc("/drivers", a.server.GetDriversHandler).Methods("GET")
	router.HandleFunc("/drivers/{driver_id}/location", a.server.UpdateDriverLocationHandler).Methods("POST")
	router.Handle("/metrics", promhttp.Handler())

	addr := ":1515"
	log.Println("Listen on %s\n", addr)
	http.ListenAndServe(addr, router)
}

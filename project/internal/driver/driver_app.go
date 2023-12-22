package driver_app

import (
	"fmt"
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
	repo := repository.NewDriverRepository("mongodb://127.0.0.1:27017")
	driv := service.NewDriverService(repo)
	server := handlers.NewHandler(driv)

	apl := &App{
		repo:   repo,
		server: server,
		driver: driv,
	}
	fmt.Println("application ready!")
	return apl
}

func (a *App) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/trips/", a.server.GetAllTripHandler).Methods("GET")
	router.HandleFunc("/trips/{trip_id}", a.server.GetTripByIdHandler).Methods("GET")

	router.HandleFunc("/trips/{trip_id}/cancel", a.server.CancelTripHandler).Methods("POST")
	router.HandleFunc("/trips/{trip_id}/accept", a.server.AcceptTripHandler).Methods("POST")
	router.HandleFunc("/trips/{trip_id}/start", a.server.StartTripHandler).Methods("POST")
	router.HandleFunc("/trips/{trip_id}/end", a.server.EndTripHandler).Methods("POST")

	router.HandleFunc("/trips/new", a.server.PutNewTripHandler).Methods("POST")

	addr := ":2564"
	fmt.Printf(" listen %s\n", addr)

	//kafka_prod.Cust(a.driver)
	http.ListenAndServe(addr, router)
	//kafka_prod.Cust(a.driver)
	//go func() {
	//	fmt.Println("listen: 2555!")
	//	if err := http.ListenAndServe(":2556", router); err != nil {
	//		log.Fatal(err)
	//	}
	//}()
	//
	//fmt.Println("HTTP server started")
}

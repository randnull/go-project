package location_app

import (
	"github.com/gorilla/mux"
	"net/http"
	"project/internal/location/handlers"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/drivers", handlers.LocationHandler.GetDriversHandler).Methods("GET")
	router.HandleFunc("/drivers/{driver_id}/location", handlers.LocationHandler.UpdateDriverLocationHandler).Methods("POST")

	http.ListenAndServe(":8000", router)
}

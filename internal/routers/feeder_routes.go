package routes

import (
	"feeder-backend/internal/controllers"
  "net/http"

	"github.com/gorilla/mux"
)

func RegisterFeederRoutes(r *mux.Router, controller *controllers.FeederController) {
	// r.HandleFunc("/details", controllers.GetFeedDetails).Methods(http.MethodGet)
  /// r.HandleFunc("/details", controllers.CreateDetail).Methods(http.MethodPost)
	r.HandleFunc("/feeders", controller.Create).Methods(http.MethodPost)
}



package routes

import (
	"feeder-backend/internal/controllers"
  "net/http"

	"github.com/gorilla/mux"
)

func RegisterFeederRoutes(r *mux.Router) {
	r.HandleFunc("/details", controllers.GetFeedDetails).Methods(http.MethodGet)
  r.HandleFunc("/details", controllers.CreateDetail).Methods(http.MethodPost)
}



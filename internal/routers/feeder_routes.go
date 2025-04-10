package routes

import (
	"feeder-backend/internal/controllers"
	"feeder-backend/internal/config"
  "net/http"

	"github.com/gorilla/mux"
)

func RegisterFeederRoutes(r *mux.Router) {
	r.HandleFunc("/details", controllers.GetFeedDetails).Methods(http.MethodGet)
  r.HandleFunc("/details", controllers.CreateDetail).Methods(http.MethodPost)
}

func StartServer(r *mux.Router, s *http.Server) error {
  
  var port = config.ServerConfig["port"]
	s := &http.Server{
	Handler: r,
	Addr: port,
	}

  // s.ListenAndServe()
	fmt.Println("Server started at " + port)
	fmt.Println(s.ListenAndServe())

  err := s.ListenAndServe()
  if err != nil {
    fmt.Println("Server failed start", err)
  }
}

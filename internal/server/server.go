package server

import (
  "feeder-backend/internal/config"
  cors "feeder-backend/internal/middleware"

  "fmt"
  "net/http"

  "github.com/gorilla/mux"
)

func StartServer(r *mux.Router) error {

  var port = ":" + config.Server.Port
	
	if port == "" {
		port = "8000"
	}

	fmt.Println("Server started at " + port)

  corsHandler := cors.SetupCORS(r)
  s := HandlerServer(port, corsHandler)

  err := s.ListenAndServe()
  if err != nil {
    fmt.Println("Server failed start", err)
  }

  return nil
}

func HandlerServer(port string, handler http.Handler) *http.Server {

  return &http.Server {
    Addr: port,
    Handler: handler,
  }
}

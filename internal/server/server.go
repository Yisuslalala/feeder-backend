package server

import (
  "feeder-backend/internal/config"
  cors "feeder-backend/internal/middleware"

  "fmt"
  "net/http"

  "github.com/gorilla/mux"
)

func StartServer(r *mux.Router) error {
  
  var port = ":" + config.ServerConfig["port"]

	// s.Handler = r
	// s.Addr = port

  // s.ListenAndServe()
	fmt.Println("Server started at " + port)
	// fmt.Println(s.ListenAndServe())

  corsHandler := cors.SetupCORS(r)

  s := HandlerServer(port, corsHandler)
  
  fmt.Printf("Server: %+v \n", s)
  fmt.Printf("Handler: %T \n", s.Handler)
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

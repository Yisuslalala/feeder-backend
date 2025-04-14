package middleware

import (
  "github.com/gorilla/handlers"
  "net/http"

  "feeder-backend/internal/config"
)

func SetupCORS(h http.Handler) http.Handler {

  var host = config.ServerConfig["host"]
  var port = config.ServerConfig["port"]

  var url = "http://" + host  + ":" + port

  return handlers.CORS(
    handlers.AllowCredentials(),
    handlers.AllowedMethods([]string{"GET, POST", "PUT", "DELETE", "OPTIONS"}),
    handlers.AllowedOrigins([]string{url}),
    handlers.AllowedHeaders([]string{"Content-Type"}), 
  )(h)
}

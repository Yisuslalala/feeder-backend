package middleware

import (
  "github.com/gorilla/handlers"
  "net/http"

  "feeder-backend/internal/config"
)

func SetupCORS(h http.Handler) http.Handler {

  var host = config.Server.Port
  var clientPort = config.Client.Port

  var serverUrl = "http://" + host  + ":" + clientPort
  var localhostUrl = "http://localhost:" + clientPort

  return handlers.CORS(
    handlers.AllowCredentials(),
    handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
   // Add the ip's of the frontend for local development or production
    handlers.AllowedOrigins([]string{serverUrl, localhostUrl}),
    handlers.AllowedHeaders([]string{"Content-Type"}), 
  )(h)
}

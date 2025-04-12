package middleware

import (
  "github.com/gorilla/handlers"
  "net/http"

  "/feeder-backend/internal/config"
)

func SetupCORS(h *http.Handler) http.Handler {

  var host = config.ServerConfig["host"]
  var port = config.ServerConfig["port"]

  var url = "http://" + host + port


  return h.CORS(
    h.AllowCredentials(),
    h.AllowedMethods([]string{"GET, POST", "PUT", "DELETE", "OPTIONS"}),
    h.AllowedOrigins([]string{url}),
    h.AllowedHeaders([]string{"Content-Type"}), 
  )h
}

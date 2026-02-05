package routes

import (
	"net/http"

	"feeder-backend/internal/controllers"
	"feeder-backend/internal/middleware"
	"github.com/gorilla/mux"
)

func RegisterUserRoutes(r *mux.Router, userController *controllers.UserController) {
	// Public routes
	r.HandleFunc("/api/auth/register", userController.Register).Methods(http.MethodPost)
	r.HandleFunc("/api/auth/login", userController.Login).Methods(http.MethodPost)

	// Protected routes (example for future use)
	protected := r.PathPrefix("/api/user").Subrouter()
	protected.Use(middleware.JWTMiddleware)
	// protected.HandleFunc("/profile", userController.GetProfile).Methods(http.MethodGet)
	// protected.HandleFunc("/update", userController.UpdateProfile).Methods(http.MethodPut)
}

package main

import (
	"log"

	config "feeder-backend/internal/config"
	"feeder-backend/internal/controllers"
	"feeder-backend/internal/db"
	"feeder-backend/internal/mqtt"
	"feeder-backend/internal/repositories"
	routes "feeder-backend/internal/routers"
	"feeder-backend/internal/server"
	"feeder-backend/internal/services"

	_ "github.com/eclipse/paho.mqtt.golang"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {

	config.Load()

	database, err := db.NewMySQLConnection()
	if err != nil {
		log.Fatalf("error with database: %v", err)
	}

	mqtt.InitMQTT()

	feederRepo := repositories.NewFeederRepository(database)

	feederService := services.NewFeederService(feederRepo)

	feederController := controllers.NewFeederController(feederService)

	userRepo := repositories.NewUserRepository(database)

	userService := services.NewUserService(userRepo)

	userController := controllers.NewUserController(userService)

	r := mux.NewRouter()
	routes.RegisterFeederRoutes(r, feederController)
	routes.RegisterUserRoutes(r, userController)

	// routes.StartServer(r)

	server.StartServer(r)
}

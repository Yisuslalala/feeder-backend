package main

import (

	"log"

	"feeder-backend/internal/db"
	"feeder-backend/internal/mqtt"
	routes "feeder-backend/internal/routers"
  "feeder-backend/internal/server"
	config "feeder-backend/internal/config"
	"feeder-backend/internal/controllers"
	"feeder-backend/internal/repositories"
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

	r := mux.NewRouter()
	routes.RegisterFeederRoutes(r, feederController)
  
  // routes.StartServer(r)

  server.StartServer(r)
}

package main

import (
	"feeder-backend/internal/db"
	"feeder-backend/internal/mqtt"
	routes "feeder-backend/internal/routers"
  "feeder-backend/internal/server"
	config "feeder-backend/internal/config"

	_ "github.com/eclipse/paho.mqtt.golang"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)


func main() {

	config.LoadEnv()
  db.InitDB()
	mqtt.InitMQTT()

	r := mux.NewRouter()
	routes.RegisterFeederRoutes(r)
  
  // routes.StartServer(r)

  server.StartServer(r)
}

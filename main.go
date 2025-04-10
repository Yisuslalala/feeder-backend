package main

import (
	"feeder-backend/internal/db"
	"feeder-backend/internal/mqtt"
	routes "feeder-backend/internal/routers"
	"fmt"
	"net/http"

	_ "github.com/eclipse/paho.mqtt.golang"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)


func main() {
	db.InitDB()
	
	mqtt.InitMQTT()

	r := mux.NewRouter()
	
	routes.RegisterFeederRoutes(r)
  
 routes.StartServer(r, s)  
	
}

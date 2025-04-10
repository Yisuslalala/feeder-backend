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

	port := ":8000"

	s := &http.Server{
	Handler: r,
	Addr: port,
	}

  // s.ListenAndServe()
	fmt.Println("Server started at " + port)
	fmt.Println(s.ListenAndServe())

  err := s.ListenAndServe()
  if err != nil {
    fmt.Println("Server failed start", err)
  }
}

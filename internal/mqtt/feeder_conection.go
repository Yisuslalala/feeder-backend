package mqtt

import (
	"fmt"
	"log"
	"os"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/joho/godotenv"
)

var mqttClient mqtt.Client

func InitMQTT() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ConnectionString := fmt.Sprintf("tcp://%s:%s", 
	os.Getenv("host"),
	os.Getenv("port"),
	)

	broker := ConnectionString
	clientID := os.Getenv("client_id")

	opts := mqtt.NewClientOptions().AddBroker(broker).SetClientID(clientID)

	// Create and connect the MQTT client
	mqttClient = mqtt.NewClient(opts)
	token := mqttClient.Connect()
	token.Wait()
	if token.Error() != nil {
		log.Fatal("MQTT Connection Error:", token.Error())
	}
	fmt.Println("Connected to MQTT broker")
}
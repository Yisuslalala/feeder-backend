package mqtt

import (
	"fmt"
	"feeder-backend/internal/config"
  "log"
	"net/http"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var mqttClient mqtt.Client

func InitMQTT() {

  config.LoadEnv()

	ConnectionString := fmt.Sprintf("tcp://%s:%s", 
	config.MQTTConfig["host"],
  config.MQTTConfig["port"],
	)

	broker := ConnectionString
	clientID := config.MQTTConfig["clientId"]

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

func ActivateMotor(w http.ResponseWriter, r *http.Request) {
	publishMessage("/motor", "ON")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Motor activated"))
}

func publishMessage(topic, payload string) {
  token := mqttClient.Publish(topic, 0, false, payload)
  token.Wait()
  fmt.Println("Message published", payload)
}

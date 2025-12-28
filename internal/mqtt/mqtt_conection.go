package mqtt

import (
	"fmt"
	"feeder-backend/internal/config"
  "log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var MqttClient mqtt.Client

func InitMQTT() {

  config.LoadEnv()

	ConnectionString := fmt.Sprintf("tcp://%s:%s", 
	config.MQTTConfig["host"],
  config.MQTTConfig["port"],
	)
	
	fmt.Println(config.MQTTConfig)

	broker := ConnectionString
	clientID := config.MQTTConfig["clientId"]

	opts := mqtt.NewClientOptions().AddBroker(broker).SetClientID(clientID)

	// Create and connect the MQTT client
	MqttClient = mqtt.NewClient(opts)
	token := MqttClient.Connect()
	token.Wait()
	if token.Error() != nil {
		log.Fatal("MQTT Connection Error:", token.Error())
	}
	fmt.Println("Connected to MQTT broker")
}


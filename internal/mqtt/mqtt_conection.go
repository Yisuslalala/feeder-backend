package mqtt

import (
	"fmt"
	"feeder-backend/internal/config"
  "log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var MqttClient mqtt.Client

func InitMQTT() {

	ConnectionString := fmt.Sprintf("tcp://%s:%s", 
	config.MQTT.Host,
  config.MQTT.Port,
	)
	
	broker := ConnectionString
	clientID := config.MQTT.ClientID

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


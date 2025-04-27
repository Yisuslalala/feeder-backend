package controllers

import (
  "fmt"
  "net/http"
  mqtt "feeder-backend/internal/mqtt"
)

func ActivateMotor(w http.ResponseWriter, r *http.Request) {
	publishMessage("/motor", "ON")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Motor activated"))
}

func publishMessage(topic, payload string) {
  token := mqtt.MqttClient.Publish(topic, 0, false, payload)
  token.Wait()
  fmt.Println("Message published", payload)
}

package controllers

import (
  "fmt"
  "net/http"
  mqtt "feeder-backend/internal/mqtt"
)

func ActivateMotor(w http.ResponseWriter, r *http.Request) {
  PublishMessage("/motor", "ON")

  //  if err != nil {
    // http.Error(w, "Failed to activate motor", http.StatusInternalServerError)
    // return
  // }

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Motor activated"))
}

func PublishMessage(topic, payload string) {
  token := mqtt.MqttClient.Publish(topic, 0, false, payload)
  token.Wait()

  // if token.Error() != nil {
  //   return token.Error()
  // }

  fmt.Println("Message published", payload)
}

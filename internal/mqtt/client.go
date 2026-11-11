package mqtt

import (
	"errors"
	"fmt"
	"log"
	"time"

	"feeder-backend/internal/config"

	paho "github.com/eclipse/paho.mqtt.golang"
)

func NewClient(cfg config.MQTTConfig) (paho.Client, error) {
	broker := fmt.Sprintf("tcp://%s:%s", cfg.Host, cfg.Port)

	opts := paho.NewClientOptions().
		AddBroker(broker).
		SetClientID(cfg.ClientID).
		SetConnectTimeout(5 * time.Second)

	if cfg.User != "" {
		opts.SetUsername(cfg.User)
		opts.SetPassword(cfg.Password)
	}
	opts.SetAutoReconnect(true)
	opts.SetConnectionLostHandler(func(c paho.Client, err error) {
		log.Printf("MQTT connection lost: %v", err)
	})
	opts.SetOnConnectHandler(func(c paho.Client) {
		log.Println("MQTT connected")
	})

	client := paho.NewClient(opts)
	token := client.Connect()

	if !token.WaitTimeout(10 * time.Second) {
		return nil, errors.New("time out MQTT broker " + broker)
	}

	if err := token.Error(); err != nil {
		return nil, fmt.Errorf("connecting to MQTT broker %s: %w", broker, err)
	}

	log.Println("connection to MQTT broker", broker)

	return client, nil

}

package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	User string
	Pass string
	Host string
	Port string
	Name string
}

type MQTTConfig struct {
	Host     string
	Port     string
	ClientID string
}

type ServerConfig struct {
	Port string
}

type ClientConfig struct {
	Origin string
	Port   string
}

type JWTConfig struct {
	Secret string
}

var (
	DB     DBConfig
	MQTT   MQTTConfig
	Server ServerConfig
	Client ClientConfig
	JWT    JWTConfig
)

func Load() {

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	DB = DBConfig{
		User: os.Getenv("DB_USER"),
		Pass: os.Getenv("DB_PASS"),
		Host: os.Getenv("DB_HOST"),
		Port: os.Getenv("DB_PORT"),
		Name: os.Getenv("DB_NAME"),
	}

	MQTT = MQTTConfig{
		Host:     os.Getenv("MQTT_HOST"),
		Port:     os.Getenv("MQTT_PORT"),
		ClientID: os.Getenv("MQTT_CLIENT_ID"),
	}

	Server = ServerConfig{
		Port: os.Getenv("SERVER_PORT"),
	}

	Client = ClientConfig{
		Origin: os.Getenv("CLIENT_ORIGIN"),
		Port:   os.Getenv("CLIENT_PORT"),
	}

	JWT = JWTConfig {
		Secret: os.Getenv("JWT_SECRET"),
	}

	validate()
}

func validate() {

	if DB.User == "" || DB.Pass == "" || DB.Host == "" || DB.Port == "" || DB.Name == "" {
		log.Fatal("Missing required database environment variables")
	}

	if Server.Port == "" {
		log.Fatal("PORT environment variable is required (Render sets this automatically)")
	}

	if MQTT.Host != "" && MQTT.Port == "" {
		log.Fatal("MQTT_PORT must be set if MQTT_HOST is provided")
	}
	
	if JWT.Secret == "" {
		log.Fatal("JWT_SECRET is required")
	}

	if len(JWT.Secret) < 32 {
		log.Fatal("JWT_SECRET must contain 32 characters")
	}

	log.Println("Configuration loaded successfully")
}


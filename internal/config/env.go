package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	DBConfig = make(map[string]string)
	MQTTConfig = make(map[string]string)
	HTTPQueries = make(map[string]string)
  ServerConfig = make(map[string]string)
  ClientConfig = make(map[string]string)
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	DBConfig["user"] = os.Getenv("DB_USER")
	DBConfig["pass"] = os.Getenv("DB_PASS")
	DBConfig["host"] = os.Getenv("DB_HOST")
	DBConfig["port"] = os.Getenv("DB_PORT")
	DBConfig["dbName"] = os.Getenv("DB_NAME")

	MQTTConfig["host"] = os.Getenv("DB_PORT")
  MQTTConfig["port"] = os.Getenv("MQTT_PORT")
	MQTTConfig["clientId"] = os.Getenv("MQTT_CLIENT_ID")

	HTTPQueries["user"] = os.Getenv("DB_USER")
	HTTPQueries["pass"] = os.Getenv("DB_PASS")
	HTTPQueries["port"] = os.Getenv("DB_PORT")
	HTTPQueries["host"] = os.Getenv("DB_HOST")
	HTTPQueries["dbName"] = os.Getenv("DB_NAME")
  
  ServerConfig["host"] = os.Getenv("DB_HOST")
  ServerConfig["port"] = os.Getenv("DB_PORT")

  ClientConfig["clientHost"] = os.Getenv("CLIENT_ORIGIN")
	ClientConfig["clientPort"] = os.Getenv("CLIENT_PORT")
}



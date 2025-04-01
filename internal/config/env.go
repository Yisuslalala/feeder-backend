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
)

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	DBConfig["user"] = os.Getenv("user")
	DBConfig["pass"] = os.Getenv("pass")
	DBConfig["host"] = os.Getenv("host")
	DBConfig["port"] = os.Getenv("sql_port")
	DBConfig["dbName"] = os.Getenv("db_name")

	MQTTConfig["host"] = os.Getenv("host")
  MQTTConfig["port"] = os.Getenv("mqtt_port")
  MQTTConfig["clientId"] = os.Getenv("client_id")

	HTTPQueries["user"] = os.Getenv("user")
	HTTPQueries["pass"] = os.Getenv("pass")
	HTTPQueries["port"] = os.Getenv("sql_port")
	HTTPQueries["host"] = os.Getenv("host")
	HTTPQueries["dbName"] = os.Getenv("db_name")
}



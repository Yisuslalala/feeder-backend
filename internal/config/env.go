package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	DBConfig map[string]string
	MQTTConfig map[string]string
	HTTPQueries map[string]string
)

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	 DBConfig = map[string]string {
		"user": os.Getenv("user"),
		"pass": os.Getenv("pass"),
		"host": os.Getenv("host"),
		"port": os.Getenv("port"),
		"dbName": os.Getenv("db_name"),
	 }

	 MQTTConfig = map[string]string {
		"host": os.Getenv("host"),
		"port": os.Getenv("port"),
	 }

	 HTTPQueries = map[string] string {
		"user": os.Getenv("user"),
		"pass": os.Getenv("pass"),
		"port": os.Getenv("port"),
		"host": os.Getenv("host"),
		"dbName": os.Getenv("db_name"),
	 }

}



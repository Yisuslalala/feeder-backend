package db

import (
	"database/sql"
	"feeder-backend/internal/config"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	// err := godotenv.Load(".env") // Load environment variables
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }
  fmt.Println("DBConfig content:", config.DBConfig)
  config.LoadEnv()

	// Construct connection string
	ConnectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		config.DBConfig["user"],
		config.DBConfig["pass"],
		config.DBConfig["host"],
		config.DBConfig["port"],
		config.DBConfig["dbName"],
	)

  for key, value := range config.DBConfig {
    fmt.Println("Firts key", key)
    fmt.Println("First value", value)
  }

  fmt.Println(`string`)
  fmt.Println(config.DBConfig)
  fmt.Println(`This is the connection string: `, ConnectionString)

  var err error
	DB, err = sql.Open("mysql", ConnectionString)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Test the connection
	if err := DB.Ping(); err != nil {
		log.Fatal("Database connection is not active:", err)
	}
}

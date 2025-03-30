package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var DB *sql.DB

func InitDB() {
	err := godotenv.Load(".env") // Load environment variables
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Construct connection string
	ConnectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("user"),
		os.Getenv("pass"),
		os.Getenv("host"),
		os.Getenv("port"),
		os.Getenv("db_name"))

	DB, err = sql.Open("mysql", ConnectionString)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Test the connection
	if err := DB.Ping(); err != nil {
		log.Fatal("Database connection is not active:", err)
	}
}

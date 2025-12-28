package db

import (
	"database/sql"
	"feeder-backend/internal/config"
	"fmt"
  "time"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
  // fmt.Println("DBConfig content:", config.DBConfig)
 
	ConnectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		config.DBConfig["user"],
		config.DBConfig["pass"],
		config.DBConfig["host"],
		config.DBConfig["port"],
		config.DBConfig["dbName"],
	)

  var err error
	DB, err = sql.Open("mysql", ConnectionString)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatal("Database connection is not active:", err)
	}

  // Some documentation recommendations
  DB.SetConnMaxLifetime(time.Minute * 3)
  DB.SetMaxOpenConns(1)
  DB.SetMaxIdleConns(1)
}

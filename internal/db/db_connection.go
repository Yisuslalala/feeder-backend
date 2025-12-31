package db

import (
	"database/sql"
	"feeder-backend/internal/config"
	"fmt"
  "time"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func NewMySQLConnection() (*sql.DB, err) {
	ConnectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		config.DB.User,
		config.DB.Pass,
		config.DB.Host,
		config.DB.Port,
		config.DB.Name,
	)

	db, err = sql.Open("mysql", ConnectionString)
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to database: %w", err)
	}

	if err := DB.Ping(); err != nil {
		return nil, fmt.Errorf("Database connection is not active: %w", err)
	}

  // Some documentation recommendations
  DB.SetConnMaxLifetime(time.Minute * 3)
  DB.SetMaxOpenConns(1)
  DB.SetMaxIdleConns(1)

	return db, nil
}


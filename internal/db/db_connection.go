package db

import (
	"database/sql"
	config "feeder-backend/internal/config"
	"fmt"
  "time"

	_ "github.com/go-sql-driver/mysql"
)

func NewMySQLConnection() (*sql.DB, error) {
	ConnectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		config.DB.User,
		config.DB.Pass,
		config.DB.Host,
		config.DB.Port,
		config.DB.Name,
	)

	db, err := sql.Open("mysql", ConnectionString)
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("Database connection is not active: %w", err)
	}

  // Some documentation recommendations
  db.SetConnMaxLifetime(time.Minute * 3)
  db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(1)

	return db, nil
}


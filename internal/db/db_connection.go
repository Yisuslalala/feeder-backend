package db

import (
	"database/sql"
	"fmt"
	"time"

	config "feeder-backend/internal/config"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func NewConnection() (*sql.DB, error) {
	db, err := sql.Open("pgx", config.DB.URL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("database connection is not active: %w", err)
	}

	// Some documentation recommendations
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db, nil
}

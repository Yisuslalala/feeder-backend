package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func initDB() {
	var err error
	dsn := "user:password@tcp(127.0.0.1:3306)/feeder_backend"
	db, err = sql.Open("mysql", dsn)

	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	fmt.Println("Connected to MySQL successfully!")
}

func getUsers() {
	rows, err := db.Query("SELECT id, name, email FROM users")
	if err != nil {
		log.Fatal("Error fetching users:", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name, email string
		err := rows.Scan(&id, &name, &email)
		if err != nil {
			log.Fatal("Error scanning data:", err)
		}
		fmt.Printf("ID: %d, Name: %s, Email: %s\n", id, name, email)
	}
}


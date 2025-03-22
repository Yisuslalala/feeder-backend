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
	dsn := "yisus:1235@tcp(127.0.0.1:3306)/feeder"
	db, err = sql.Open("mysql", dsn)

	if err != nil {
		log.Fatal("Error connection to database", err)
	}	

	err = db.Ping()
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	fmt.Println("Connected to MySQL successfully!")
}

func createUser(name, email string) {
	query := "INSERT INTO users (name, email) VALUES (?, ?)"
	result, err := db.Exec(query, name, email)
	if err != nil {
		log.Fatal("Error inserting user:", err)
	}

	id, _ := result.LastInsertId()
	fmt.Println("User added with ID:", id)
}

func getFeeds() {
	rows, err := db.Query("SELECT id, feed_at FROM feed_details")
	if err != nil {
		log.Fatal("Error fetching users:", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var feed_detail string
		err := rows.Scan(&id, &feed_detail)
		if err != nil {
			log.Fatal("Error scanning data:", err)
		}
    fmt.Printf("id: %d, feeded at: %s", id, feed_detail)
		//fmt.Printf("ID: %d, Name: %s, Email: %s\n", id, name, email)
	}
}

func updateUser(id int, name, email string) {
	query := "UPDATE users SET name = ?, email = ? WHERE id = ?"
	_, err := db.Exec(query, name, email, id)
	if err != nil {
		log.Fatal("Error updating user:", err)
	}

	fmt.Println("User updated successfully.")
}

func deleteUser(id int) {
	query := "DELETE FROM users WHERE id = ?"
	_, err := db.Exec(query, id)
	if err != nil {
		log.Fatal("Error deleting user:", err)
	}

	fmt.Println("User deleted successfully.")
}

func main() {
	initDB()
	defer db.Close()

	getUsers()
}

package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var db *sql.DB
var dbDriver = "mysql"
var dbUser = "yisus"
var dbPass = "1235"
var dbName = "feeder"

type Feeder struct {
	id int
	feed_at string
}

func getFeederHandler(w http.ResponseWriter, r *http.Request) {
  db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp(192.168.1.125:3306)/"+dbName)

	if err != nil {
		http.Error(w, "Database connection failed", http.StatusInternalServerError)
		return	
	}

	defer db.Close()

  rows, err := db.Query("SELECT id, feed_at FROM feeder_details")

  if err != nil {
    http.Error(w, "Failed to fetch feeder_details", http.StatusInternalServerError)
    return
  }

  defer rows.Close()

	var feders []Feeder

  for rows.Next() {
    var  feeder Feeder
    err := rows.Scan(&feeder.id, &feeder.feed_at)

    if err != nil {
      http.Error(w, "Error scanning data from feeeder_details", http.StatusInternalServerError)
      return
    }
  
    feeders = append(feeders, feeder)
  }

  if err := rows.Err(); err != nil {
    http.Error(w. "Error readind rows of feeder_details", http.StatusInternalServerError)
    return
  }

  w.Header().Set("Content-Type", "application/json")
  json.NewEnconder(w).Encode(feeders)

}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/feeds", getFeederHandler).Methods("GET")

	log.Println("Server listening on: 8090")
}

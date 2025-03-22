package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// var db *sql.DB
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

	var feeders []Feeder

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
    http.Error(w, "Error readind rows of feeder_details", http.StatusInternalServerError)
    return
  }

  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(feeders)

}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func main() {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/feeds", getFeederHandler).Methods("GET")
	r.HandleFunc("/", homeLink)
	log.Fatal(http.ListenAndServe(":8090", r))
}

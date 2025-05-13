package controllers

import (
	"fmt"

	"database/sql"
	"encoding/json"
	"feeder-backend/internal/config"
	"feeder-backend/internal/models"
	"net/http"
)

func GetFeedDetails(w http.ResponseWriter, r *http.Request) {
		
	details, err := getDetails()
	if err == nil {
		responseSuccess(details, w)
	} else {
		responseError(err, w)
	}
}

func getDetails() ([]models.Detail, error) {
	
  details := []models.Detail{}
	
	endpointString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
	config.HTTPQueries["user"],
	config.HTTPQueries["pass"],
	config.HTTPQueries["host"],
	config.HTTPQueries["port"],
	config.HTTPQueries["dbName"],
	)

	db, err := sql.Open("mysql", endpointString)
	if err != nil {
		return details, err
	}

	rows, err := db.Query("SELECT id, feed_at FROM feeder_details")
	if err != nil {
		return details, err
	}

	for rows.Next() {
		var detail models.Detail
		err = rows.Scan(&detail.ID, &detail.FeedAt)
		if err != nil {
			return details, err
		}

		details = append(details, detail)
	}

	return details, nil
}

// TODO: Create function for adding a new feed detail
func CreateDetail(w http.ResponseWriter, r *http.Request) {
  // params := mux.Vars(r)
  // fmt.Print(`Params: `, params)
  // Create endpointString
  endpointString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
    config.HTTPQueries["user"],
    config.HTTPQueries["pass"],
    config.HTTPQueries["host"],
    config.HTTPQueries["port"],
    config.HTTPQueries["dbName"],
  )

  // Open session for sql interface
  db, err := sql.Open("mysql", endpointString)
  if err != nil {
    fmt.Println("Error at opening sql interface")
  }
  // Prepare sql query
  query, err := db.Prepare("INSERT INTO feeder_details VALUES()")
  if err != nil {
    fmt.Println("Error at prepare sql query")
  }

  // Execute it and handle errors
  _, err = query.Exec()
  // fmt.Println("res", res)
  if err != nil {
    http.Error(w, "Failed to create feeding", http.StatusInternalServerError)
  }

  // Add move motor mqtt controller
  PublishMessage("/motor", "ON")
  if err != nil {
    http.Error(w, "Failed to activate motor" , http.StatusInternalServerError)
    return
  }

  w.WriteHeader(http.StatusCreated)
  fmt.Println("Feeding detail created successfully")
}

func responseSuccess(data interface{}, w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
  w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func responseError(err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(err.Error())
}

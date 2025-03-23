
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"net/http"
	"encoding/json"
)	

type Detail struct {
	id int
	feed_at string
}

func main() {
	db, err:= sql.Open("mysql", "yisus:1235@tcp(192.168.1.125:3306)/feeder")
	if err != nil {
		fmt.Println("Error with database connection: " + err.Error())
	} else {
		err = db.Ping()
		if err != nil {
			fmt.Println("error making connection to DB, verify credentials" + err.Error())
		}	
	}

	r := mux.NewRouter()
	
	r.HandleFunc("/details", func(w http.ResponseWriter, r *http.Request) {
		
		details, err := getDetails()
		if err == nil {
			responseSuccess(details, w)
		} else {
			responseError(err, w)
		}

	}).Methods(http.MethodGet)

	port := ":8000"

	s := &http.Server{
	Handler: r,
	Addr: port,
	}

	fmt.Println("Server started at " + port)
	fmt.Println(s.ListenAndServe())
}	

func getDetails() ([]Detail, error) {
	
	details := []Detail{}
		
	db, err := sql.Open("mysql", "yisus:1235@tcp(192.168.1.125:3306)/feeder")
	if err != nil {
		return details, err
	}

	rows, err := db.Query("SELECT id, feed_at FROM feeder_details")
	if err != nil {
		return details, err
	}

	for rows.Next() {
		var detail Detail
		err = rows.Scan(&detail.id, &detail.feed_at)
		if err != nil {
			return details, err
		}

		details = append(details, detail)
	}

	return details, nil
}

func responseSuccess(data interface{}, w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func responseError(err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(err.Error())
}

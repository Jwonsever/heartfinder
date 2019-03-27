package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

type row struct {
	Name string
}

// Handler returns http.Handler for API endpoint
func Handler() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {

		res.Header().Set("Content-Type", "application/json")

		connStr := "user=jwonsever dbname=pqgotest sslmode=disable"
		db, err := sql.Open("postgres", connStr)

		if err != nil {
			fmt.Printf("Could not open the db: %s", err.Error())
			res.WriteHeader(500)
			return
		}

		rows, err := db.Query("SELECT name FROM hearts")
		if err != nil {
			fmt.Printf("Could not connect to db: %s", err.Error())
			res.WriteHeader(500)
			return
		}

		got := []row{}
		for rows.Next() {
			var r row
			err = rows.Scan(&r.Name)
			if err != nil {
				fmt.Printf("Scan: %v", err)
			}
			got = append(got, r)
		}

		//TODO, I DONT GET JSON MARSHALLING YET
		body, err := json.Marshal(map[string]interface{}{
			"data": got,
		})

		if err != nil {
			res.WriteHeader(500)
			return
		}

		res.WriteHeader(200)
		res.Write(body)
	}
}

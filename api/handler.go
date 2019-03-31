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

		//https://github.com/lib/pq/issues/645 - why i need host...
		connStr := "user=james host=/run/postgresql dbname=james sslmode=disable"
		db, err := sql.Open("postgres", connStr)

		if err != nil {
			fmt.Printf("Could not open the db: %s", err.Error())
			res.WriteHeader(500)
			return
		}
		rows, err := db.Query("SELECT heart_name FROM hearts")
		if err != nil {
			db.Close()
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

		db.Close()

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

// Handler returns http.Handler for API endpoint
func PostHandler() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json")

		//https://github.com/lib/pq/issues/645 - why i need host...
		connStr := "user=james host=/run/postgresql dbname=james sslmode=disable"
		db, err := sql.Open("postgres", connStr)

		if req.Method != "POST" {
			fmt.Printf("Wrong method for POST handler: %s", req.Method)
			res.WriteHeader(404)
			return
		}

		if err := req.ParseForm(); err != nil {
			fmt.Fprintf(res, "ParseForm() err: %v", err)
			return
		}

		fmt.Printf("Post from website! req.PostFrom = %v\n", req.PostForm)
		name := req.FormValue("name")
		fmt.Printf("Name = %s\n", name)

		if len(name) == 0 {
			fmt.Printf("Name is empty, return a 400\n")
			res.WriteHeader(400)
			return
		}

		_, err = db.Query("INSERT into hearts (heart_name, created_on) values ('" + name + "', CURRENT_TIMESTAMP);")

		if err != nil {
			fmt.Printf("Could not insert into db: %s", err.Error())
			res.WriteHeader(500)
			return
		}

		res.WriteHeader(200)
	}
}

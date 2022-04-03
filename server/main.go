package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

// var Db *sql.DB

func init() {
	_, err := sql.Open("postgres", "host=twitter_db user=shion password=pass dbname=twitter port=5432 sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	// db の　name 先頭を表示する
	fmt.Fprintln(w, "hellso")
}

func main() {
	http.HandleFunc("/", hello)
	// rows, err := Db.Query("select * from ")
	http.ListenAndServe(":8080", nil)
}

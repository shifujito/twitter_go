package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type User struct {
	Id   int    `gorm:"column:id`
	Name string `gorm:column:name`
}

// var db *gorm.DB

func init() {
	db, err := gorm.Open("postgres", "host=twitter_db user=shion password=pass dbname=twitter port=5432 sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()
}

func dbConnect() (db *gorm.DB) {
	db, err := gorm.Open("postgres", "host=twitter_db user=shion password=pass dbname=twitter port=5432 sslmode=disable")
	if err != nil {
		log.Fatalln("not conncetd postgres database", err)
	}
	return
}

func getFirstRow(firstUser *User) {
	// connect db
	db := dbConnect()
	// get first data
	db.First(&firstUser)
	// disconnect
	defer db.Close()
}

func hello(w http.ResponseWriter, r *http.Request) {
	// db の　name 先頭を表示する
	var firstUser User
	getFirstRow(&firstUser)
	// defer db.Close()
	fmt.Fprintln(w, firstUser)
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8080", nil)
}

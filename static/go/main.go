package main

import (
	"database/sql"
	"fmt"
	"net/http"
)

var db *sql.DB
var err error
var userID int
var userRole string
var userName string
var userPassword string

func main() {
	db, err = sql.Open("sqlite", "../db/OnlineLibrary.db")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	createTables()

	if DataBaseIsEmpty() {
		fmt.Println("DataBase is Empty")
	}

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("../css"))))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("../img"))))
	http.Handle("/html/", http.StripPrefix("/html/", http.FileServer(http.Dir("../html"))))

	http.HandleFunc("/Login", login)

	fmt.Println(":8080")
	http.ListenAndServe(":8000", nil)
}

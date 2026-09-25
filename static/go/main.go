package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
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
func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		id := r.FormValue("id")
		userID, err = strconv.Atoi(id)
		if err != nil {
			fmt.Println(err)
		}
		password := r.FormValue("password")

		fmt.Println("id:", id)
		fmt.Println("Password:", password)

		db.QueryRow("SELECT name,password,role FROM users WHERE id = ?", userID).Scan(&userName, &userPassword, &userRole)
		if password == userPassword {
			template, err := template.ParseFiles("../html/main.html")
			if err != nil {
				fmt.Println(err)
				return
			}
			template.Execute(w, userName)

			return
		} else {
			template, err := template.ParseFiles("../html/Login.html")
			if err != nil {
				fmt.Println(err)
				return
			}
			template.Execute(w, "Wrong user id or password")
			return
		}
	}
	template, err := template.ParseFiles("../html/Login.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	template.Execute(w, "")
}

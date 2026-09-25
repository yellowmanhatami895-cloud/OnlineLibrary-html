package main

import (
	"database/sql"
	"fmt"
)

var db *sql.DB
var err error

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
		// record := getInputs([]string{"Enter first username", "Enter email", "Enter first password"})
		//insertIntoUsers(append(record, "Admin"))
	}
}

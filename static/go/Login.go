package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

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
			switch userRole {
			case "Customer":
				template, err := template.ParseFiles("../html/CustomerHome.html")
				if err != nil {
					fmt.Println(err)
					return
				}
				template.Execute(w, userName)
			case "Admin":

				template, err := template.ParseFiles("../html/AdminHome.html")
				if err != nil {
					fmt.Println(err)
					return
				}
				template.Execute(w, userName)
			case "Author":
				template, err := template.ParseFiles("../html/AuthorHome.html")
				if err != nil {
					fmt.Println(err)
					return
				}
				template.Execute(w, userName)
			}
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

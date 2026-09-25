package main

import "fmt"

func checkEmail(oldEmail string, newEmail string) {
	row, err := db.Query("SELECT email FROM users")
	if err != nil {
		fmt.Println(err)
		return
	}
	var email string
	for row.Next() {
		row.Scan(&email)
		if email == oldEmail {
			continue
		}
		if email == newEmail {
			fmt.Println("Email is already used")
		}
	}
	err = row.Err()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func insertCustomerIntoUsers() {
	//record := getInputs([]string{"Enter name", "Email", "Password"})
	//record = append(record, "Customer")
	//insertIntoUsers(record)
}
func insertIntoUsers(record []string) {
	if len(record) != 4 {
		fmt.Println("wrong length")
		return
	}
	checkEmail("", record[1])
	result, err := db.Exec("INSERT INTO users(name,email,password,role) VALUES(?,?,?,?)", record[0], record[1], record[2], record[3])
	if err != nil {
		fmt.Println(err)
		return
	}
	id, err := result.LastInsertId()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Your id  :", id)
}

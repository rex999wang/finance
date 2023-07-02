package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func addUser(reader *bufio.Reader) {
	fmt.Print("Enter username: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	fmt.Print("Enter password: ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	_, err := db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", username, password)
	if err != nil {
		fmt.Println("Failed to add user:", err)
	}
}

func viewUsers() {
	rows, err := db.Query("SELECT id, username, password FROM users")
	if err != nil {
		fmt.Println("Failed to fetch users:", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var username string
		var password string
		err = rows.Scan(&id, &username, &password)
		if err != nil {
			fmt.Println("Failed to fetch user:", err)
		}
		fmt.Printf("ID: %d, Username: %s, Password: %s\n", id, username, password)
	}
}

func deleteUser(reader *bufio.Reader) {
	fmt.Print("Enter the ID of the user you want to delete: ")
	text, _ := reader.ReadString('\n')
	id, err := strconv.Atoi(strings.TrimSpace(text))
	if err != nil {
		fmt.Println("Invalid ID. Please enter a number.")
		return
	}

	_, err = db.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		fmt.Println("Failed to delete user:", err)
	}
}

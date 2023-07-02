package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func addExpense(reader *bufio.Reader) {
	fmt.Print("Enter expense category: ")
	category, _ := reader.ReadString('\n')
	category = strings.TrimSpace(category)

	fmt.Print("Enter expense amount: ")
	text, _ := reader.ReadString('\n')
	amount, err := strconv.ParseFloat(strings.TrimSpace(text), 64)
	if err != nil {
		fmt.Println("Invalid amount. Please enter a number.")
		return
	}

	_, err = db.Exec("INSERT INTO expenses (category, amount) VALUES (?, ?)", category, amount)
	if err != nil {
		fmt.Println("Failed to add expense:", err)
	}
}

func viewExpenses() {
	rows, err := db.Query("SELECT id, category, amount FROM expenses")
	if err != nil {
		fmt.Println("Failed to fetch expenses:", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var category string
		var amount float64
		err = rows.Scan(&id, &category, &amount)
		if err != nil {
			fmt.Println("Failed to fetch expense:", err)
		}
		fmt.Printf("ID: %d, Category: %s, Amount: %.2f\n", id, category, amount)
	}
}

func deleteExpense(reader *bufio.Reader) {
	fmt.Print("Enter the ID of the expense you want to delete: ")
	text, _ := reader.ReadString('\n')
	id, err := strconv.Atoi(strings.TrimSpace(text))
	if err != nil {
		fmt.Println("Invalid ID. Please enter a number.")
		return
	}

	_, err = db.Exec("DELETE FROM expenses WHERE id = ?", id)
	if err != nil {
		fmt.Println("Failed to delete expense:", err)
	}
}

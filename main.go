package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("mysql", "root:wzx920929@/finance")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Choose an operation:")
		fmt.Println("1 - Manage Users")
		fmt.Println("2 - Manage Expenses")
		fmt.Println("3 - Quit")
		fmt.Print("Enter the number of your choice: ")

		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		switch text {
		case "1":
			userOption(reader)
		case "2":
			expenseOption(reader)
		case "3":
			return
		default:
			fmt.Println("Invalid choice. Please enter a number between 1 and 3.")
		}
	}
}

func userOption(reader *bufio.Reader) {
	for {
		fmt.Println("Choose an operation:")
		fmt.Println("1 - Add user")
		fmt.Println("2 - View users")
		fmt.Println("3 - Delete user")
		fmt.Println("4 - Back to main")
		fmt.Print("Enter the number of your choice: ")

		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		switch text {
		case "1":
			addUser(reader)
		case "2":
			viewUsers()
		case "3":
			deleteUser(reader)
		case "4":
			return
		default:
			fmt.Println("Invalid choice. Please enter a number between 1 and 4.")
		}
	}
}

func expenseOption(reader *bufio.Reader) {
	for {
		fmt.Println("Choose an operation:")
		fmt.Println("1 - Add expense")
		fmt.Println("2 - View expenses")
		fmt.Println("3 - Delete expense")
		fmt.Println("4 - Back to main")
		fmt.Print("Enter the number of your choice: ")

		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		switch text {
		case "1":
			addExpense(reader)
		case "2":
			viewExpenses()
		case "3":
			deleteExpense(reader)
		case "4":
			return
		default:
			fmt.Println("Invalid choice. Please enter a number between 1 and 4.")
		}
	}
}

package main

import (
	"expence-tracker/internal/expense"
	"fmt"
	"time"
)

func showHeader() {
	fmt.Println("================================")
	fmt.Println("       CLI EXPENSE TRACKER")
	fmt.Println("================================")
	fmt.Println("Welcome to Expense Tracker!")
}

func showMenu() {
	fmt.Println()
	fmt.Println("1. Add Expense")
	fmt.Println("2. List Expense")
	fmt.Println("3. Summary")
	fmt.Println("4. Exit")
	fmt.Println("5. Help")
	fmt.Println("================================")
}

func showFooter() {
	fmt.Println("================================")
	fmt.Println(" Developed with Go Standard Lib")
	fmt.Println("================================")
}

func main() {
	showHeader()
	showMenu()

	transaction := expense.Expense{
		ID:          1,
		Description: "Makan Siang",
		Amount:      25000,
		Category:    "Food",
		Date:        time.Now(),
	}
	fmt.Println("================================")
	fmt.Println("       EXPENSE DETAIL")
	fmt.Println("================================")

	fmt.Printf("ID          : %d\n", transaction.ID)
	fmt.Printf("Description : %s\n", transaction.Description)
	fmt.Printf("Amount      : Rp%.0f\n", transaction.Amount)
	fmt.Printf("Category    : %s\n", transaction.Category)
	fmt.Printf("Date        : %s\n", transaction.Date.Format("2006-01-02"))

	showFooter()
}

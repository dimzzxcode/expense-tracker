package main

import (
	"bufio"
	"expense-tracker/internal/expense"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func addExpense(scanner *bufio.Scanner, id int) expense.Expense {
	var transaction expense.Expense
	
	// input description
	for {
		fmt.Print("Description: ")
		if !scanner.Scan() {
			return transaction
		}

		description := strings.TrimSpace(scanner.Text())
		if description == "" {
			fmt.Println("Invalid description. Please enter a desciption.")
			continue
		}

		transaction.Description = description
		break
	}

	// input amount
	for {
		fmt.Print("Amount: ")
		if !scanner.Scan() {
			return transaction
		}
		
		amountInput := strings.TrimSpace(scanner.Text())

		amount, err := strconv.ParseFloat(amountInput, 64)
		if err != nil || amount <= 0 {
			fmt.Println("Invalid amount. Please enter a positive number.")
			continue
		}

		transaction.Amount = amount
		break
	}

	// input category
	for {
		fmt.Print("Category: ")
		if !scanner.Scan() {
			return transaction
		}

		category := strings.TrimSpace(scanner.Text())
		if category == "" {
			fmt.Println("Invalid category. Please enter a category.")
			continue
		}

		transaction.Category = category
		break
	}

	// set transaction ID and Date
	transaction.ID = id
	transaction.Date = time.Now()

	return transaction
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	showHeader()
	showMenu()

	var totalToAdd int
	for {
		fmt.Print("\nHow many expenses would you like to input? : ")
		if !scanner.Scan() {
			return
		}

		input := strings.TrimSpace(scanner.Text())
		n, err := strconv.Atoi(input)
		if err != nil || n <= 0 {
			fmt.Println("Invalid input. Enter a valid positive number.")
			continue
		}
		totalToAdd = n
		break
	}

	// slice menampung banyak expense
	var expenses []expense.Expense

	for i := 1; i <= totalToAdd; i++ {
		fmt.Printf("\n=== Add New Expense (%d/%d) ===\n", i, totalToAdd)
		// kirim i sebagai ID transaksi
		transaction := addExpense(scanner, i)
		// masukan kedalam slice
		expenses = append(expenses, transaction)
	}

	err := expense.SaveExpense(expenses)
	if err != nil {
		fmt.Println("failed to save expenses:", err)
		return
	}
	fmt.Println("Expenses saved successfully!")

	fmt.Println("================================")
	fmt.Println("       EXPENSE DETAIL")
	fmt.Println("================================")

	for _, tx := range expenses {
		fmt.Printf("ID          : %d\n", tx.ID)
		fmt.Printf("Description : %s\n", tx.Description)
		fmt.Printf("Amount      : Rp%.0f\n", tx.Amount)
		fmt.Printf("Category    : %s\n", tx.Category)
		fmt.Printf("Date        : %s\n", tx.Date.Format("2006-01-02"))
		fmt.Println("--------------------------------")
	}

	showFooter()
}

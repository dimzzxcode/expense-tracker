package main

import (
	"bufio"
	"expense-tracker/internal/expense"
	"flag"
	"fmt"
	"os"
	"slices"
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
	fmt.Println("6. Update Expense")
	fmt.Println("7. Delete Expense")
	fmt.Println("8. Search Expense")
	fmt.Println("9. Filter Expense")
	fmt.Println("10. Sort Expense")
	fmt.Println("11. Export CSV")
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

func listExpenses() {
	expenses, err := expense.LoadExpenses()
	if err != nil {
		fmt.Println("failed to load expenses:", err)
		return
	}

	if len(expenses) == 0 {
		fmt.Println("no expenses found.")
		return
	}

	slices.SortFunc(expenses, func(a, b expense.Expense) int {
		return a.ID - b.ID
	})

	fmt.Println("\n========== EXPENSE LIST ==========")

    for _, tx := range expenses {
        fmt.Printf("ID          : %d\n", tx.ID)
        fmt.Printf("Description : %s\n", tx.Description)
        fmt.Printf("Amount      : Rp%.0f\n", tx.Amount)
        fmt.Printf("Category    : %s\n", tx.Category)
        fmt.Printf("Date        : %s\n", tx.Date.Format("2006-01-02"))
        fmt.Println("----------------------------------")
    }
}

func updateExpense(scanner *bufio.Scanner) {
    expenses, err := expense.LoadExpenses()
    if err != nil {
        fmt.Println("Failed to load expenses:", err)
        return
    }

    fmt.Print("Enter expense ID to update: ")
    if !scanner.Scan() {
        return
    }

    id, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
    if err != nil {
        fmt.Println("Invalid ID.")
        return
    }

    updated := addExpense(scanner, id)

    expenses, err = expense.UpdateExpense(expenses, id, updated)
    if err != nil {
        fmt.Println(err)
        return
    }

    if err := expense.SaveExpense(expenses); err != nil {
        fmt.Println("Failed to save expenses:", err)
        return
    }

    fmt.Println("Expense updated successfully!")
}

func deleteExpense(scanner *bufio.Scanner) {
    expenses, err := expense.LoadExpenses()
    if err != nil {
        fmt.Println("Failed to load expenses:", err)
        return
    }

    fmt.Print("Enter expense ID to delete: ")
    if !scanner.Scan() {
        return
    }

    id, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
    if err != nil {
        fmt.Println("Invalid ID.")
        return
    }

    expenses, err = expense.DeleteExpense(expenses, id)
    if err != nil {
        fmt.Println(err)
        return
    }

    if err := expense.SaveExpense(expenses); err != nil {
        fmt.Println("Failed to save expenses:", err)
        return
    }

    fmt.Println("Expense deleted successfully!")
}

func searchExpense(scanner *bufio.Scanner) {
    expenses, err := expense.LoadExpenses()
    if err != nil {
        fmt.Println("Failed to load expenses:", err)
        return
    }

    fmt.Print("Enter keyword: ")
    if !scanner.Scan() {
        return
    }

    keyword := strings.TrimSpace(scanner.Text())
    results := expense.SearchExpenses(expenses, keyword)

    if len(results) == 0 {
        fmt.Println("No matching expenses found.")
        return
    }

    for _, tx := range results {
        fmt.Printf("%d | %s | Rp%.0f | %s\n",
            tx.ID, tx.Description, tx.Amount, tx.Category)
    }
}

func filterExpense(scanner *bufio.Scanner) {
    expenses, err := expense.LoadExpenses()
    if err != nil {
        fmt.Println("Failed to load expenses:", err)
        return
    }

    fmt.Print("Enter category: ")
    if !scanner.Scan() {
        return
    }

    category := strings.TrimSpace(scanner.Text())
    results := expense.FilterExpenses(expenses, category)

    if len(results) == 0 {
        fmt.Println("No expenses found in this category.")
        return
    }

    for _, tx := range results {
        fmt.Printf("%d | %s | Rp%.0f | %s\n",
            tx.ID, tx.Description, tx.Amount, tx.Category)
    }
}

func sortExpenses() {
    expenses, err := expense.LoadExpenses()
    if err != nil {
        fmt.Println("Failed to load expenses:", err)
        return
    }

    expense.SortExpensesByAmount(expenses)

    for _, tx := range expenses {
        fmt.Printf("%d | %s | Rp%.0f | %s\n",
            tx.ID, tx.Description, tx.Amount, tx.Category)
    }
}

func showSummary() {
    expenses, err := expense.LoadExpenses()
    if err != nil {
        fmt.Println("Failed to load expenses:", err)
        return
    }

    if len(expenses) == 0 {
        fmt.Println("No expenses found.")
        return
    }

    summary := expense.CalculateSummary(expenses)

    fmt.Println("\n========== EXPENSE SUMMARY ==========")
    fmt.Printf("Total Expense   : Rp%.0f\n", summary.Total)
    fmt.Printf("Average Expense : Rp%.0f\n", summary.Average)
    fmt.Printf("Highest Expense : Rp%.0f\n", summary.Highest)
    fmt.Printf("Lowest Expense  : Rp%.0f\n", summary.Lowest)
    fmt.Println("=====================================")
}

func exportExpenses() {
    expenses, err := expense.LoadExpenses()
    if err != nil {
        fmt.Println("Failed to load expenses:", err)
        return
    }

    if err := expense.ExportCSV(expenses); err != nil {
        fmt.Println("Failed to export CSV:", err)
        return
    }

    fmt.Println("Expenses exported successfully to expenses.csv!")
}

func main() {
	exportFlag := flag.Bool("export", false, "Export expenses to CSV")
	flag.Parse()

	if *exportFlag {
		expenses, err := expense.LoadExpenses()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		if err := expense.ExportCSV(expenses); err != nil {
			fmt.Println("Export failed:", err)
			return
		}

		fmt.Println("Expenses exported to expenses.csv")
		return
	}

    scanner := bufio.NewScanner(os.Stdin)

    showHeader()

    for {
        showMenu()

        fmt.Print("Choose menu: ")
        if !scanner.Scan() {
            return
        }

        choice := strings.TrimSpace(scanner.Text())

        switch choice {
        case "1":
            expenses, err := expense.LoadExpenses()
            if err != nil {
                expenses = []expense.Expense{}
            }

            id := 1
            if len(expenses) > 0 {
                id = expenses[len(expenses)-1].ID + 1
            }

            transaction := addExpense(scanner, id)
            expenses = append(expenses, transaction)

            if err := expense.SaveExpense(expenses); err != nil {
                fmt.Println("Failed to save:", err)
                continue
            }

            fmt.Println("Expense added successfully!")

        case "2":
            listExpenses()

		case "3":
    		showSummary()

        case "6":
            updateExpense(scanner)

        case "7":
            deleteExpense(scanner)
		
		case "8":
			searchExpense(scanner)
		
		case "9":
			filterExpense(scanner)
		
		case "10":
			sortExpenses()

		case "11":
    		exportExpenses()

        case "4":
            fmt.Println("Thank you for using Expense Tracker!")
            showFooter()
            return

        default:
            fmt.Println("Invalid menu choice.")
        }
    }
}

package main

import "fmt"

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
	fmt.Println(" Developed with Go Standard Lib")
    fmt.Println("================================")
}

func main() {
	fmt.Println("================================")
    fmt.Println("       CLI EXPENSE TRACKER")
    fmt.Println("================================")

    fmt.Println("Welcome to Expense Tracker!")
	showMenu()
	showFooter()
}
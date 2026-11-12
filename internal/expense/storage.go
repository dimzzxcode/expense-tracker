package expense

import (
	"encoding/json"
	"os"
)

func SaveExpense(expense []Expense) error {
	data, err := json.MarshalIndent(expense, "", " ")
	if err != nil {
		return err
	}

	err = os.WriteFile("expenses.json", data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func LoadExpenses() ([]Expense, error) {
	data, err := os.ReadFile("expenses.json")
	if err != nil {
		return nil, err
	}

	var expenses []Expense

	err = json.Unmarshal(data, &expenses)
	if err != nil {
		return nil, err
	}

	return expenses, nil
}
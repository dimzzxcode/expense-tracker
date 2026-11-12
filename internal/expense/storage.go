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
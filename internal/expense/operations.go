package expense

import "errors"

func UpdateExpense(expenses []Expense, id int, updated Expense) ([]Expense, error) {
	for i, tx := range expenses {
		if tx.ID == id {
			updated.ID = id
			expenses[i] = updated
			return expenses, nil
		}
	}

	return nil, errors.New("expense not found")
}

func DeleteExpense(expenses []Expense, id int) ([]Expense, error) {
	for i, tx := range expenses {
		if tx.ID == id {
			expenses = append(expenses[:i], expenses[i+1:]...)
			return expenses, nil
		}
	}

	return nil, errors.New("expense not found")
}
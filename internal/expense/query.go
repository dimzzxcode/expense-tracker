package expense

import (
	"regexp"
	"sort"
	"strings"
)

func SearchExpenses(expenses []Expense, keyword string) []Expense {
    var result []Expense

    keyword = strings.ToLower(keyword)

    for _, tx := range expenses {
        description := strings.ToLower(tx.Description)

        if strings.Contains(description, keyword) {
            result = append(result, tx)
        }
    }

    return result
}

func FilterExpenses(expenses []Expense, category string) []Expense {
    var result []Expense

    for _, tx := range expenses {
        if strings.EqualFold(tx.Category, category) {
            result = append(result, tx)
        }
    }

    return result
}

func SearchExpensesRegex(expenses []Expense, pattern string) ([]Expense, error) {
    regex, err := regexp.Compile(pattern)
    if err != nil {
        return nil, err
    }

    var result []Expense

    for _, tx := range expenses {
        if regex.MatchString(tx.Description) {
            result = append(result, tx)
        }
    }

    return result, nil
}

func SortExpensesByAmount(expenses []Expense) {
    sort.Slice(expenses, func(i, j int) bool {
        return expenses[i].Amount < expenses[j].Amount
    })
}
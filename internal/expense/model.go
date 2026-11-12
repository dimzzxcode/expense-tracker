package expense

import "time"

type Expense struct {
	ID          int
	Description string
	Amount      float64
	Category    string
	Date        time.Time
}

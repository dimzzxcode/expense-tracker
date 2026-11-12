package expense

import "math"

type Summary struct {
    Total   float64
    Average float64
    Highest float64
    Lowest  float64
}

func CalculateSummary(expenses []Expense) Summary {
    var summary Summary

    if len(expenses) == 0 {
        return summary
    }

    summary.Lowest = math.MaxFloat64

    for _, tx := range expenses {
        summary.Total += tx.Amount

        if tx.Amount > summary.Highest {
            summary.Highest = tx.Amount
        }

        if tx.Amount < summary.Lowest {
            summary.Lowest = tx.Amount
        }
    }

    summary.Average = summary.Total / float64(len(expenses))

    return summary
}
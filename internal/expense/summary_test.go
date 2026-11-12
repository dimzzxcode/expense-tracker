package expense

import "testing"

func TestCalculateSummary(t *testing.T) {
    expenses := []Expense{
        {Amount: 100},
        {Amount: 200},
        {Amount: 300},
    }

    result := CalculateSummary(expenses)

    if result.Total != 600 {
        t.Errorf("expected total 600, got %v", result.Total)
    }

    if result.Average != 200 {
        t.Errorf("expected average 200, got %v", result.Average)
    }

    if result.Highest != 300 {
        t.Errorf("expected highest 300, got %v", result.Highest)
    }

    if result.Lowest != 100 {
        t.Errorf("expected lowest 100, got %v", result.Lowest)
    }
}
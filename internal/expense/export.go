package expense

import (
	"encoding/csv"
	"os"
	"strconv"
)

func ExportCSV(expenses []Expense) error {
    file, err := os.Create("expenses.csv")
    if err != nil {
        return err
    }
    defer file.Close()

    writer := csv.NewWriter(file)
    defer writer.Flush()

    header := []string{
        "ID",
        "Description",
        "Amount",
        "Category",
        "Date",
    }

    if err := writer.Write(header); err != nil {
        return err
    }

    for _, tx := range expenses {
        record := []string{
            strconv.Itoa(tx.ID),
            tx.Description,
            strconv.FormatFloat(tx.Amount, 'f', 2, 64),
            tx.Category,
            tx.Date.Format("2006-01-02"),
        }

        if err := writer.Write(record); err != nil {
            return err
        }
    }

    return writer.Error()
}
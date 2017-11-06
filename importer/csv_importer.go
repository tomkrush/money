package importer

import (
	"encoding/csv"
	"io"
	"money/finance"
	"os"
	"strconv"
	"time"
)

// ImportTransactionsCSV accepts a file path be imported as Transaction types.
func TransactionsCSV(path string) (finance.Transactions, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	reader := csv.NewReader(file)

	transactions := make([]finance.Transaction, 0)

	skippedHeader := false

	for {
		record, err := reader.Read()

		if skippedHeader == false {
			skippedHeader = true
			continue
		}

		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		transaction := finance.Transaction{}
		transaction.BankID = record[0]
		transaction.AccountNumber = record[1]
		transaction.AccountType = record[2]

		transaction.Balance = finance.NewCurrencyFromDollars(record[3])

		transaction.StartDate, err = time.Parse("2006-01-02", record[4])

		if err != nil {
			return nil, err
		}

		transaction.EndDate, err = time.Parse("2006-01-02", record[5])

		if err != nil {
			return nil, err
		}

		transaction.Type = record[6]

		transaction.Date, err = time.Parse("2006-01-02", record[7])

		if err != nil {
			return nil, err
		}

		transaction.Amount = finance.NewCurrencyFromDollars(record[8])

		uniqueID, err := strconv.ParseInt(record[9], 10, 32)

		if err != nil {
			return nil, err
		}

		transaction.UniqueID = int(uniqueID)
		transaction.Description = record[10]

		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

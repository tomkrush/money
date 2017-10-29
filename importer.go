package main

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"
	"time"
)

func importTransactions(path string) ([]Transaction, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	reader := csv.NewReader(file)

	transactions := make([]Transaction, 0)

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

		transaction := Transaction{}
		transaction.BankID = record[0]
		transaction.AccountNumber = record[1]
		transaction.AccountType = record[2]
		balance, err := strconv.ParseFloat(record[3], 64)
		transaction.Balance = int(balance * 100)

		if err != nil {
			return nil, err
		}

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

		amount, err := strconv.ParseFloat(record[8], 64)
		transaction.Amount = int(amount * 100)

		if err != nil {
			return nil, err
		}

		transaction.UniqueID = record[9]
		transaction.Description = record[10]

		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

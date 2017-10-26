package main

import (
	"encoding/csv"
	"io"
	"os"
)

func importTransactions(path string) ([]Transaction, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	reader := csv.NewReader(file)

	transactions := make([]Transaction, 0)

	for {
		_, err := reader.Read()

		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		transactions = append(transactions, Transaction{})
	}

	return transactions, nil
}

package main

import (
	"encoding/csv"
	"io"
	"os"
)

func importTransactions(path string) ([][]string, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	reader := csv.NewReader(file)

	records := make([]([]string), 0)

	for {
		record, err := reader.Read()

		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		records = append(records, record)
	}

	return records, nil
}

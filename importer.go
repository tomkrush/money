package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func importTransactions(path string) [][]string {
	file, err := os.Open(path)
	records := make([][]string, 0)

	if err != nil {
		fmt.Println("Error: ", err)
		return records
	}

	defer file.Close()

	reader := csv.NewReader(file)

	for {
		record, err := reader.Read()

		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error: ", err)
		}

		records = append(records, record)
	}

	return records
}

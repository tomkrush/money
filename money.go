package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	path := flag.String("path", "", "Import CSV file from specified path")
	flag.Parse()

	file, err := os.Open(*path)

	if err != nil {
		fmt.Println("Error: ", err)
		return
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

		fmt.Println("Record", record)
	}
}

package main

import (
	"flag"
	"fmt"
)

func main() {
	path := flag.String("path", "", "Import CSV file from specified path")
	flag.Parse()

	records, _ := importTransactionsCSV(*path)
	fmt.Println(records)
}

package main

import (
	"flag"
	"fmt"
)

func main() {
	path := flag.String("path", "", "Import CSV file from specified path")
	flag.Parse()

	transactions, _ := ImportTransactionsCSV(*path)
	transactions.Sort()
	ledger := Ledger{transactions}
	sum := ledger.Sum()

	fmt.Println("Account Total: ", sum.FormatToDollars())
}

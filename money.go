package main

import (
	"flag"
)

func main() {
	path := flag.String("path", "", "Import CSV file from specified path")
	flag.Parse()

	transactions, _ := ImportTransactionsCSV(*path)
	transactions.Sort()
	// ledger := Ledger{transactions}
	// startingBalance := ledger.StartingBalance()
	// sum := ledger.Sum()

	// fmt.Println("Starting Balance: ", startingBalance.FormatToDollars())
	// fmt.Println("Account Total: ", sum.FormatToDollars())
}

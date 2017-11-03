package main

import (
	"flag"
	"fmt"
)

func main() {
	path := flag.String("path", "", "Import CSV file from specified path")
	flag.Parse()

	transactions, _ := ImportTransactionsCSV(*path)
	// transactions.Sort()
	accounts := transactions.SplitIntoAccounts()

	for accountNumber, account := range accounts {
		amount := account.Sum()
		starting := account.StartingBalance()

		// if accountNumber == "5555555555" {
		// 	for _, t := range account.Transactions {
		// 		fmt.Println(t.Balance.FormatToDollars(), t.Amount.FormatToDollars())
		// 	}
		// }

		fmt.Println(accountNumber, starting.FormatToDollars(), amount.FormatToDollars())
	}

	// ledger := Ledger{transactions}
	// startingBalance := ledger.StartingBalance()
	// sum := ledger.Sum()

	// fmt.Println("Starting Balance: ", startingBalance.FormatToDollars())
	// fmt.Println("Account Total: ", sum.FormatToDollars())
}

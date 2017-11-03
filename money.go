package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

func main() {
	path := flag.String("path", "", "Import CSV file from specified path")
	flag.Parse()

	transactions, _ := ImportTransactionsCSV(*path)
	accounts := transactions.SplitIntoAccounts()

	fmt.Println("# Account Balances")
	fmt.Println()

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Account Number", "Starting Balance", "Ending Balance"})
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.SetAlignment(tablewriter.ALIGN_LEFT)

	for accountNumber, account := range accounts {
		amount := account.Sum()
		starting := account.StartingBalance()

		table.Append([]string{accountNumber, starting.FormatToDollars(), amount.FormatToDollars()})
	}

	table.Render() // Send output
}

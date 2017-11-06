package main

import (
	"flag"
	"fmt"
	"money/importer"
	"money/rules"
	"os"

	"github.com/olekukonko/tablewriter"
)

func main() {
	transactionsPath := flag.String("path", "", "Import CSV file from specified path")
	rulesPath := flag.String("rules", "", "Import json file of rules")

	flag.Parse()

	personalRules := rules.New(*rulesPath)
	transactions, _ := importer.TransactionsCSV(*transactionsPath)

	transactions = personalRules.Apply(transactions)

	for _, t := range transactions {
		fmt.Println(t.GetCategory(), t.GetDescription(), t.Amount.FormatToDollars())
	}

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

	starting := accounts.StartingBalance()
	table.SetFooter([]string{"–", starting.FormatToDollars(), "–"})

	table.Render() // Send output
}

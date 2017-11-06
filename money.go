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

	fmt.Println("# Personal Balance")
	fmt.Println()

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Starting Balance", "Ending Balance"})
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.SetAlignment(tablewriter.ALIGN_LEFT)

	starting := transactions.StartingBalance()
	amount := transactions.Sum()
	table.Append([]string{starting.FormatToDollars(), amount.FormatToDollars()})

	table.Render() // Send output
}

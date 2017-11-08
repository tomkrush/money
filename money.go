package main

import (
	"flag"
	"money/importer"
	"money/reports"
	"money/rules"
	"time"
)

func main() {
	transactionsPath := flag.String("path", "", "Import CSV file from specified path")
	rulesPath := flag.String("rules", "", "Import json file of rules")

	flag.Parse()

	personalRules := rules.New(*rulesPath)
	transactions, _ := importer.TransactionsCSV(*transactionsPath)

	transactions.Sort()

	transactions = personalRules.Apply(transactions)

	start, _ := time.Parse("2006-01-02", "2017-10-01")
	end, _ := time.Parse("2006-01-02", "2017-10-31")

	transactions = transactions.DateRange(start, end)

	// transactions = transactions.FilterByCategory("Uncategorized")
	reports.Transactions(transactions)
	// reports.Categories(transactions)
}

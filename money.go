package main

import (
	"flag"
	"fmt"
	"money/importer"
	"money/rules"
	"time"
)

func main() {
	transactionsPath := flag.String("path", "", "Import CSV file from specified path")
	rulesPath := flag.String("rules", "", "Import json file of rules")

	flag.Parse()

	rules := rules.New(*rulesPath)
	transactions, _ := importer.TransactionsCSV(*transactionsPath)

	transactions.Sort()

	transactions = rules.Apply(transactions)

	start, _ := time.Parse("2006-01-02", "2017-11-01")
	end, _ := time.Parse("2006-01-02", "2017-11-30")

	transactions = transactions.DateRange(start, end)

	bills := rules.Bills(transactions)

	// transactions = transactions.FilterByCategory("Uncategorized")
	// reports.Transactions(transactions)
	// reports.Bills(bills)

	other := transactions.TotalExpenses().Subtract(bills.ActualAmount())
	income := rules.Income.Amount
	projectedAmount := bills.ProjectedAmount()
	allowance := income.Add(other.Add(projectedAmount))
	actualAmount := bills.ActualAmount()
	sum := transactions.TotalExpenses()

	fmt.Println("Sum: ", sum.FormatToDollars())
	fmt.Println("Income: ", income.FormatToDollars())
	fmt.Println("Actual Bills: ", actualAmount.FormatToDollars())
	fmt.Println("Projected Bills: ", projectedAmount.FormatToDollars())
	fmt.Println("Other Expenses: ", other.FormatToDollars())

	fmt.Printf("Allowance %s\n\n", allowance.FormatToDollars())

	// reports.Categories(transactions)
}

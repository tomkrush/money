package main

import (
	"flag"
	"fmt"
	"money/importer"
	"money/reports"
	"money/rules"
	"time"
)

func main() {
	transactionsPath := flag.String("path", "", "Import CSV file from specified path")
	rulesPath := flag.String("rules", "", "Import json file of rules")

	flag.Parse()

	transactionRules := rules.New(*rulesPath)
	transactions, _ := importer.TransactionsCSV(*transactionsPath)

	transactions.Sort()

	transactions = transactionRules.Apply(transactions)

	start, _ := time.Parse("2006-01-02", "2017-11-01")
	end, _ := time.Parse("2006-01-02", "2017-11-30")

	transactions = transactions.DateRange(start, end)

	bills := transactionRules.Bills(transactions)

	// transactions = transactions.FilterByCategory("Uncategorized")
	// reports.Transactions(transactions)

	income := transactionRules.Income

	allowance := rules.Allowance(income, bills, transactions)

	unplannedExpense := rules.UnplannedExpenses(bills, transactions)
	projectedAmount := bills.ProjectedAmount()
	actualAmount := bills.ActualAmount()
	sum := transactions.TotalExpenses()

	fmt.Println("# Personal Finances")

	fmt.Println("### Summary")
	fmt.Println("- Sum: ", sum.FormatToDollars())
	fmt.Println("- Income: ", income.Amount.FormatToDollars())
	fmt.Println("- Actual Bills: ", actualAmount.FormatToDollars())
	fmt.Println("- Projected Bills: ", projectedAmount.FormatToDollars())
	fmt.Println("- Unplanned Expenses: ", unplannedExpense.FormatToDollars())

	fmt.Printf("Allowance %s\n\n", allowance.FormatToDollars())

	reports.Bills(bills)

	reports.Categories(transactions)
}

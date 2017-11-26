package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/tomkrush/money/importer"
	"github.com/tomkrush/money/rules"
)

type Path struct {
	TransactionPath string
	RulesPath       string
}

type Data struct {
	LastUpdated       string
	UnplannedExpenses string
	Allowance         string
	ProjectedBills    string
	RemainingBills    string
	Savings           string
	Categories        rules.CategoryList
	Bills             []rules.Bill
}

func indexHandler(writer http.ResponseWriter, request *http.Request, paths Path) {
	transactionRules := rules.New(paths.RulesPath)
	transactions, _ := importer.TransactionsCSV(paths.TransactionPath)

	fileInfo, err := os.Stat(paths.TransactionPath)

	if err != nil {
		log.Fatal(err)
	}

	lastUpdated := fileInfo.ModTime().Format("2006-01-02")

	if err != nil {
		log.Fatal(err)
	}

	transactions.Sort()

	transactions = transactionRules.Apply(transactions)

	start, _ := time.Parse("2006-01-02", "2017-11-01")
	end, _ := time.Parse("2006-01-02", "2017-11-30")

	transactions = transactions.DateRange(start, end)

	bills := transactionRules.Bills(transactions)

	income := transactionRules.TotalIncome(transactions)

	allowance := rules.Allowance(income.ProjectedAmount, bills, transactions)

	unplannedExpense := rules.UnplannedExpenses(bills, transactions)
	projectedAmount := bills.ProjectedAmount()
	remainingAmount := bills.RemainingAmount()

	data := Data{
		LastUpdated:       lastUpdated,
		UnplannedExpenses: unplannedExpense.FormatToDollars(),
		Allowance:         allowance.FormatToDollars(),
		ProjectedBills:    projectedAmount.FormatToDollars(),
		RemainingBills:    remainingAmount.FormatToDollars(),
		Categories:        rules.GetCategories(transactions),
		Savings:           "N/A",
		Bills:             bills.List(),
	}

	t, err := template.ParseFiles("templates/index.html")

	err = t.Execute(writer, data) // merge.

	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	transactionsPath := flag.String("path", "", "Import CSV file from specified path")
	rulesPath := flag.String("rules", "", "Import json file of rules")

	flag.Parse()

	paths := Path{
		TransactionPath: *transactionsPath,
		RulesPath:       *rulesPath,
	}

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		indexHandler(writer, request, paths)
	})

	fileServer := http.StripPrefix("/css", http.FileServer(http.Dir("css")))
	http.Handle("/css/", fileServer)

	err := http.ListenAndServe(":9696", nil) // set listen port

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

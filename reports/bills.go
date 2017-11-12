package reports

import (
	"fmt"
	"money/finance"
	"money/rules"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
)

func getTransactionForBill(transactions finance.Transactions, rule rules.TransactionRule) (finance.Transaction, bool) {
	for _, transaction := range transactions {
		if transaction.GetDescription() == rule.Bill.Description {
			return transaction, true
		}
	}

	return finance.Transaction{}, false
}

func abs(value int) int {
	if value < 0 {
		return value * -1
	}

	return value
}

// Bills displays a list of all bills provided by the rules.
// Each list item displays the what was actually spent on the bill
// and how much should have been spent on the bill.
//
// The purpose of this view is to provide a simple summary of what bills
// have been handled and which ones still need to be handled.
//
// Another important value in this view is the projected bill spend.
// The projected bill spend is the total amount spent on bills so far based
// on transaction history plus the total amount remaining on bills. This value
// can change over the course of the month. It is possible to have a lower projected
// goal if you spend less on your bills than expected. It is pssible to have a higher
// projected goal if you spend more on a bill. Variable bills such as an Electric
// or Gas bill are a good reason why the projected value can change.
func Bills(bills rules.Bills) {
	fmt.Println("# Bills")
	fmt.Println()

	transactions := bills.Transactions

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Bill", "Pay Date", "Transaction Date", "Estimated Amount", "Actual Amount", "Need"})
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.SetAlignment(tablewriter.ALIGN_LEFT)

	actualTotalAmount := 0
	projectedTotalAmount := 0

	for _, rule := range bills.Rules {
		need := "No"
		actualAmount := "–"
		transaction, ok := getTransactionForBill(transactions, rule)
		transactionDate := "–"

		if ok {
			actualAmount = transaction.Amount.FormatToDollars()
			transactionDate = transaction.Date.Format("2006-01-02")
			actualTotalAmount += abs(transaction.Amount.Amount)
			projectedTotalAmount += abs(transaction.Amount.Amount)
		} else {
			projectedTotalAmount += rule.Bill.Amount.Amount
		}

		if rule.Need {
			need = "Yes"
		}

		day := strconv.FormatInt(int64(rule.Bill.Day), 10)

		table.Append([]string{
			rule.Bill.Description,
			day,
			transactionDate,
			rule.Bill.Amount.FormatToDollars(),
			actualAmount,
			need,
		})
	}

	table.Render() // Send output

	estimatedTotalAmount := bills.GoalAmount()
	formattedActualAmount := finance.NewCurrency(actualTotalAmount)
	formattedProjectedAmount := finance.NewCurrency(projectedTotalAmount)
	remainingAmount := finance.NewCurrency(projectedTotalAmount - actualTotalAmount)

	fmt.Printf("\nGoal Amount: %s    Already Paid: %s    Projected Amount: %s    Remaining Amount: %s\n\n",
		estimatedTotalAmount.FormatToDollars(),
		formattedActualAmount.FormatToDollars(),
		formattedProjectedAmount.FormatToDollars(),
		remainingAmount.FormatToDollars(),
	)

}

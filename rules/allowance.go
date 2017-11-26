package rules

import "github.com/tomkrush/money/finance"

// Allowance is calculated based on the (total income + projected bill total
// + all other expenses)
func Allowance(income finance.Currency, bills Bills, transactions finance.Transactions) finance.Currency {
	unplannedExpense := UnplannedExpenses(bills, transactions)
	projectedAmount := bills.ProjectedAmount()

	allowance := income.Add(unplannedExpense.Add(projectedAmount))

	return allowance
}

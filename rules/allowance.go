package rules

import "money/finance"

// Allowance is calculated based on the (total income + projected bill total
// + all other expenses)
func Allowance(income Income, bills Bills, transactions finance.Transactions) finance.Currency {
	unplannedExpense := UnplannedExpenses(bills, transactions)
	projectedAmount := bills.ProjectedAmount()

	allowance := income.Amount.Add(unplannedExpense.Add(projectedAmount))

	return allowance
}

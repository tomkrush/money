package rules

import "money/finance"

// UnplannedExpenses returns the total amount of money that has been spent
// that isn't a bill or planned transaction.
func UnplannedExpenses(bills Bills, transactions finance.Transactions) finance.Currency {
	return transactions.TotalExpenses().Subtract(bills.ActualAmount())
}

package rules

import "money/finance"

// Allowance is calculated based on the (total income + projected bill total
// + all other expenses)
func Allowance(income Income, bills Bills, transactions finance.Transactions) finance.Currency {
	other := transactions.TotalExpenses().Subtract(bills.ActualAmount())
	projectedAmount := bills.ProjectedAmount()
	allowance := income.Amount.Add(other.Add(projectedAmount))

	return allowance
}

package rules

import "money/finance"

type Allowance struct {
	Bills  Bills
	Income Income
}

func (a Allowance) Monthly() finance.Currency {
	remaining := a.Income.Amount.Subtract(a.Bills.ProjectedAmount())

	return remaining
}

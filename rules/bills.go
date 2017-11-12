package rules

import (
	"money/finance"
)

// Bills contains a list of bill rules
type Bills struct {
	Rules []TransactionRule
}

func (bills Bills) GoalAmount() finance.Currency {
	amount := 0

	for _, rule := range bills.Rules {
		amount += rule.Bill.Amount.Amount
	}

	return finance.NewCurrency(amount)
}

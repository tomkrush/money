package rules

import (
	"money/finance"
)

// Bills contains a list of bill rules
type Bills struct {
	Rules        []TransactionRule
	Transactions finance.Transactions
	calculated   bool
	goalAmount   finance.Currency
}

func (b *Bills) Calculate() {
	if b.calculated == false {
		amount := 0

		for _, rule := range b.Rules {
			amount += rule.Bill.Amount.Amount
		}

		b.goalAmount = finance.NewCurrency(amount)
	}
}

func (b Bills) GoalAmount() finance.Currency {
	return b.goalAmount
}

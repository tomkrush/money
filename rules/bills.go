package rules

import (
	"money/finance"
)

func abs(value int) int {
	if value < 0 {
		return value * -1
	}

	return value
}

// Bills contains a list of bill rules
type Bills struct {
	Rules           []TransactionRule
	Transactions    finance.Transactions
	calculated      bool
	goalAmount      finance.Currency
	actualAmount    finance.Currency
	projectedAmount finance.Currency
	remainingAmount finance.Currency
}

func (b *Bills) Calculate() {
	if b.calculated == false {
		goalAmount := 0
		actualAmount := 0
		projectedAmount := 0

		for _, rule := range b.Rules {
			transaction, ok := b.getTransaction(rule)

			if ok {
				actualAmount += abs(transaction.Amount.Amount)
				projectedAmount += abs(transaction.Amount.Amount)
			} else {
				projectedAmount += rule.Bill.Amount.Amount
			}

			goalAmount += rule.Bill.Amount.Amount
		}

		b.goalAmount = finance.NewCurrency(goalAmount)
		b.actualAmount = finance.NewCurrency(actualAmount)
		b.projectedAmount = finance.NewCurrency(projectedAmount)
		b.remainingAmount = finance.NewCurrency(projectedAmount - actualAmount)
	}
}

func (b Bills) ProjectedAmount() finance.Currency {
	return b.projectedAmount
}

func (b Bills) RemainingAmount() finance.Currency {
	return b.remainingAmount
}

func (b Bills) GoalAmount() finance.Currency {
	return b.goalAmount
}

func (b Bills) ActualAmount() finance.Currency {
	return b.actualAmount
}

func (b *Bills) getTransaction(rule TransactionRule) (finance.Transaction, bool) {
	for _, transaction := range b.Transactions {
		if transaction.GetDescription() == rule.Bill.Description {
			return transaction, true
		}
	}

	return finance.Transaction{}, false
}

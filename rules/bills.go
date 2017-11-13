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

// Calculate iterates over bill rules and transactions and internally holds
// the results for future use.
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

// ProjectedAmount returns the amount of money that is going to be spent
// this month based on the actual amount of money already spent on bills plus
// the the remaining amount of unpaid bills.
func (b Bills) ProjectedAmount() finance.Currency {
	return b.projectedAmount
}

// RemainingAmount returns the amount of money that has yet to be paid to bills.
func (b Bills) RemainingAmount() finance.Currency {
	return b.remainingAmount
}

// GoalAmount returns the amount of money that will ideally be spent on bills
// this month. This is calculated by adding up the expected bill amounts.
func (b Bills) GoalAmount() finance.Currency {
	return b.goalAmount
}

// ActualAmount returns the amount of money that has already been spent on bills.
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

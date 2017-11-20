package rules

import (
	"money/finance"
	"strconv"
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
	bills           []Bill
}

// Bill holds the summarized info for a user to understand if the bill has
// been paid, and if so what amount.
type Bill struct {
	Day         string
	Amount      string
	Description string
	Paid        string
}

// Calculate iterates over bill rules and transactions and internally holds
// the results for future use.
func (b *Bills) Calculate() {
	if b.calculated == false {
		goalAmount := 0
		actualAmount := 0
		projectedAmount := 0
		var bills []Bill

		for _, rule := range b.Rules {
			transaction, ok := b.Transactions.GetByDescription(rule.Bill.Description)

			billedAmount := rule.Bill.Amount
			paid := "estimate"

			if ok {
				actualAmount -= abs(transaction.Amount.Amount)
				projectedAmount -= abs(transaction.Amount.Amount)
				billedAmount = transaction.Amount
				paid = "expense"
			} else {
				projectedAmount -= rule.Bill.Amount.Amount
			}

			day := strconv.FormatInt(int64(rule.Bill.Day), 10)

			bills = append(bills, Bill{
				Description: rule.Bill.Description,
				Day:         day,
				Amount:      billedAmount.FormatToDollars(),
				Paid:        paid,
			})

			goalAmount -= rule.Bill.Amount.Amount
		}

		b.goalAmount = finance.NewCurrency(goalAmount)
		b.actualAmount = finance.NewCurrency(actualAmount)
		b.projectedAmount = finance.NewCurrency(projectedAmount)
		b.remainingAmount = finance.NewCurrency(projectedAmount - actualAmount)
		b.bills = bills
		b.calculated = true
	}
}

// List returns a list of bills that either have been or haven't been paid.
// Depending on paid status, the bill will have estimated or actual values.
func (b *Bills) List() []Bill {
	b.Calculate()
	return b.bills
}

// ProjectedAmount returns the amount of money that is going to be spent
// this month based on the actual amount of money already spent on bills plus
// the the remaining amount of unpaid bills.
func (b *Bills) ProjectedAmount() finance.Currency {
	b.Calculate()
	return b.projectedAmount
}

// RemainingAmount returns the amount of money that has yet to be paid to bills.
func (b *Bills) RemainingAmount() finance.Currency {
	b.Calculate()
	return b.remainingAmount
}

// GoalAmount returns the amount of money that will ideally be spent on bills
// this month. This is calculated by adding up the expected bill amounts.
func (b *Bills) GoalAmount() finance.Currency {
	b.Calculate()
	return b.goalAmount
}

// ActualAmount returns the amount of money that has already been spent on bills.
func (b *Bills) ActualAmount() finance.Currency {
	b.Calculate()
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

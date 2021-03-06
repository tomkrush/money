package rules

import (
	"github.com/tomkrush/money/finance"
	"testing"
)

func createBillRule(description string, amount int) TransactionRule {
	return TransactionRule{
		Bill: BillRule{
			Description: description,
			Day:         1,
			Amount:      finance.NewCurrency(amount),
		},
	}
}

func TestRules_Bills(t *testing.T) {
	bills := Bills{
		Rules: []TransactionRule{
			createBillRule("hello", 1000),
			createBillRule("world", 750),
			TransactionRule{},
		},
		Transactions: finance.Transactions{
			finance.Transaction{
				Description: "hello",
				Amount:      finance.NewCurrency(1200),
			},
			finance.Transaction{
				Description: "test",
				Amount:      finance.NewCurrency(1500),
			},
		},
	}

	goalAmount := bills.GoalAmount()
	actualAmount := bills.ActualAmount()
	projectedAmount := bills.ProjectedAmount()
	remainingAmount := bills.RemainingAmount()

	if goalAmount.Amount != -1750 {
		t.Error("Goal amount incorrect")
	}

	if actualAmount.Amount != -1200 {
		t.Errorf("Actual amount incorrect %d, %d", actualAmount, -1200)
	}

	if projectedAmount.Amount != -1950 {
		t.Errorf("Projected amount incorrect %d, %d", projectedAmount, -1950)
	}

	if remainingAmount.Amount != -750 {
		t.Errorf("Remaining amount incorrect %d, %d", projectedAmount, -750)
	}
}

func TestRules_Bill(t *testing.T) {
	bills := Bills{
		Rules: []TransactionRule{
			createBillRule("hello", 1000),
		},
		Transactions: finance.Transactions{
			finance.Transaction{
				Description: "hello",
				Amount:      finance.NewCurrency(1200),
			},
		},
	}

	list := bills.List()

	bill := list[0]

	if bill.Description != "hello" {
		t.Error("Description Incorrect")
	}

	if bill.Amount != "$12.00" {
		t.Error("Amount Incorrect")
	}

	if bill.Paid != "expense" {
		t.Error("Should be expense")
	}

	if bill.Day != "1" {
		t.Error("Day is incorrect")
	}
}

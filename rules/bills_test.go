package rules

import (
	"money/finance"
	"testing"
)

func createBillRule(description string, amount int) TransactionRule {
	return TransactionRule{
		Bill: BillRule{
			Description: description,
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
				Amount:      finance.NewCurrency(1000),
			},
			finance.Transaction{
				Description: "test",
				Amount:      finance.NewCurrency(1500),
			},
		},
	}

	bills.Calculate()

	goalAmount := bills.GoalAmount()
	actualAmount := bills.ActualAmount()

	if goalAmount.Amount != 1750 {
		t.Error("Goal amount incorrect")
	}

	if actualAmount.Amount != 1000 {
		t.Errorf("Actual amount incorrect %d, %d", actualAmount, 1000)
	}
}

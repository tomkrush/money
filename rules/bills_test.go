package rules

import (
	"money/finance"
	"testing"
)

func createBillRule(amount int) TransactionRule {
	return TransactionRule{
		Bill: BillRule{
			Amount: finance.NewCurrency(amount),
		},
	}
}

func TestRules_Bills(t *testing.T) {
	bills := Bills{
		Rules: []TransactionRule{
			createBillRule(1000),
			createBillRule(750),
			TransactionRule{},
		},
	}

	bills.Calculate()

	goalAmount := bills.GoalAmount()

	if goalAmount.Amount != 1750 {
		t.Error("Goal amount incorrect")
	}
}

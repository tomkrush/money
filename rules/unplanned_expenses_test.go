package rules

import (
	"money/finance"
	"testing"
)

func TestUnplannedExpense_Expenses(t *testing.T) {
	transactions := finance.Transactions{
		finance.Transaction{
			Description: "hello",
			Amount:      finance.NewCurrency(-1200),
		},
		finance.Transaction{
			Description: "test",
			Amount:      finance.NewCurrency(-1500),
		},
	}

	bills := Bills{
		Rules: []TransactionRule{
			createBillRule("hello", 1000),
		},
		Transactions: transactions,
	}

	unplannedExpense := UnplannedExpenses(bills, transactions)

	if unplannedExpense.Amount != -1500 {
		t.Errorf("Unplanned expense incorrect %d %d", unplannedExpense.Amount, -1500)
	}
}

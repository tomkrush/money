package rules

import (
	"money/finance"
	"testing"
)

func TestAllowance_IncomeOnly(t *testing.T) {
	income := finance.NewCurrency(100)

	bills := Bills{}
	transactions := finance.Transactions{}

	allowance := Allowance(income, bills, transactions)

	if allowance.Amount != 100 {
		t.Errorf("Allowance incorrect %d %d", allowance.Amount, 100)
	}
}

func TestAllowance_Bills(t *testing.T) {
	income := finance.NewCurrency(10000)

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
			createBillRule("world", 750),
		},
		Transactions: transactions,
	}

	allowance := Allowance(income, bills, transactions)

	if allowance.Amount != 6550 {
		t.Errorf("Allowance incorrect %d %d", allowance.Amount, 6550)
	}
}

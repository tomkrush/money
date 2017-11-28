package rules

import (
	"testing"

	"github.com/tomkrush/money/finance"
)

func TestAllowance_IncomeOnly(t *testing.T) {
	income := finance.NewCurrency(100)

	bills := Bills{}
	transactions := finance.Transactions{}
	plannedExpenses := finance.NewCurrency(-10)

	allowance := Allowance(income, plannedExpenses, bills, transactions)

	if allowance.Amount != 90 {
		t.Errorf("Allowance incorrect %d %d", allowance.Amount, 90)
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

	plannedExpenses := finance.NewCurrency(0)

	bills := Bills{
		Rules: []TransactionRule{
			createBillRule("hello", 1000),
			createBillRule("world", 750),
		},
		Transactions: transactions,
	}

	allowance := Allowance(income, plannedExpenses, bills, transactions)

	if allowance.Amount != 6550 {
		t.Errorf("Allowance incorrect %d %d", allowance.Amount, 6550)
	}
}

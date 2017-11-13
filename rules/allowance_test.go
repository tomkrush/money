package rules

import (
	"money/finance"
	"testing"
)

func TestAllowance_Allowance(t *testing.T) {
	allowance := Allowance{
		Income: Income{
			Amount: finance.NewCurrency(100),
		},
	}

	monthly := allowance.Monthly()

	if monthly.Amount != 100 {
		t.Errorf("Allowance incorrect %d %d", monthly.Amount, 100)
	}
}

func TestAllowance_IncludingBills(t *testing.T) {
	allowance := Allowance{
		Income: Income{
			Amount: finance.NewCurrency(3000),
		},
		Bills: Bills{
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
		},
	}

	monthly := allowance.Monthly()

	if monthly.Amount != 1050 {
		t.Errorf("Allowance incorrect %d %d", monthly.Amount, 1050)
	}
}

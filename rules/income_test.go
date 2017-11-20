package rules

import (
	"money/finance"
	"testing"
)

func TestRules_getIncomeRuleWithDescription(t *testing.T) {
	rules := []TransactionRule{
		TransactionRule{
			Income: Paycheck{
				Description: "Test",
			},
		},
		TransactionRule{
			Income: Paycheck{
				Description: "Test #2",
			},
		},
	}

	_, ok := getIncomeRuleWithDescription(rules, "Test")

	if ok == false {
		t.Error("Income Rule not found")
	}
}

func TestRules_getIncomeRuleWithDescriptionFail(t *testing.T) {
	rules := []TransactionRule{
		TransactionRule{
			Income: Paycheck{
				Description: "Test",
			},
		},
		TransactionRule{
			Income: Paycheck{
				Description: "Test #2",
			},
		},
	}

	_, ok := getIncomeRuleWithDescription(rules, "Foobar")

	if ok == true {
		t.Error("Income Rule should not have been found")
	}
}

func TestTransaction_GetDescription(t *testing.T) {
	rules := Rules{
		Transactions: []TransactionRule{
			TransactionRule{
				Contains: "Work",
				Income: Paycheck{
					Description: "Work Income",
					Frequency:   1,
					Amount:      finance.NewCurrency(1000),
				},
			},
		},
	}

	tests := []struct {
		name         string
		transactions finance.Transactions
		want         IncomeSummary
	}{
		{
			"Add up remaining income",
			finance.Transactions{
				finance.Transaction{},
			},
			IncomeSummary{
				ProjectedAmount: finance.NewCurrency(1000),
				RemainingAmount: finance.NewCurrency(1000),
			},
		},
		{
			"Add up all income",
			finance.Transactions{
				finance.Transaction{
					Description: "Work Income",
					Amount:      finance.NewCurrency(1100),
					Income:      true,
				},
			},
			IncomeSummary{
				ProjectedAmount: finance.NewCurrency(1100),
				ActualAmount:    finance.NewCurrency(1100),
			},
		},
		{
			"Unplanned  income",
			finance.Transactions{
				finance.Transaction{
					Description: "Random Income",
					Amount:      finance.NewCurrency(800),
					Income:      true,
				},
			},
			IncomeSummary{
				ProjectedAmount: finance.NewCurrency(1800),
				ActualAmount:    finance.NewCurrency(800),
				RemainingAmount: finance.NewCurrency(1000),
				UnPlannedAmount: finance.NewCurrency(800),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			summary := rules.TotalIncome(tt.transactions)

			if tt.want != summary {
				t.Errorf("TotalIncome returned %v, want %v", summary, tt.want)
			}
		})
	}
}

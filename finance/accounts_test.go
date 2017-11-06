package finance

import (
	"reflect"
	"testing"
)

func TestAccounts_StartingBalance(t *testing.T) {
	tests := []struct {
		name     string
		accounts Accounts
		want     Currency
	}{
		{
			"Starting balance is 0",
			Accounts{
				"1": Account{
					Transactions: Transactions{
						Transaction{Amount: NewCurrency(500), Balance: NewCurrency(600)},
						Transaction{Amount: NewCurrency(500), Balance: NewCurrency(1000)},
					},
				},
				"2": Account{
					Transactions: Transactions{
						Transaction{Amount: NewCurrency(250), Balance: NewCurrency(1311)},
						Transaction{Amount: NewCurrency(300), Balance: NewCurrency(300)},
					},
				},
			},
			NewCurrency(1161),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := tt.accounts
			if got := a.StartingBalance(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Accounts.StartingBalance() = %v, want %v", got, tt.want)
			}
		})
	}
}

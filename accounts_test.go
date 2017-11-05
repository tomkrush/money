package main

import (
	"money/currency"
	"reflect"
	"testing"
)

func TestAccounts_StartingBalance(t *testing.T) {
	tests := []struct {
		name     string
		accounts Accounts
		want     currency.Currency
	}{
		{
			"Starting balance is 0",
			Accounts{
				"1": Account{
					Transactions: Transactions{
						Transaction{Amount: currency.New(500), Balance: currency.New(600), AccountNumber: "1"},
						Transaction{Amount: currency.New(500), Balance: currency.New(1000), AccountNumber: "1"},
					},
				},
				"2": Account{
					Transactions: Transactions{
						Transaction{Amount: currency.New(250), Balance: currency.New(1311), AccountNumber: "2"},
						Transaction{Amount: currency.New(300), Balance: currency.New(300), AccountNumber: "2"},
					},
				},
			},
			currency.New(1161),
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

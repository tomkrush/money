package currency

import (
	"testing"
)

func TestCurrency_CentsToDollars(t *testing.T) {
	type fields struct {
		Amount int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"Convert 1 penny to $0.00", fields{1}, "$0.01"},
		{"Convert 120 pennies to $1.20", fields{120}, "$1.20"},
		{"Convert 1005 pennies to $10.05", fields{1005}, "$10.05"},
		{"Convert 6050400 pennies to $60504.00", fields{6050400}, "$60504.00"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			currency := &Currency{
				Amount: tt.fields.Amount,
			}
			if got := currency.FormatToDollars(); got != tt.want {
				t.Errorf("Currency.CentsToDollars() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCurrency_FromDollars(t *testing.T) {
	type args struct {
		dollars int
	}
	tests := []struct {
		name    string
		dollars string
		want    int
	}{
		{"Convert 2193.00", "2193.00", 219300},
		{"Convert 2193.97", "2193.97", 219397},
		{"Convert 2193.97", "2193", 219300},
		{"Convert 2193.97", "-1046", -104600},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Currency{}
			c.FromDollars(tt.dollars)

			if got := c.Amount; got != tt.want {
				t.Errorf("Currency.FromDollars() = %v, want %v", got, tt.want)
			}
		})
	}
}

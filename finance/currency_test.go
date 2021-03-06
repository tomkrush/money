package finance

import (
	"encoding/json"
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

func TestCurrency_Subtract(t *testing.T) {
	firstAmount := NewCurrency(10)
	secondAmount := NewCurrency(10)
	expectedAmount := NewCurrency(0)
	actualAmount := firstAmount.Subtract(secondAmount)

	if actualAmount != expectedAmount {
		t.Errorf("Currency subtract failed %d, %d", actualAmount.Amount, expectedAmount.Amount)
	}
}

func TestCurrency_Add(t *testing.T) {
	firstAmount := NewCurrency(10)
	secondAmount := NewCurrency(10)
	expectedAmount := NewCurrency(20)
	actualAmount := firstAmount.Add(secondAmount)

	if actualAmount != expectedAmount {
		t.Errorf("Currency add failed %d, %d", actualAmount.Amount, expectedAmount.Amount)
	}
}

func TestCurrency_UnmarshalJSON(t *testing.T) {
	var c Currency

	if err := json.Unmarshal([]byte("15000"), &c); err != nil {
		t.Errorf("Error unmarshalling %s", err)
	}

	if c.Amount != 15000 {
		t.Errorf("Currency value incorrect %d", c.Amount)
	}
}

func TestCurrency_UnmarshalJSONError(t *testing.T) {
	var c Currency

	if err := json.Unmarshal([]byte("test"), &c); err == nil {
		t.Errorf("Should have errored unmarshalling %s", err)
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

func TestNew(t *testing.T) {
	tests := []struct {
		name   string
		amount int
		want   int
	}{
		{"Get 1000", 1000, 1000},
		{"Get 500", 500, 500},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCurrency(tt.amount); got.Amount != tt.want {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewFromDollars(t *testing.T) {
	tests := []struct {
		name   string
		amount string
		want   int
	}{
		{"Get 10.00", "10.00", 1000},
		{"Get 45.99", "45.99", 4599},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCurrencyFromDollars(tt.amount); got.Amount != tt.want {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

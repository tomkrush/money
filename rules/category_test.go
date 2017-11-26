package rules

import (
	"github.com/tomkrush/money/finance"
	"reflect"
	"testing"
)

func TestCategory_GetCategory(t *testing.T) {
	tests := []struct {
		name         string
		transactions finance.Transactions
		want         CategoryList
	}{
		{
			"Add up categories",
			finance.Transactions{
				finance.Transaction{
					Category: "Gas",
					Amount:   finance.NewCurrency(5000),
				},
				finance.Transaction{
					Category: "Gas",
					Amount:   finance.NewCurrency(2500),
				},
				finance.Transaction{
					Category: "Gas",
					Amount:   finance.NewCurrency(2500),
				},
				finance.Transaction{
					Category: "Shopping",
					Amount:   finance.NewCurrency(5000),
				},
			},
			CategoryList{
				Category{
					Category: "Gas",
					Amount:   finance.NewCurrency(10000),
				},
				Category{
					Category: "Shopping",
					Amount:   finance.NewCurrency(5000),
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			categories := GetCategories(tt.transactions)

			if !reflect.DeepEqual(categories, tt.want) {
				t.Errorf("Categories returned %v, want %v", categories, tt.want)
			}
		})
	}
}

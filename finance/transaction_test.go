package finance

import (
	"testing"
)

func TestTransaction_GetDescription(t *testing.T) {
	tests := []struct {
		name        string
		transaction Transaction
		want        string
	}{
		{"Should return description", Transaction{Description: "Test Description"}, "Test Description"},
		{"Should return description", Transaction{Description: "Test Description", UserDescription: "This is modified"}, "This is modified"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.want != tt.transaction.GetDescription() {
				t.Errorf("GetDescription returned %v, want %v", tt.transaction.GetDescription(), tt.want)
			}
		})
	}
}

func TestTransaction_GetCategory(t *testing.T) {
	tests := []struct {
		name        string
		transaction Transaction
		want        string
	}{
		{"Should return category", Transaction{Category: "Fast Food"}, "Fast Food"},
		{"Should return uncategorized", Transaction{}, "Uncategorized"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.want != tt.transaction.GetCategory() {
				t.Errorf("GetCategory returned %v, want %v", tt.transaction.GetCategory(), tt.want)
			}
		})
	}
}

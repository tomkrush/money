package rules

import (
	"testing"
)

func TestPersonalRules_New(t *testing.T) {
	p := New("testdata/personal-rules.json")

	if p.MonthlyIncome != 3000 {
		t.Error("Monthly Income incorrect")
	}

	if len(p.Categories) == 0 {
		t.Error("No categories imported")
	}

	if len(p.Transactions) == 0 {
		t.Error("No transactions imported")
	}
}

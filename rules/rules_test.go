package rules

import (
	"testing"
	"time"

	"github.com/tomkrush/money/finance"
)

func TestPersonalRules_New(t *testing.T) {
	p := New("testdata/personal-rules.json")

	if len(p.Transactions) == 0 {
		t.Error("No transactions imported")
	}
}

func TestRules_TransactionRule_Apply_Category(t *testing.T) {
	transaction := finance.Transaction{Description: "Target"}

	transactionRule := TransactionRule{
		Contains: "Target",
		Category: "Store",
	}

	transaction, matched := transactionRule.Apply(transaction)

	if transaction.Category != "Store" && matched {
		t.Error("Category failed to assign")
	}
}

func TestRules_TransactionRule_Apply_Replace(t *testing.T) {
	transaction := finance.Transaction{Description: "CHIPOTLE 1408 ELK GROVE VI IL"}

	transactionRule := TransactionRule{
		Contains: "Chipotle",
		Replace:  "Chipotle",
	}

	transaction, matched := transactionRule.Apply(transaction)

	if transaction.GetDescription() != "Chipotle" && matched {
		t.Errorf("New description failed to replace")
	}
}

func TestRules_TransactionRule_Apply_Need(t *testing.T) {
	transaction := finance.Transaction{Description: "BP Gas Station"}

	transactionRule := TransactionRule{
		Contains: "BP Gas Station",
		Need:     true,
	}

	transaction, matched := transactionRule.Apply(transaction)

	if transaction.Need != true && matched {
		t.Error("Need wasn't assigned true")
	}
}

func TestRules_TransactionRule_Apply_FindReplace(t *testing.T) {
	transaction := finance.Transaction{Description: "BP Gas Station"}

	transactionRule := TransactionRule{
		FindReplace: "BP",
	}

	transaction, matched := transactionRule.Apply(transaction)

	if transaction.GetDescription() != "BP" && matched {
		t.Error("FindReplace didnt't find or assign")
	}
}

func TestRules_Rules_Apply(t *testing.T) {
	rules := Rules{
		Transactions: []TransactionRule{
			TransactionRule{
				FindReplace: "BP",
			},
		},
	}

	transactions := rules.Apply(finance.Transactions{
		finance.Transaction{Description: "BP Gas Station"},
	})

	transaction := transactions[0]
	actual := transaction.GetDescription()
	expects := "BP"

	if transaction.GetDescription() != expects {
		t.Errorf("Rules failed to apply description to transaction %s %s", expects, actual)
	}
}

func TestRules_TransactionRule_Apply_Remove(t *testing.T) {
	transaction := finance.Transaction{Description: "BP Gas Station"}

	transactionRule := TransactionRule{
		Remove: "Gas Station",
	}

	transaction, matched := transactionRule.Apply(transaction)

	if transaction.GetDescription() != "BP" && matched {
		t.Errorf("Remove didn't work %s", transaction.GetDescription())
	}
}

func TestRules_TransactionRule_Apply_Some(t *testing.T) {
	transaction := finance.Transaction{Description: "Check #1311"}

	transactionRule := TransactionRule{
		Some:    []string{"check", "cashed"},
		Replace: "Townhouse Rent",
	}

	transaction, matched := transactionRule.Apply(transaction)

	if transaction.GetDescription() != "Townhouse Rent" && matched {
		t.Errorf("Some didn't work %s", transaction.GetDescription())
	}
}

func TestRules_TransactionRule_Apply_SomeAmount(t *testing.T) {
	transaction := finance.Transaction{
		Description: "Check #1311",
		Amount:      finance.NewCurrency(130),
	}

	transactionRule := TransactionRule{
		Some:    []string{"check", "cashed"},
		Replace: "Townhouse Rent",
		Amount:  finance.NewCurrency(150000),
	}

	transaction, matched := transactionRule.Apply(transaction)

	if transaction.GetDescription() == "Townhouse Rent" && matched {
		t.Errorf("Description should not have applied")
	}
}

func TestRules_TransactionRule_Income(t *testing.T) {
	transaction := finance.Transaction{
		Description: "My Work",
		Amount:      finance.NewCurrency(150000),
	}

	transactionRule := TransactionRule{
		Contains: "Work",
		Income: Paycheck{
			Description: "Work Check",
			Frequency:   2,
			Amount:      finance.NewCurrency(150000),
		},
	}

	transaction, _ = transactionRule.Apply(transaction)

	if transaction.GetDescription() != "Work Check" {
		t.Errorf("Transaction description incorrect %s", transaction.GetDescription())
	}

	if transaction.Income != true {
		t.Errorf("Income wasn't identified")
	}
}

func TestRules_TransactionRule_Apply_Bill(t *testing.T) {
	transaction := finance.Transaction{
		Description: "Check #1311",
		Amount:      finance.NewCurrency(150000),
	}

	transactionRule := TransactionRule{
		Some:   []string{"check", "cashed"},
		Amount: finance.NewCurrency(150000),
		Bill: BillRule{
			Description: "Townhouse Rent",
			Day:         1,
			Amount:      finance.NewCurrency(150000),
		},
	}

	transaction, _ = transactionRule.Apply(transaction)

	if transaction.GetDescription() != "Townhouse Rent" {
		t.Errorf("Transaction description Incorrect")
	}

	if transaction.Bill != true {
		t.Errorf("Bill wasn't identified")
	}
}

func TestRules_TransactionRule_Bills(t *testing.T) {
	rules := Rules{
		Transactions: []TransactionRule{
			TransactionRule{
				Some:   []string{"check", "cashed"},
				Amount: finance.NewCurrency(150000),
				Bill: BillRule{
					Description: "Townhouse Rent",
					Day:         1,
					Amount:      finance.NewCurrency(150000),
				},
			},
			TransactionRule{
				Some:   []string{"check", "cashed"},
				Amount: finance.NewCurrency(150000),
				Bill: BillRule{
					Description: "Townhouse Rent",
					Day:         1,
					Amount:      finance.NewCurrency(150000),
				},
			},
			TransactionRule{},
			TransactionRule{},
		},
	}

	transactions := finance.Transactions{}

	bills := rules.Bills(transactions)

	if len(bills.Rules) != 2 {
		t.Error("Incorrect number of bills returned")
	}
}

func TestRules_Rules_PlannedTransactions(t *testing.T) {

	plannedExpenses := PlannedExpenses{
		PlannedExpense{
			Description: "Some Purchase",
			Amount:      finance.NewCurrency(3000),
		},
		createPlannedExpense("Some Purchase", 2000, "2017-11-10"),
	}

	actual := plannedExpenses.TotalExpenses()

	if actual.Amount != 3000 {
		t.Errorf("Planned Expenses should is %d, wants, %d", actual.Amount, 2000)
	}
}

func createPlannedExpense(description string, amount int, date string) PlannedExpense {
	paid, _ := time.Parse("2006-01-02", date)

	return PlannedExpense{
		Description: description,
		Amount:      finance.NewCurrency(amount),
		Paid:        paid,
	}
}

func TestRules_Rules_PlannedTransactions_InMonth(t *testing.T) {
	tests := []struct {
		name            string
		plannedExpenses PlannedExpenses
		want            int
	}{
		{
			"Add up categories",
			PlannedExpenses{
				createPlannedExpense("Some Purchase #1", 2500, "2017-10-01"),
				createPlannedExpense("Some Purchase #2", 2500, "2017-11-10"),
				PlannedExpense{Description: "Some Purchase #3", Amount: finance.NewCurrency(500)},
			},
			500,
		},
		{
			"Add up categories",
			PlannedExpenses{
				createPlannedExpense("Some Purchase #1", 2500, "2017-10-01"),
				PlannedExpense{Description: "Some Purchase #2", Amount: finance.NewCurrency(1500)},
				PlannedExpense{Description: "Some Purchase #3", Amount: finance.NewCurrency(500)},
			},
			2000,
		},
	}

	monthTimestamp, _ := time.Parse("2006-01-02", "2017-11-01")

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			plannedInMonth := tt.plannedExpenses.InMonth(monthTimestamp)
			total := plannedInMonth.TotalExpenses()

			if total.Amount != tt.want {
				t.Errorf("Total expenses returned %v, want %v", total.Amount, tt.want)
			}
		})
	}
}

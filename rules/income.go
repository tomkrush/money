package rules

import (
	"github.com/tomkrush/money/finance"
)

// Paycheck is amount of money that a family adds to their bank.
type Paycheck struct {
	Amount      finance.Currency `json:"amount"`
	Frequency   int              `json:"frequency"`
	Description string           `json:"description"`
}

// IncomeSummary is a summary of income calculations.
//
// - Projected mmount is the amount of income expected by the end of
//   the current month.
// - Actual amount is the amount of income already received for the current month.
// - Remaining amount is the amount of income that is still expected
// - Unplanned amount is the amount of tncome that has been recieved that wasn't
//   expected this month.
type IncomeSummary struct {
	ProjectedAmount finance.Currency
	ActualAmount    finance.Currency
	RemainingAmount finance.Currency
	UnPlannedAmount finance.Currency
}

func getIncomeRuleWithDescription(incomeRules []TransactionRule, description string) (TransactionRule, bool) {
	for _, rule := range incomeRules {
		if rule.Income.Description == description {
			return rule, true
		}
	}

	return TransactionRule{}, false
}

// TotalIncome calculates the total income for a given list of transactions and set
// of rules.
func (r Rules) TotalIncome(transactions finance.Transactions) IncomeSummary {
	var summary IncomeSummary
	var incomeRules []TransactionRule

	// Collect all rules that contain a income frequency
	for _, rule := range r.Transactions {
		if rule.Income.Frequency != 0 {
			incomeRules = append(incomeRules, rule)
		}
	}

	// Iterate over all transactions and check if income rule exists.
	// If rule does exist, than the income is planned. Add the amount to the
	// projected amount and add the amount to the actual amount recieved.
	// If rule doesn't exist, than the income is unplanned. Add the amount to the
	// projected total and also add the amount of the unplanned amount.
	for _, transaction := range transactions {
		_, ok := getIncomeRuleWithDescription(incomeRules, transaction.GetDescription())

		if ok {
			summary.ProjectedAmount = summary.ProjectedAmount.Add(transaction.Amount)
			summary.ActualAmount = summary.ActualAmount.Add(transaction.Amount)
		} else if transaction.Amount.Amount > 0 {
			summary.ProjectedAmount = summary.ProjectedAmount.Add(transaction.Amount)
			summary.UnPlannedAmount = summary.UnPlannedAmount.Add(transaction.Amount)
			summary.ActualAmount = summary.ActualAmount.Add(transaction.Amount)
		}
	}

	// Iterate over all rules and identify the amount of planned income transactions
	// that haven't taken place yet. Add the this difference to the remaining and projected
	// amount fields.
	for _, rule := range incomeRules {
		transactions := transactions.GetAllByDescription(rule.Income.Description)
		remainingCount := (rule.Income.Frequency - len(transactions))
		remainingAmount := finance.NewCurrency(rule.Income.Amount.Amount * remainingCount)

		summary.RemainingAmount = summary.RemainingAmount.Add(remainingAmount)
		summary.ProjectedAmount = summary.ProjectedAmount.Add(remainingAmount)
	}

	return summary
}

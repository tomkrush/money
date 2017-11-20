package rules

import (
	"money/finance"
)

// Paycheck is amount of money that a family adds to their bank.
type Paycheck struct {
	Amount      finance.Currency `json:"amount"`
	Frequency   int              `json:"frequency"`
	Description string           `json:"description"`
}
}

package rules

import "money/finance"

// Income is amount of money that a family adds to their bank.
type Income struct {
	Amount finance.Currency `json:"amount"`
}

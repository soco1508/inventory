package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type Expense struct {
	ExpenseID string          `db:"expenseId" json:"expenseId"`
	Category  string          `db:"category" json:"category"`
	Amount    decimal.Decimal `db:"amount" json:"amount"`
	Timestamp time.Time       `db:"timestamp" json:"timestamp"`
}

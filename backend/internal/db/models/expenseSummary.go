package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type ExpenseSummary struct {
	ExpenseSummaryId string          `db:"expenseSummaryId" json:"expenseSummaryId"`
	TotalExpenses    decimal.Decimal `db:"totalExpenses" json:"totalExpenses"`
	Date             time.Time       `db:"date" json:"date"`
}

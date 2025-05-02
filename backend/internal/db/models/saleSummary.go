package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type SaleSummary struct {
	SalesSummaryID   string          `db:"salesSummaryId" json:"salesSummaryId"`
	TotalValue       decimal.Decimal `db:"totalValue" json:"totalValue"`
	ChangePercentage decimal.Decimal `db:"changePercentage" json:"changePercentage"`
	Date             time.Time       `db:"date" json:"date"`
}

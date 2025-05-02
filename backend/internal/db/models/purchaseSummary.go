package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type PurchaseSummary struct {
	PurchaseSummaryID string          `db:"purchaseSummaryId" json:"purchaseSummaryId"`
	TotalPurchased    decimal.Decimal `db:"totalPurchased" json:"totalPurchased"`
	ChangePercentage  decimal.Decimal `db:"changePercentage" json:"changePercentage"`
	Date              time.Time       `db:"date" json:"date"`
}

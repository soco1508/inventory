package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type Purchase struct {
	PurchaseID string          `db:"purchaseId" json:"purchaseId"`
	ProductID  string          `db:"productId" json:"productId"`
	Timestamp  time.Time       `db:"timestamp" json:"timestamp"`
	Quantity   int32           `db:"quantity" json:"quantity"`
	UnitCost   decimal.Decimal `db:"unitCost" json:"unitCost"`
	TotalCost  decimal.Decimal `db:"totalCost" json:"totalCost"`
}

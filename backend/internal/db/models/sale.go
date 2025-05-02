package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type Sale struct {
	SaleID      string          `db:"saleId" json:"saleId"`
	ProductID   string          `db:"productId" json:"productId"`
	Timestamp   time.Time       `db:"timestamp" json:"timestamp"`
	Quantity    int32           `db:"quantity" json:"quantity"`
	UnitPrice   decimal.Decimal `db:"unitPrice" json:"unitPrice"`
	TotalAmount decimal.Decimal `db:"totalAmount" json:"totalAmount"`
}

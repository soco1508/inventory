package repository

import (
	"backend/internal/db/models"
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PurchaseSummaryRepository interface {
	GetPurchaseSummary(ctx context.Context) ([]*models.PurchaseSummary, error)
}

type purchaseSummaryRepository struct {
	db *sqlx.DB
}

func NewPurchaseSummaryRepository(db *sqlx.DB) PurchaseSummaryRepository {
	return &purchaseSummaryRepository{db: db}
}

func (p *purchaseSummaryRepository) GetPurchaseSummary(ctx context.Context) ([]*models.PurchaseSummary, error) {
	sql := `SELECT * FROM "PurchaseSummary" ORDER BY "date" DESC LIMIT 5`
	rows, err := p.db.QueryxContext(ctx, sql)
	if err != nil {
		return nil, fmt.Errorf("query PurchaseSummary err: %v", err)
	}

	purchaseSummary := []*models.PurchaseSummary{}
	for rows.Next() {
		item := models.PurchaseSummary{}
		if err = rows.StructScan(&item); err != nil {
			return nil, fmt.Errorf("StructScan PurchaseSummary err: %v", err)
		}
		purchaseSummary = append(purchaseSummary, &item)
	}

	return purchaseSummary, nil
}

package repository

import (
	"backend/internal/db/models"
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type ProductRepository interface {
	GetPopularProducts(ctx context.Context) ([]*models.Product, error)
}

type productRepository struct {
	db *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) ProductRepository {
	return &productRepository{db: db}
}

func (p *productRepository) GetPopularProducts(ctx context.Context) ([]*models.Product, error) {
	sql := `SELECT * FROM products ORDER BY stock_quantity DESC LIMIT 15`
	rows, err := p.db.QueryxContext(ctx, sql)
	if err != nil {
		return nil, fmt.Errorf("query products err:\n %+v", err)
	}
	defer rows.Close()

	products := []*models.Product{}
	for rows.Next() {
		product := models.Product{}
		if err = rows.StructScan(&product); err != nil {
			return nil, fmt.Errorf("StructScan product err:\n %+v", err)
		}
		products = append(products, &product)
	}

	return products, nil
}

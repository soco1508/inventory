package repository

import (
	"backend/internal/db/models"
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type ExpenseByCategoryRepository interface {
	GetExpenseByCategory(ctx context.Context) ([]*models.ExpenseByCategory, error)
}

type expenseByCategoryRepository struct {
	db *sqlx.DB
}

func NewExpenseByCategoryRepository(db *sqlx.DB) ExpenseByCategoryRepository {
	return &expenseByCategoryRepository{db: db}
}

func (r *expenseByCategoryRepository) GetExpenseByCategory(ctx context.Context) ([]*models.ExpenseByCategory, error) {
	sql := `SELECT * FROM expense_by_category ORDER BY date DESC LIMIT 5`
	rows, err := r.db.QueryxContext(ctx, sql)
	if err != nil {
		return nil, fmt.Errorf("query ExpenseByCategory err: %v", err)
	}
	defer rows.Close()

	expenseByCategory := []*models.ExpenseByCategory{}
	for rows.Next() {
		item := models.ExpenseByCategory{}
		if err = rows.StructScan(&item); err != nil {
			return nil, fmt.Errorf("StructScan ExpenseByCategory err: %v", err)
		}
		expenseByCategory = append(expenseByCategory, &item)
	}

	return expenseByCategory, nil
}

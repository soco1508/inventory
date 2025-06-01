package repository

import (
	"backend/internal/db/models"
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type UserRepository interface {
	SaveUser(ctx context.Context, user models.User) error
	GetUsers(ctx context.Context) ([]*models.User, error)
}

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepository{db: db}
}

func (u *userRepository) SaveUser(ctx context.Context, user models.User) error {
	sql := `INSERT INTO users(user_id, name, email) VALUES (:user_id, :name, :email)`
	_, err := u.db.NamedExecContext(ctx, sql, &user)
	return err
}

func (u *userRepository) GetUsers(ctx context.Context) ([]*models.User, error) {
	sql := `SELECT user_id, name, email FROM users`
	users := []*models.User{}
	if err := u.db.SelectContext(ctx, &users, sql); err != nil {
		return nil, fmt.Errorf("query users error: %v", err)
	}

	return users, nil
}

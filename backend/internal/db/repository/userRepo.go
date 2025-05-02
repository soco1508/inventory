package repository

import (
	"backend/internal/db/models"
	"context"

	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	SaveUser(ctx context.Context, user models.User) error
}

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepository{db: db}
}

func (u *userRepository) SaveUser(ctx context.Context, user models.User) error {
	sql := `INSERT INTO user(userId, name, email) VALUES (:userId, :name, :email)`
	_, err := u.db.NamedExecContext(ctx, sql, &user)
	return err
}

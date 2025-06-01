package service

import (
	"backend/internal/db/models"
	"backend/internal/db/repository"
	"context"
)

type UserService interface {
	GetUsers(ctx context.Context) ([]*models.User, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) GetUsers(ctx context.Context) ([]*models.User, error) {
	return s.userRepo.GetUsers(ctx)
}

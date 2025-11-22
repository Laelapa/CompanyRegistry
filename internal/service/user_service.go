package service

import (
	"context"

	"github.com/Laelapa/CompanyRegistry/internal/domain"
)

type UserRepository interface {
	Create(ctx context.Context, u *domain.User) (*domain.User, error)
	GetByUsername(ctx context.Context, username string) (*domain.User, error)
}

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

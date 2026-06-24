package repository

import (
	"context"

	"github.com/jesee-kuya/transport-system/domain"
)

type AuthRepository interface {
	CreateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
	GetUserByUsername(ctx context.Context, username string) (*domain.User, error)
}

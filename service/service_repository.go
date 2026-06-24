package service

import (
	"context"

	"github.com/jesee-kuya/transport-system/domain"
)

type Authentication interface {
	SignUp(ctx context.Context, req *domain.SignUpRequest) (*domain.SignUpResponse, error)
}

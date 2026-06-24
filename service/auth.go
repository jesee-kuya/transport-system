package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/jesee-kuya/transport-system/domain"
	"golang.org/x/crypto/bcrypt"
)

func (s *AuthService) SignUp(ctx context.Context, req *domain.SignUpRequest) (*domain.SignUpResponse, error) {
	_, err := s.AuthRepo.GetUserByEmail(ctx, req.Email)
	if err == nil {
		return nil, domain.ErrEmailInUse
	}
	if !errors.Is(err, domain.ErrNotFound) {
		return nil, fmt.Errorf("check email: %w", err)
	}

	_, err = s.AuthRepo.GetUserByUsername(ctx, req.Username)
	if err == nil {
		return nil, domain.ErrUsernameInUse
	}
	if !errors.Is(err, domain.ErrNotFound) {
		return nil, fmt.Errorf("check username: %w", err)
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("hash password: %w", err)
	}

	created, err := s.AuthRepo.CreateUser(ctx, &domain.User{
		Username:     req.Username,
		Email:        req.Email,
		PasswordHash: string(hash),
		Role:         req.Role,
	})
	if err != nil {
		return nil, fmt.Errorf("create user: %w", err)
	}

	return &domain.SignUpResponse{
		ID:       created.ID,
		Username: created.Username,
		Email:    created.Email,
		Role:     created.Role,
	}, nil
}

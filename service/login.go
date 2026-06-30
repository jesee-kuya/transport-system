package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/jesee-kuya/transport-system/domain"
	"github.com/jesee-kuya/transport-system/utils"
	"golang.org/x/crypto/bcrypt"
)

func (s *AuthService) Login(ctx context.Context, req *domain.LoginRequest) (*domain.LoginResponse, error) {
	user, err := s.AuthRepo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return nil, domain.ErrInvalidCredentials
		}
		return nil, fmt.Errorf("get user: %w", err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		return nil, domain.ErrInvalidCredentials
	}

	token, err := utils.GenerateToken(user, s.JWTConfig)
	if err != nil {
		return nil, fmt.Errorf("generate token: %w", err)
	}

	return &domain.LoginResponse{
		Token:    token,
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Role:     user.Role,
	}, nil
}

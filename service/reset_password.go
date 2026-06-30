package service

import (
	"context"
	"fmt"

	"github.com/jesee-kuya/transport-system/domain"
	"github.com/jesee-kuya/transport-system/utils"
	"golang.org/x/crypto/bcrypt"
)

func (s *AuthService) ResetPassword(ctx context.Context, req *domain.ResetPasswordRequest) error {
	claims, err := utils.ParseResetToken(req.Token, s.JWTConfig)
	if err != nil {
		return domain.ErrInvalidToken
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("hash password: %w", err)
	}

	return s.AuthRepo.UpdatePassword(ctx, claims.UserID, string(hash))
}

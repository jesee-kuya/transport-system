package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/jesee-kuya/transport-system/domain"
	"github.com/jesee-kuya/transport-system/utils"
)

func (s *AuthService) ForgotPassword(ctx context.Context, req *domain.ForgotPasswordRequest) (*domain.ForgotPasswordResponse, error) {
	user, err := s.AuthRepo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return &domain.ForgotPasswordResponse{
				Message: "if that email is registered, a reset link has been sent",
			}, nil
		}
		return nil, fmt.Errorf("get user: %w", err)
	}

	resetToken, err := utils.GenerateResetToken(user, s.JWTConfig)
	if err != nil {
		return nil, fmt.Errorf("generate reset token: %w", err)
	}

	return &domain.ForgotPasswordResponse{
		ResetToken: resetToken,
		Message:    "if that email is registered, a reset link has been sent",
	}, nil
}

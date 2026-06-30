package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jesee-kuya/transport-system/domain"
)

func GenerateToken(user *domain.User, cfg *domain.JWTConfig) (string, error) {
	claims := domain.Claims{
		UserID:   user.ID,
		Username: user.Username,
		Email:    user.Email,
		Role:     user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(cfg.TokenExpiration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString([]byte(cfg.SecretKey))
	if err != nil {
		return "", fmt.Errorf("sign token: %w", err)
	}
	return signed, nil
}

func GenerateResetToken(user *domain.User, cfg *domain.JWTConfig) (string, error) {
	claims := domain.ResetClaims{
		UserID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString([]byte(cfg.SecretKey))
	if err != nil {
		return "", fmt.Errorf("sign reset token: %w", err)
	}
	return signed, nil
}

func ParseResetToken(tokenStr string, cfg *domain.JWTConfig) (*domain.ResetClaims, error) {
	claims := &domain.ResetClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(cfg.SecretKey), nil
	})
	if err != nil || !token.Valid {
		return nil, domain.ErrInvalidToken
	}
	return claims, nil
}

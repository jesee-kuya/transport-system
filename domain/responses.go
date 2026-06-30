package domain

import "github.com/google/uuid"

type SignUpResponse struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Role     string    `json:"role"`
}

type LoginResponse struct {
	Token    string    `json:"token"`
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Role     string    `json:"role"`
}

type ForgotPasswordResponse struct {
	ResetToken string `json:"reset_token"`
	Message    string `json:"message"`
}

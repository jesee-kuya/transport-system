package middleware

import "github.com/jesee-kuya/transport-system/domain"

func NewMiddleware(cfg *domain.JWTConfig) Middleware {
	return &MiddlewareStruct{JWT: cfg}
}

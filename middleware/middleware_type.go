package middleware

import "github.com/jesee-kuya/transport-system/domain"

const ClaimsKey = "claims"

type MiddlewareStruct struct {
	JWT *domain.JWTConfig
}

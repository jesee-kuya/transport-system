package service

import (
	"github.com/jesee-kuya/transport-system/domain"
	"github.com/jesee-kuya/transport-system/repository"
)

type AuthService struct {
	AuthRepo  repository.AuthRepository
	JWTConfig *domain.JWTConfig
}

type SchoolSvc struct {
	SchoolRepo repository.SchoolRepository
}

type GuardianSvc struct {
	GuardianRepo repository.GuardianRepository
}

type PrivateParentSvc struct {
	ParentRepo repository.PrivateParentRepository
}

type SchoolDriverSvc struct {
	DriverRepo repository.SchoolDriverRepository
}

type PrivateDriverSvc struct {
	DriverRepo repository.PrivateDriverRepository
}

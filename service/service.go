package service

import (
	"github.com/jesee-kuya/transport-system/domain"
	"github.com/jesee-kuya/transport-system/repository"
)

func NewAuthService(repo repository.AuthRepository, cfg *domain.JWTConfig) Authentication {
	return &AuthService{AuthRepo: repo, JWTConfig: cfg}
}

func NewSchoolService(repo repository.SchoolRepository) SchoolService {
	return &SchoolSvc{SchoolRepo: repo}
}

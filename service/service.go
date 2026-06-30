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

func NewGuardianService(repo repository.GuardianRepository) GuardianService {
	return &GuardianSvc{GuardianRepo: repo}
}

func NewPrivateParentService(repo repository.PrivateParentRepository) PrivateParentService {
	return &PrivateParentSvc{ParentRepo: repo}
}

func NewSchoolDriverService(repo repository.SchoolDriverRepository) SchoolDriverService {
	return &SchoolDriverSvc{DriverRepo: repo}
}

func NewPrivateDriverService(repo repository.PrivateDriverRepository) PrivateDriverService {
	return &PrivateDriverSvc{DriverRepo: repo}
}

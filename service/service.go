package service

import "github.com/jesee-kuya/transport-system/repository"

func NewAuthService(repo repository.AuthRepository) Authentication {
	return &AuthService{AuthRepo: repo}
}

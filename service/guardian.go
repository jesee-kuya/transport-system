package service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jesee-kuya/transport-system/domain"
)

func (s *GuardianSvc) getGuardian(ctx context.Context, userID uuid.UUID) (*domain.Guardian, error) {
	guardian, err := s.GuardianRepo.GetGuardianByUserID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("get guardian: %w", err)
	}
	return guardian, nil
}

func (s *GuardianSvc) GetMyStudents(ctx context.Context, userID uuid.UUID) ([]*domain.Student, error) {
	guardian, err := s.getGuardian(ctx, userID)
	if err != nil {
		return nil, err
	}
	return s.GuardianRepo.GetStudentsByGuardianID(ctx, guardian.ID)
}

func (s *GuardianSvc) GetMyStudent(ctx context.Context, userID uuid.UUID, studentID uuid.UUID) (*domain.Student, error) {
	guardian, err := s.getGuardian(ctx, userID)
	if err != nil {
		return nil, err
	}
	return s.GuardianRepo.GetStudentByIDForGuardian(ctx, guardian.ID, studentID)
}

func (s *GuardianSvc) TrackMyStudent(ctx context.Context, userID uuid.UUID, studentID uuid.UUID) (*domain.Trip, error) {
	if _, err := s.GetMyStudent(ctx, userID, studentID); err != nil {
		return nil, err
	}
	return s.GuardianRepo.GetActiveSchoolTripForStudent(ctx, studentID)
}

func (s *GuardianSvc) GetMyProfile(ctx context.Context, userID uuid.UUID) (*domain.Guardian, error) {
	return s.getGuardian(ctx, userID)
}

func (s *GuardianSvc) EditMyProfile(ctx context.Context, userID uuid.UUID, req *domain.EditGuardianProfileRequest) (*domain.Guardian, error) {
	guardian, err := s.getGuardian(ctx, userID)
	if err != nil {
		return nil, err
	}
	guardian.FirstName = req.FirstName
	guardian.LastName = req.LastName
	guardian.Phone = req.Phone
	return s.GuardianRepo.UpdateGuardian(ctx, guardian)
}

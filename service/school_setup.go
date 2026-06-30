package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jesee-kuya/transport-system/domain"
)

func (s *SchoolSvc) AddMySchool(ctx context.Context, adminID uuid.UUID, req *domain.AddSchoolRequest) (*domain.School, error) {
	_, err := s.SchoolRepo.GetSchoolByAdminID(ctx, adminID)
	if err == nil {
		return nil, domain.ErrSchoolAlreadyExists
	}
	if !errors.Is(err, domain.ErrNotFound) {
		return nil, fmt.Errorf("check existing school: %w", err)
	}

	return s.SchoolRepo.CreateSchool(ctx, &domain.School{
		AdminID:      adminID,
		Name:         req.Name,
		Address:      req.Address,
		ContactEmail: req.ContactEmail,
		ContactPhone: req.ContactPhone,
	})
}

package service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jesee-kuya/transport-system/domain"
)

func (s *SchoolSvc) AddBus(ctx context.Context, adminID uuid.UUID, req *domain.AddBusRequest) (*domain.Bus, error) {
	schoolID, err := s.getSchoolID(ctx, adminID)
	if err != nil {
		return nil, err
	}
	return s.SchoolRepo.CreateBus(ctx, &domain.Bus{
		SchoolID:    schoolID,
		PlateNumber: req.PlateNumber,
		Model:       req.Model,
		Capacity:    req.Capacity,
	})
}

func (s *SchoolSvc) GetBuses(ctx context.Context, adminID uuid.UUID) ([]*domain.Bus, error) {
	schoolID, err := s.getSchoolID(ctx, adminID)
	if err != nil {
		return nil, err
	}
	return s.SchoolRepo.GetBusesBySchoolID(ctx, schoolID)
}

func (s *SchoolSvc) GetBus(ctx context.Context, adminID uuid.UUID, busID uuid.UUID) (*domain.Bus, error) {
	schoolID, err := s.getSchoolID(ctx, adminID)
	if err != nil {
		return nil, err
	}
	bus, err := s.SchoolRepo.GetBusByID(ctx, busID)
	if err != nil {
		return nil, err
	}
	if bus.SchoolID != schoolID {
		return nil, domain.ErrNotFound
	}
	return bus, nil
}

func (s *SchoolSvc) EditBus(ctx context.Context, adminID uuid.UUID, busID uuid.UUID, req *domain.EditBusRequest) (*domain.Bus, error) {
	bus, err := s.GetBus(ctx, adminID, busID)
	if err != nil {
		return nil, err
	}
	bus.PlateNumber = req.PlateNumber
	bus.Model = req.Model
	bus.Capacity = req.Capacity
	bus.IsActive = req.IsActive
	return s.SchoolRepo.UpdateBus(ctx, bus)
}

func (s *SchoolSvc) FilterBuses(ctx context.Context, adminID uuid.UUID, isActive *bool) ([]*domain.Bus, error) {
	schoolID, err := s.getSchoolID(ctx, adminID)
	if err != nil {
		return nil, err
	}
	return s.SchoolRepo.FilterBuses(ctx, schoolID, isActive)
}

func (s *SchoolSvc) SearchBuses(ctx context.Context, adminID uuid.UUID, query string) ([]*domain.Bus, error) {
	schoolID, err := s.getSchoolID(ctx, adminID)
	if err != nil {
		return nil, err
	}
	return s.SchoolRepo.SearchBuses(ctx, schoolID, query)
}

func (s *SchoolSvc) RemoveBus(ctx context.Context, adminID uuid.UUID, busID uuid.UUID) error {
	if _, err := s.GetBus(ctx, adminID, busID); err != nil {
		return err
	}
	return s.SchoolRepo.DeactivateBus(ctx, busID)
}

func (s *SchoolSvc) TrackBus(ctx context.Context, adminID uuid.UUID, busID uuid.UUID) (*domain.Trip, error) {
	if _, err := s.GetBus(ctx, adminID, busID); err != nil {
		return nil, err
	}
	trip, err := s.SchoolRepo.GetActiveTripByBusID(ctx, busID)
	if err != nil {
		return nil, fmt.Errorf("get active trip: %w", err)
	}
	return trip, nil
}

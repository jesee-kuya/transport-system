package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jesee-kuya/transport-system/domain"
)

func (s *PrivateDriverSvc) getDriver(ctx context.Context, userID uuid.UUID) (*domain.PrivateDriver, error) {
	driver, err := s.DriverRepo.GetPrivateDriverByUserID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("get private driver: %w", err)
	}
	return driver, nil
}

func (s *PrivateDriverSvc) getDriverTrip(ctx context.Context, driver *domain.PrivateDriver, tripID uuid.UUID) (*domain.PrivateTrip, error) {
	trip, err := s.DriverRepo.GetPrivateTripByID(ctx, tripID)
	if err != nil {
		return nil, err
	}
	if trip.DriverID != driver.ID {
		return nil, domain.ErrNotFound
	}
	return trip, nil
}

func (s *PrivateDriverSvc) KYCDriver(ctx context.Context, userID uuid.UUID, req *domain.KYCDriverRequest) (*domain.PrivateDriver, error) {
	_, err := s.DriverRepo.GetPrivateDriverByUserID(ctx, userID)
	if err == nil {
		return nil, domain.ErrKYCAlreadySubmitted
	}
	if !errors.Is(err, domain.ErrNotFound) {
		return nil, fmt.Errorf("check existing driver: %w", err)
	}
	return s.DriverRepo.CreatePrivateDriver(ctx, &domain.PrivateDriver{
		UserID:        userID,
		FirstName:     req.FirstName,
		LastName:      req.LastName,
		Phone:         req.Phone,
		LicenseNumber: req.LicenseNumber,
	})
}

func (s *PrivateDriverSvc) MatchWithParent(ctx context.Context, userID uuid.UUID, matchID uuid.UUID, req *domain.RespondToMatchRequest) (*domain.DriverParentMatch, error) {
	driver, err := s.getDriver(ctx, userID)
	if err != nil {
		return nil, err
	}
	match, err := s.DriverRepo.GetMatchByID(ctx, matchID)
	if err != nil {
		return nil, err
	}
	if match.DriverID != driver.ID {
		return nil, domain.ErrNotFound
	}
	if match.Status != "pending" {
		return nil, domain.ErrMatchNotPending
	}
	return s.DriverRepo.UpdateMatchStatus(ctx, matchID, req.Status)
}

func (s *PrivateDriverSvc) StartTrip(ctx context.Context, userID uuid.UUID, req *domain.StartPrivateTripRequest) (*domain.PrivateTrip, error) {
	driver, err := s.getDriver(ctx, userID)
	if err != nil {
		return nil, err
	}
	_, err = s.DriverRepo.GetActivePrivateTripByDriverID(ctx, driver.ID)
	if err == nil {
		return nil, domain.ErrTripAlreadyActive
	}
	if !errors.Is(err, domain.ErrNotFound) {
		return nil, fmt.Errorf("check active trip: %w", err)
	}
	matchID, err := uuid.Parse(req.MatchID)
	if err != nil {
		return nil, fmt.Errorf("parse match id: %w", err)
	}
	match, err := s.DriverRepo.GetMatchByID(ctx, matchID)
	if err != nil {
		return nil, err
	}
	if match.DriverID != driver.ID || match.Status != "accepted" {
		return nil, domain.ErrNotFound
	}
	now := time.Now()
	return s.DriverRepo.CreatePrivateTrip(ctx, &domain.PrivateTrip{
		DriverID:  driver.ID,
		MatchID:   matchID,
		Status:    "in_progress",
		StartedAt: &now,
	})
}

func (s *PrivateDriverSvc) EndTrip(ctx context.Context, userID uuid.UUID, tripID uuid.UUID) (*domain.PrivateTrip, error) {
	driver, err := s.getDriver(ctx, userID)
	if err != nil {
		return nil, err
	}
	trip, err := s.getDriverTrip(ctx, driver, tripID)
	if err != nil {
		return nil, err
	}
	now := time.Now()
	trip.Status = "completed"
	trip.EndedAt = &now
	return s.DriverRepo.UpdatePrivateTrip(ctx, trip)
}

func (s *PrivateDriverSvc) OnboardPrivateStudent(ctx context.Context, userID uuid.UUID, tripID uuid.UUID, req *domain.OnboardPrivateStudentRequest) (*domain.PrivateTripChild, error) {
	driver, err := s.getDriver(ctx, userID)
	if err != nil {
		return nil, err
	}
	if _, err := s.getDriverTrip(ctx, driver, tripID); err != nil {
		return nil, err
	}
	childID, err := uuid.Parse(req.ChildID)
	if err != nil {
		return nil, fmt.Errorf("parse child id: %w", err)
	}
	_, err = s.DriverRepo.GetPrivateTripChildByTripAndChildID(ctx, tripID, childID)
	if err == nil {
		return nil, domain.ErrStudentAlreadyBoarded
	}
	if !errors.Is(err, domain.ErrNotFound) {
		return nil, fmt.Errorf("check child boarding: %w", err)
	}
	return s.DriverRepo.AddPrivateTripChild(ctx, &domain.PrivateTripChild{
		TripID:  tripID,
		ChildID: childID,
	})
}

func (s *PrivateDriverSvc) UpdatePrivateTripStatus(ctx context.Context, userID uuid.UUID, tripID uuid.UUID, req *domain.UpdatePrivateTripStatusRequest) (*domain.PrivateTrip, error) {
	driver, err := s.getDriver(ctx, userID)
	if err != nil {
		return nil, err
	}
	trip, err := s.getDriverTrip(ctx, driver, tripID)
	if err != nil {
		return nil, err
	}
	trip.Status = req.Status
	return s.DriverRepo.UpdatePrivateTrip(ctx, trip)
}

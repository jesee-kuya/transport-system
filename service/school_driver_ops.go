package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jesee-kuya/transport-system/domain"
)

func (s *SchoolDriverSvc) getDriver(ctx context.Context, userID uuid.UUID) (*domain.SchoolDriver, error) {
	driver, err := s.DriverRepo.GetSchoolDriverByUserID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("get school driver: %w", err)
	}
	return driver, nil
}

func (s *SchoolDriverSvc) getDriverTrip(ctx context.Context, driver *domain.SchoolDriver, tripID uuid.UUID) (*domain.Trip, error) {
	trip, err := s.DriverRepo.GetTripByID(ctx, tripID)
	if err != nil {
		return nil, err
	}
	if trip.DriverID != driver.ID {
		return nil, domain.ErrNotFound
	}
	return trip, nil
}

func (s *SchoolDriverSvc) SearchStudents(ctx context.Context, userID uuid.UUID, query string) ([]*domain.Student, error) {
	driver, err := s.getDriver(ctx, userID)
	if err != nil {
		return nil, err
	}
	return s.DriverRepo.SearchStudentsBySchoolID(ctx, driver.SchoolID, query)
}

func (s *SchoolDriverSvc) StartTrip(ctx context.Context, userID uuid.UUID, req *domain.StartTripRequest) (*domain.Trip, error) {
	driver, err := s.getDriver(ctx, userID)
	if err != nil {
		return nil, err
	}
	_, err = s.DriverRepo.GetActiveTripByDriverID(ctx, driver.ID)
	if err == nil {
		return nil, domain.ErrTripAlreadyActive
	}
	if !errors.Is(err, domain.ErrNotFound) {
		return nil, fmt.Errorf("check active trip: %w", err)
	}
	busID, err := uuid.Parse(req.BusID)
	if err != nil {
		return nil, fmt.Errorf("parse bus id: %w", err)
	}
	now := time.Now()
	return s.DriverRepo.CreateTrip(ctx, &domain.Trip{
		SchoolID:  driver.SchoolID,
		DriverID:  driver.ID,
		BusID:     busID,
		TripType:  req.TripType,
		Status:    "in_progress",
		StartedAt: &now,
	})
}

func (s *SchoolDriverSvc) EndTrip(ctx context.Context, userID uuid.UUID, tripID uuid.UUID) (*domain.Trip, error) {
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
	return s.DriverRepo.UpdateTrip(ctx, trip)
}

func (s *SchoolDriverSvc) UpdateTripStatus(ctx context.Context, userID uuid.UUID, tripID uuid.UUID, req *domain.UpdateTripStatusRequest) (*domain.Trip, error) {
	driver, err := s.getDriver(ctx, userID)
	if err != nil {
		return nil, err
	}
	trip, err := s.getDriverTrip(ctx, driver, tripID)
	if err != nil {
		return nil, err
	}
	trip.Status = req.Status
	return s.DriverRepo.UpdateTrip(ctx, trip)
}

func (s *SchoolDriverSvc) OnboardStudent(ctx context.Context, userID uuid.UUID, tripID uuid.UUID, studentID uuid.UUID) (*domain.TripStudent, error) {
	driver, err := s.getDriver(ctx, userID)
	if err != nil {
		return nil, err
	}
	if _, err := s.getDriverTrip(ctx, driver, tripID); err != nil {
		return nil, err
	}
	student, err := s.DriverRepo.GetStudentByID(ctx, studentID)
	if err != nil {
		return nil, err
	}
	if student.SchoolID != driver.SchoolID {
		return nil, domain.ErrNotFound
	}
	_, err = s.DriverRepo.GetTripStudentByTripAndStudentID(ctx, tripID, studentID)
	if err == nil {
		return nil, domain.ErrStudentAlreadyBoarded
	}
	if !errors.Is(err, domain.ErrNotFound) {
		return nil, fmt.Errorf("check student boarding: %w", err)
	}
	return s.DriverRepo.AddTripStudent(ctx, &domain.TripStudent{
		TripID:    tripID,
		StudentID: studentID,
	})
}

func (s *SchoolDriverSvc) ViewBoardedStudents(ctx context.Context, userID uuid.UUID, tripID uuid.UUID) ([]*domain.Student, error) {
	driver, err := s.getDriver(ctx, userID)
	if err != nil {
		return nil, err
	}
	if _, err := s.getDriverTrip(ctx, driver, tripID); err != nil {
		return nil, err
	}
	return s.DriverRepo.GetBoardedStudents(ctx, tripID)
}

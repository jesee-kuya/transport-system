package service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jesee-kuya/transport-system/domain"
)

func (s *SchoolSvc) AddDriver(ctx context.Context, adminID uuid.UUID, req *domain.AddDriverRequest) (*domain.SchoolDriver, error) {
	schoolID, err := s.getSchoolID(ctx, adminID)
	if err != nil {
		return nil, err
	}
	userID, err := uuid.Parse(req.UserID)
	if err != nil {
		return nil, fmt.Errorf("parse user id: %w", err)
	}
	return s.SchoolRepo.CreateSchoolDriver(ctx, &domain.SchoolDriver{
		UserID:        userID,
		SchoolID:      schoolID,
		LicenseNumber: req.LicenseNumber,
	})
}

func (s *SchoolSvc) GetDrivers(ctx context.Context, adminID uuid.UUID) ([]*domain.SchoolDriver, error) {
	schoolID, err := s.getSchoolID(ctx, adminID)
	if err != nil {
		return nil, err
	}
	return s.SchoolRepo.GetSchoolDriversBySchoolID(ctx, schoolID)
}

func (s *SchoolSvc) GetDriver(ctx context.Context, adminID uuid.UUID, driverID uuid.UUID) (*domain.SchoolDriver, error) {
	schoolID, err := s.getSchoolID(ctx, adminID)
	if err != nil {
		return nil, err
	}
	driver, err := s.SchoolRepo.GetSchoolDriverByID(ctx, driverID)
	if err != nil {
		return nil, err
	}
	if driver.SchoolID != schoolID {
		return nil, domain.ErrNotFound
	}
	return driver, nil
}

func (s *SchoolSvc) EditDriver(ctx context.Context, adminID uuid.UUID, driverID uuid.UUID, req *domain.EditDriverRequest) (*domain.SchoolDriver, error) {
	driver, err := s.GetDriver(ctx, adminID, driverID)
	if err != nil {
		return nil, err
	}
	driver.LicenseNumber = req.LicenseNumber
	driver.IsActive = req.IsActive
	return s.SchoolRepo.UpdateSchoolDriver(ctx, driver)
}

func (s *SchoolSvc) FilterDrivers(ctx context.Context, adminID uuid.UUID, isActive *bool) ([]*domain.SchoolDriver, error) {
	schoolID, err := s.getSchoolID(ctx, adminID)
	if err != nil {
		return nil, err
	}
	return s.SchoolRepo.FilterSchoolDrivers(ctx, schoolID, isActive)
}

func (s *SchoolSvc) SearchDrivers(ctx context.Context, adminID uuid.UUID, query string) ([]*domain.SchoolDriver, error) {
	schoolID, err := s.getSchoolID(ctx, adminID)
	if err != nil {
		return nil, err
	}
	return s.SchoolRepo.SearchSchoolDrivers(ctx, schoolID, query)
}

func (s *SchoolSvc) RemoveDriver(ctx context.Context, adminID uuid.UUID, driverID uuid.UUID) error {
	if _, err := s.GetDriver(ctx, adminID, driverID); err != nil {
		return err
	}
	return s.SchoolRepo.DeactivateSchoolDriver(ctx, driverID)
}

func (s *SchoolSvc) ConnectPrivateDriver(ctx context.Context, adminID uuid.UUID, req *domain.ConnectPrivateDriverRequest) error {
	schoolID, err := s.getSchoolID(ctx, adminID)
	if err != nil {
		return err
	}
	driverID, err := uuid.Parse(req.DriverID)
	if err != nil {
		return fmt.Errorf("parse driver id: %w", err)
	}
	if _, err := s.SchoolRepo.GetPrivateDriverByID(ctx, driverID); err != nil {
		return err
	}
	if err := s.SchoolRepo.ConnectPrivateDriver(ctx, schoolID, driverID); err != nil {
		return fmt.Errorf("connect private driver: %w", err)
	}
	return nil
}

func (s *SchoolSvc) GetPrivateDrivers(ctx context.Context, adminID uuid.UUID) ([]*domain.PrivateDriver, error) {
	schoolID, err := s.getSchoolID(ctx, adminID)
	if err != nil {
		return nil, err
	}
	return s.SchoolRepo.GetPrivateDriversBySchoolID(ctx, schoolID)
}

func (s *SchoolSvc) GetPrivateDriver(ctx context.Context, adminID uuid.UUID, driverID uuid.UUID) (*domain.PrivateDriver, error) {
	schoolID, err := s.getSchoolID(ctx, adminID)
	if err != nil {
		return nil, err
	}
	drivers, err := s.SchoolRepo.GetPrivateDriversBySchoolID(ctx, schoolID)
	if err != nil {
		return nil, err
	}
	for _, d := range drivers {
		if d.ID == driverID {
			return d, nil
		}
	}
	return nil, domain.ErrNotFound
}

func (s *SchoolSvc) TrackPrivateTrips(ctx context.Context, adminID uuid.UUID) ([]*domain.PrivateTrip, error) {
	schoolID, err := s.getSchoolID(ctx, adminID)
	if err != nil {
		return nil, err
	}
	return s.SchoolRepo.GetActivePrivateTripsBySchoolID(ctx, schoolID)
}

func (s *SchoolSvc) GetPrivateTripData(ctx context.Context, adminID uuid.UUID, tripID uuid.UUID) (*domain.PrivateTrip, error) {
	schoolID, err := s.getSchoolID(ctx, adminID)
	if err != nil {
		return nil, err
	}
	trip, err := s.SchoolRepo.GetPrivateTripByID(ctx, tripID)
	if err != nil {
		return nil, err
	}
	// verify the trip belongs to a driver connected to this school
	drivers, err := s.SchoolRepo.GetPrivateDriversBySchoolID(ctx, schoolID)
	if err != nil {
		return nil, err
	}
	for _, d := range drivers {
		if d.ID == trip.DriverID {
			return trip, nil
		}
	}
	return nil, domain.ErrNotFound
}

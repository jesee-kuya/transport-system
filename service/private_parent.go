package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jesee-kuya/transport-system/domain"
)

func (s *PrivateParentSvc) getParent(ctx context.Context, userID uuid.UUID) (*domain.PrivateParent, error) {
	parent, err := s.ParentRepo.GetPrivateParentByUserID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("get private parent: %w", err)
	}
	return parent, nil
}

func (s *PrivateParentSvc) CollectKYC(ctx context.Context, userID uuid.UUID, req *domain.CollectKYCRequest) (*domain.PrivateParent, error) {
	_, err := s.ParentRepo.GetPrivateParentByUserID(ctx, userID)
	if err == nil {
		return nil, domain.ErrKYCAlreadySubmitted
	}
	if !errors.Is(err, domain.ErrNotFound) {
		return nil, fmt.Errorf("check existing parent: %w", err)
	}
	return s.ParentRepo.CreatePrivateParent(ctx, &domain.PrivateParent{
		UserID:    userID,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Phone:     req.Phone,
	})
}

func (s *PrivateParentSvc) GetPrivateProfile(ctx context.Context, userID uuid.UUID) (*domain.PrivateParent, error) {
	return s.getParent(ctx, userID)
}

func (s *PrivateParentSvc) AddMyChild(ctx context.Context, userID uuid.UUID, req *domain.AddPrivateChildRequest) (*domain.PrivateChild, error) {
	parent, err := s.getParent(ctx, userID)
	if err != nil {
		return nil, err
	}
	return s.ParentRepo.CreatePrivateChild(ctx, &domain.PrivateChild{
		ParentID:  parent.ID,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Grade:     req.Grade,
	})
}

func (s *PrivateParentSvc) GetMyChildren(ctx context.Context, userID uuid.UUID) ([]*domain.PrivateChild, error) {
	parent, err := s.getParent(ctx, userID)
	if err != nil {
		return nil, err
	}
	return s.ParentRepo.GetPrivateChildrenByParentID(ctx, parent.ID)
}

func (s *PrivateParentSvc) EditMyPrivateProfile(ctx context.Context, userID uuid.UUID, req *domain.EditPrivateProfileRequest) (*domain.PrivateParent, error) {
	parent, err := s.getParent(ctx, userID)
	if err != nil {
		return nil, err
	}
	parent.FirstName = req.FirstName
	parent.LastName = req.LastName
	parent.Phone = req.Phone
	return s.ParentRepo.UpdatePrivateParent(ctx, parent)
}

func (s *PrivateParentSvc) EditMyChild(ctx context.Context, userID uuid.UUID, childID uuid.UUID, req *domain.EditPrivateChildRequest) (*domain.PrivateChild, error) {
	parent, err := s.getParent(ctx, userID)
	if err != nil {
		return nil, err
	}
	child, err := s.ParentRepo.GetPrivateChildByID(ctx, childID)
	if err != nil {
		return nil, err
	}
	if child.ParentID != parent.ID {
		return nil, domain.ErrNotFound
	}
	child.FirstName = req.FirstName
	child.LastName = req.LastName
	child.Grade = req.Grade
	return s.ParentRepo.UpdatePrivateChild(ctx, child)
}

func (s *PrivateParentSvc) DeleteMyAccount(ctx context.Context, userID uuid.UUID) error {
	if _, err := s.getParent(ctx, userID); err != nil {
		return err
	}
	return s.ParentRepo.DeleteAccount(ctx, userID)
}

func (s *PrivateParentSvc) MatchWithDriver(ctx context.Context, userID uuid.UUID, req *domain.MatchWithDriverRequest) (*domain.DriverParentMatch, error) {
	parent, err := s.getParent(ctx, userID)
	if err != nil {
		return nil, err
	}
	driverID, err := uuid.Parse(req.DriverID)
	if err != nil {
		return nil, fmt.Errorf("parse driver id: %w", err)
	}
	if _, err := s.ParentRepo.GetPrivateDriverByID(ctx, driverID); err != nil {
		return nil, err
	}
	_, err = s.ParentRepo.GetMatchByParentAndDriverID(ctx, parent.ID, driverID)
	if err == nil {
		return nil, domain.ErrMatchAlreadyExists
	}
	if !errors.Is(err, domain.ErrNotFound) {
		return nil, fmt.Errorf("check existing match: %w", err)
	}
	return s.ParentRepo.CreateDriverParentMatch(ctx, &domain.DriverParentMatch{
		ParentID: parent.ID,
		DriverID: driverID,
		Status:   "pending",
	})
}

func (s *PrivateParentSvc) ConnectWithSchool(ctx context.Context, userID uuid.UUID, req *domain.ConnectWithSchoolRequest) (*domain.ParentSchoolConnection, error) {
	parent, err := s.getParent(ctx, userID)
	if err != nil {
		return nil, err
	}
	schoolID, err := uuid.Parse(req.SchoolID)
	if err != nil {
		return nil, fmt.Errorf("parse school id: %w", err)
	}
	if _, err := s.ParentRepo.GetSchoolByID(ctx, schoolID); err != nil {
		return nil, err
	}
	_, err = s.ParentRepo.GetParentSchoolConnection(ctx, parent.ID, schoolID)
	if err == nil {
		return nil, domain.ErrSchoolAlreadyConnected
	}
	if !errors.Is(err, domain.ErrNotFound) {
		return nil, fmt.Errorf("check existing connection: %w", err)
	}
	return s.ParentRepo.CreateParentSchoolConnection(ctx, &domain.ParentSchoolConnection{
		ParentID: parent.ID,
		SchoolID: schoolID,
	})
}

func (s *PrivateParentSvc) GetSchools(ctx context.Context) ([]*domain.School, error) {
	return s.ParentRepo.GetAllSchools(ctx)
}

func (s *PrivateParentSvc) SearchSchools(ctx context.Context, query string) ([]*domain.School, error) {
	return s.ParentRepo.SearchSchools(ctx, query)
}

func (s *PrivateParentSvc) FilterSchools(ctx context.Context, address string) ([]*domain.School, error) {
	return s.ParentRepo.FilterSchools(ctx, address)
}

func (s *PrivateParentSvc) GetSchool(ctx context.Context, schoolID uuid.UUID) (*domain.School, error) {
	return s.ParentRepo.GetSchoolByID(ctx, schoolID)
}

func (s *PrivateParentSvc) ReceiveStudent(ctx context.Context, userID uuid.UUID, tripID uuid.UUID) (*domain.PrivateTrip, error) {
	parent, err := s.getParent(ctx, userID)
	if err != nil {
		return nil, err
	}
	if _, err := s.ParentRepo.GetPrivateTripByIDForParent(ctx, parent.ID, tripID); err != nil {
		return nil, err
	}
	return s.ParentRepo.UpdatePrivateTripStatus(ctx, tripID, "student_received")
}

func (s *PrivateParentSvc) ConfirmStudentBoarding(ctx context.Context, userID uuid.UUID, tripID uuid.UUID) (*domain.PrivateTrip, error) {
	parent, err := s.getParent(ctx, userID)
	if err != nil {
		return nil, err
	}
	if _, err := s.ParentRepo.GetPrivateTripByIDForParent(ctx, parent.ID, tripID); err != nil {
		return nil, err
	}
	return s.ParentRepo.UpdatePrivateTripStatus(ctx, tripID, "boarding_confirmed")
}

func (s *PrivateParentSvc) GetTrips(ctx context.Context, userID uuid.UUID) ([]*domain.PrivateTrip, error) {
	parent, err := s.getParent(ctx, userID)
	if err != nil {
		return nil, err
	}
	return s.ParentRepo.GetPrivateTripsByParentID(ctx, parent.ID)
}

func (s *PrivateParentSvc) TrackTrip(ctx context.Context, userID uuid.UUID, tripID uuid.UUID) (*domain.PrivateTrip, error) {
	parent, err := s.getParent(ctx, userID)
	if err != nil {
		return nil, err
	}
	return s.ParentRepo.GetPrivateTripByIDForParent(ctx, parent.ID, tripID)
}

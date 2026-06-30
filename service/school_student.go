package service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jesee-kuya/transport-system/domain"
)

func (s *SchoolSvc) getSchoolID(ctx context.Context, adminID uuid.UUID) (uuid.UUID, error) {
	school, err := s.SchoolRepo.GetSchoolByAdminID(ctx, adminID)
	if err != nil {
		return uuid.Nil, fmt.Errorf("get school: %w", err)
	}
	return school.ID, nil
}

func (s *SchoolSvc) EnrollStudent(ctx context.Context, adminID uuid.UUID, req *domain.EnrollStudentRequest) (*domain.Student, error) {
	schoolID, err := s.getSchoolID(ctx, adminID)
	if err != nil {
		return nil, err
	}
	return s.SchoolRepo.CreateStudent(ctx, &domain.Student{
		SchoolID:  schoolID,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Grade:     req.Grade,
	})
}

func (s *SchoolSvc) GetStudents(ctx context.Context, adminID uuid.UUID) ([]*domain.Student, error) {
	schoolID, err := s.getSchoolID(ctx, adminID)
	if err != nil {
		return nil, err
	}
	return s.SchoolRepo.GetStudentsBySchoolID(ctx, schoolID)
}

func (s *SchoolSvc) GetStudent(ctx context.Context, adminID uuid.UUID, studentID uuid.UUID) (*domain.Student, error) {
	schoolID, err := s.getSchoolID(ctx, adminID)
	if err != nil {
		return nil, err
	}
	student, err := s.SchoolRepo.GetStudentByID(ctx, studentID)
	if err != nil {
		return nil, err
	}
	if student.SchoolID != schoolID {
		return nil, domain.ErrNotFound
	}
	return student, nil
}

func (s *SchoolSvc) GetStudentGuardians(ctx context.Context, adminID uuid.UUID, studentID uuid.UUID) ([]*domain.Guardian, error) {
	if _, err := s.GetStudent(ctx, adminID, studentID); err != nil {
		return nil, err
	}
	return s.SchoolRepo.GetStudentGuardians(ctx, studentID)
}

func (s *SchoolSvc) AddStudentGuardian(ctx context.Context, adminID uuid.UUID, studentID uuid.UUID, req *domain.AddStudentGuardianRequest) (*domain.StudentGuardian, error) {
	if _, err := s.GetStudent(ctx, adminID, studentID); err != nil {
		return nil, err
	}
	guardianID, err := uuid.Parse(req.GuardianID)
	if err != nil {
		return nil, fmt.Errorf("parse guardian id: %w", err)
	}
	if _, err := s.SchoolRepo.GetGuardianByID(ctx, guardianID); err != nil {
		return nil, err
	}
	return s.SchoolRepo.AddStudentGuardian(ctx, &domain.StudentGuardian{
		StudentID:    studentID,
		GuardianID:   guardianID,
		Relationship: req.Relationship,
		IsPrimary:    req.IsPrimary,
	})
}

func (s *SchoolSvc) ManageStudent(ctx context.Context, adminID uuid.UUID, studentID uuid.UUID, req *domain.ManageStudentRequest) (*domain.Student, error) {
	student, err := s.GetStudent(ctx, adminID, studentID)
	if err != nil {
		return nil, err
	}
	student.IsActive = req.IsActive
	return s.SchoolRepo.UpdateStudent(ctx, student)
}

func (s *SchoolSvc) FilterStudents(ctx context.Context, adminID uuid.UUID, grade string, isActive *bool) ([]*domain.Student, error) {
	schoolID, err := s.getSchoolID(ctx, adminID)
	if err != nil {
		return nil, err
	}
	return s.SchoolRepo.FilterStudents(ctx, schoolID, grade, isActive)
}

func (s *SchoolSvc) SearchStudents(ctx context.Context, adminID uuid.UUID, query string) ([]*domain.Student, error) {
	schoolID, err := s.getSchoolID(ctx, adminID)
	if err != nil {
		return nil, err
	}
	return s.SchoolRepo.SearchStudents(ctx, schoolID, query)
}

func (s *SchoolSvc) RemoveStudent(ctx context.Context, adminID uuid.UUID, studentID uuid.UUID) error {
	if _, err := s.GetStudent(ctx, adminID, studentID); err != nil {
		return err
	}
	return s.SchoolRepo.DeactivateStudent(ctx, studentID)
}

func (s *SchoolSvc) EditStudent(ctx context.Context, adminID uuid.UUID, studentID uuid.UUID, req *domain.EditStudentRequest) (*domain.Student, error) {
	student, err := s.GetStudent(ctx, adminID, studentID)
	if err != nil {
		return nil, err
	}
	student.FirstName = req.FirstName
	student.LastName = req.LastName
	student.Grade = req.Grade
	return s.SchoolRepo.UpdateStudent(ctx, student)
}

package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jesee-kuya/transport-system/domain"
)

func (r *GuardianRepositoryStruct) GetGuardianByUserID(ctx context.Context, userID uuid.UUID) (*domain.Guardian, error) {
	var guardian domain.Guardian
	query := `SELECT id, user_id, first_name, last_name, phone, created_at, updated_at FROM guardians WHERE user_id = $1`
	err := r.DB.GetContext(ctx, &guardian, query, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrNotFound
		}
		return nil, fmt.Errorf("get guardian by user id: %w", err)
	}
	return &guardian, nil
}

func (r *GuardianRepositoryStruct) GetStudentsByGuardianID(ctx context.Context, guardianID uuid.UUID) ([]*domain.Student, error) {
	var students []*domain.Student
	query := `
		SELECT s.id, s.school_id, s.first_name, s.last_name, s.grade, s.is_active, s.created_at, s.updated_at
		FROM students s
		JOIN student_guardians sg ON sg.student_id = s.id
		WHERE sg.guardian_id = $1
		ORDER BY s.last_name, s.first_name
	`
	if err := r.DB.SelectContext(ctx, &students, query, guardianID); err != nil {
		return nil, fmt.Errorf("get students by guardian: %w", err)
	}
	return students, nil
}

func (r *GuardianRepositoryStruct) GetStudentByIDForGuardian(ctx context.Context, guardianID uuid.UUID, studentID uuid.UUID) (*domain.Student, error) {
	var student domain.Student
	query := `
		SELECT s.id, s.school_id, s.first_name, s.last_name, s.grade, s.is_active, s.created_at, s.updated_at
		FROM students s
		JOIN student_guardians sg ON sg.student_id = s.id
		WHERE sg.guardian_id = $1 AND s.id = $2
	`
	err := r.DB.GetContext(ctx, &student, query, guardianID, studentID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrNotFound
		}
		return nil, fmt.Errorf("get student by id for guardian: %w", err)
	}
	return &student, nil
}

func (r *GuardianRepositoryStruct) UpdateGuardian(ctx context.Context, guardian *domain.Guardian) (*domain.Guardian, error) {
	query := `
		UPDATE guardians SET first_name = :first_name, last_name = :last_name, phone = :phone, updated_at = now()
		WHERE id = :id
		RETURNING id, user_id, first_name, last_name, phone, created_at, updated_at
	`
	rows, err := r.DB.NamedQueryContext(ctx, query, guardian)
	if err != nil {
		return nil, fmt.Errorf("update guardian: %w", err)
	}
	defer rows.Close()

	var updated domain.Guardian
	if rows.Next() {
		if err := rows.StructScan(&updated); err != nil {
			return nil, fmt.Errorf("scan updated guardian: %w", err)
		}
	} else {
		return nil, domain.ErrNotFound
	}
	return &updated, nil
}

func (r *GuardianRepositoryStruct) GetActiveSchoolTripForStudent(ctx context.Context, studentID uuid.UUID) (*domain.Trip, error) {
	var trip domain.Trip
	query := `
		SELECT t.id, t.school_id, t.driver_id, t.bus_id, t.trip_type, t.status, t.started_at, t.ended_at, t.created_at, t.updated_at
		FROM trips t
		JOIN trip_students ts ON ts.trip_id = t.id
		WHERE ts.student_id = $1 AND t.status = 'in_progress'
		LIMIT 1
	`
	err := r.DB.GetContext(ctx, &trip, query, studentID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrNotFound
		}
		return nil, fmt.Errorf("get active trip for student: %w", err)
	}
	return &trip, nil
}

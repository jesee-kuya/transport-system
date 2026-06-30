package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jesee-kuya/transport-system/domain"
)

func (r *SchoolDriverRepositoryStruct) GetSchoolDriverByUserID(ctx context.Context, userID uuid.UUID) (*domain.SchoolDriver, error) {
	var driver domain.SchoolDriver
	query := `SELECT id, user_id, school_id, license_number, is_active, created_at, updated_at FROM school_drivers WHERE user_id = $1`
	err := r.DB.GetContext(ctx, &driver, query, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrNotFound
		}
		return nil, fmt.Errorf("get school driver by user id: %w", err)
	}
	return &driver, nil
}

func (r *SchoolDriverRepositoryStruct) SearchStudentsBySchoolID(ctx context.Context, schoolID uuid.UUID, query string) ([]*domain.Student, error) {
	var students []*domain.Student
	q := `
		SELECT id, school_id, first_name, last_name, grade, is_active, created_at, updated_at
		FROM students
		WHERE school_id = $1 AND is_active = true
		  AND (first_name ILIKE '%' || $2 || '%' OR last_name ILIKE '%' || $2 || '%')
		ORDER BY last_name, first_name
	`
	if err := r.DB.SelectContext(ctx, &students, q, schoolID, query); err != nil {
		return nil, fmt.Errorf("search students by school: %w", err)
	}
	return students, nil
}

func (r *SchoolDriverRepositoryStruct) CreateTrip(ctx context.Context, trip *domain.Trip) (*domain.Trip, error) {
	query := `
		INSERT INTO trips (school_id, driver_id, bus_id, trip_type, status, started_at)
		VALUES (:school_id, :driver_id, :bus_id, :trip_type, :status, :started_at)
		RETURNING id, school_id, driver_id, bus_id, trip_type, status, started_at, ended_at, created_at, updated_at
	`
	rows, err := r.DB.NamedQueryContext(ctx, query, trip)
	if err != nil {
		return nil, fmt.Errorf("create trip: %w", err)
	}
	defer rows.Close()

	var created domain.Trip
	if rows.Next() {
		if err := rows.StructScan(&created); err != nil {
			return nil, fmt.Errorf("scan trip: %w", err)
		}
	}
	return &created, nil
}

func (r *SchoolDriverRepositoryStruct) GetTripByID(ctx context.Context, id uuid.UUID) (*domain.Trip, error) {
	var trip domain.Trip
	query := `SELECT id, school_id, driver_id, bus_id, trip_type, status, started_at, ended_at, created_at, updated_at FROM trips WHERE id = $1`
	err := r.DB.GetContext(ctx, &trip, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrNotFound
		}
		return nil, fmt.Errorf("get trip by id: %w", err)
	}
	return &trip, nil
}

func (r *SchoolDriverRepositoryStruct) GetActiveTripByDriverID(ctx context.Context, driverID uuid.UUID) (*domain.Trip, error) {
	var trip domain.Trip
	query := `SELECT id, school_id, driver_id, bus_id, trip_type, status, started_at, ended_at, created_at, updated_at FROM trips WHERE driver_id = $1 AND status = 'in_progress' LIMIT 1`
	err := r.DB.GetContext(ctx, &trip, query, driverID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrNotFound
		}
		return nil, fmt.Errorf("get active trip by driver id: %w", err)
	}
	return &trip, nil
}

func (r *SchoolDriverRepositoryStruct) UpdateTrip(ctx context.Context, trip *domain.Trip) (*domain.Trip, error) {
	query := `
		UPDATE trips SET status = :status, ended_at = :ended_at, updated_at = now()
		WHERE id = :id
		RETURNING id, school_id, driver_id, bus_id, trip_type, status, started_at, ended_at, created_at, updated_at
	`
	rows, err := r.DB.NamedQueryContext(ctx, query, trip)
	if err != nil {
		return nil, fmt.Errorf("update trip: %w", err)
	}
	defer rows.Close()

	var updated domain.Trip
	if rows.Next() {
		if err := rows.StructScan(&updated); err != nil {
			return nil, fmt.Errorf("scan updated trip: %w", err)
		}
	} else {
		return nil, domain.ErrNotFound
	}
	return &updated, nil
}

func (r *SchoolDriverRepositoryStruct) GetStudentByID(ctx context.Context, id uuid.UUID) (*domain.Student, error) {
	var student domain.Student
	query := `SELECT id, school_id, first_name, last_name, grade, is_active, created_at, updated_at FROM students WHERE id = $1`
	err := r.DB.GetContext(ctx, &student, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrNotFound
		}
		return nil, fmt.Errorf("get student by id: %w", err)
	}
	return &student, nil
}

func (r *SchoolDriverRepositoryStruct) AddTripStudent(ctx context.Context, ts *domain.TripStudent) (*domain.TripStudent, error) {
	query := `
		INSERT INTO trip_students (trip_id, student_id)
		VALUES (:trip_id, :student_id)
		RETURNING id, trip_id, student_id, boarded_at
	`
	rows, err := r.DB.NamedQueryContext(ctx, query, ts)
	if err != nil {
		return nil, fmt.Errorf("add trip student: %w", err)
	}
	defer rows.Close()

	var created domain.TripStudent
	if rows.Next() {
		if err := rows.StructScan(&created); err != nil {
			return nil, fmt.Errorf("scan trip student: %w", err)
		}
	}
	return &created, nil
}

func (r *SchoolDriverRepositoryStruct) GetTripStudentByTripAndStudentID(ctx context.Context, tripID, studentID uuid.UUID) (*domain.TripStudent, error) {
	var ts domain.TripStudent
	query := `SELECT id, trip_id, student_id, boarded_at FROM trip_students WHERE trip_id = $1 AND student_id = $2`
	err := r.DB.GetContext(ctx, &ts, query, tripID, studentID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrNotFound
		}
		return nil, fmt.Errorf("get trip student: %w", err)
	}
	return &ts, nil
}

func (r *SchoolDriverRepositoryStruct) GetBoardedStudents(ctx context.Context, tripID uuid.UUID) ([]*domain.Student, error) {
	var students []*domain.Student
	query := `
		SELECT s.id, s.school_id, s.first_name, s.last_name, s.grade, s.is_active, s.created_at, s.updated_at
		FROM students s
		JOIN trip_students ts ON ts.student_id = s.id
		WHERE ts.trip_id = $1
		ORDER BY s.last_name, s.first_name
	`
	if err := r.DB.SelectContext(ctx, &students, query, tripID); err != nil {
		return nil, fmt.Errorf("get boarded students: %w", err)
	}
	return students, nil
}

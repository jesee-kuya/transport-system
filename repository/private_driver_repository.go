package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jesee-kuya/transport-system/domain"
)

func (r *PrivateDriverRepositoryStruct) CreatePrivateDriver(ctx context.Context, driver *domain.PrivateDriver) (*domain.PrivateDriver, error) {
	query := `
		INSERT INTO private_drivers (user_id, first_name, last_name, phone, license_number)
		VALUES (:user_id, :first_name, :last_name, :phone, :license_number)
		RETURNING id, user_id, first_name, last_name, phone, license_number, kyc_verified, created_at, updated_at
	`
	rows, err := r.DB.NamedQueryContext(ctx, query, driver)
	if err != nil {
		return nil, fmt.Errorf("create private driver: %w", err)
	}
	defer rows.Close()

	var created domain.PrivateDriver
	if rows.Next() {
		if err := rows.StructScan(&created); err != nil {
			return nil, fmt.Errorf("scan private driver: %w", err)
		}
	}
	return &created, nil
}

func (r *PrivateDriverRepositoryStruct) GetPrivateDriverByUserID(ctx context.Context, userID uuid.UUID) (*domain.PrivateDriver, error) {
	var driver domain.PrivateDriver
	query := `SELECT id, user_id, first_name, last_name, phone, license_number, kyc_verified, created_at, updated_at FROM private_drivers WHERE user_id = $1`
	err := r.DB.GetContext(ctx, &driver, query, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrNotFound
		}
		return nil, fmt.Errorf("get private driver by user id: %w", err)
	}
	return &driver, nil
}

func (r *PrivateDriverRepositoryStruct) GetMatchByID(ctx context.Context, matchID uuid.UUID) (*domain.DriverParentMatch, error) {
	var match domain.DriverParentMatch
	query := `SELECT id, parent_id, driver_id, status, created_at, updated_at FROM driver_parent_matches WHERE id = $1`
	err := r.DB.GetContext(ctx, &match, query, matchID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrNotFound
		}
		return nil, fmt.Errorf("get match by id: %w", err)
	}
	return &match, nil
}

func (r *PrivateDriverRepositoryStruct) UpdateMatchStatus(ctx context.Context, matchID uuid.UUID, status string) (*domain.DriverParentMatch, error) {
	var match domain.DriverParentMatch
	query := `
		UPDATE driver_parent_matches SET status = $2, updated_at = now()
		WHERE id = $1
		RETURNING id, parent_id, driver_id, status, created_at, updated_at
	`
	err := r.DB.GetContext(ctx, &match, query, matchID, status)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrNotFound
		}
		return nil, fmt.Errorf("update match status: %w", err)
	}
	return &match, nil
}

func (r *PrivateDriverRepositoryStruct) CreatePrivateTrip(ctx context.Context, trip *domain.PrivateTrip) (*domain.PrivateTrip, error) {
	query := `
		INSERT INTO private_trips (driver_id, match_id, status, started_at)
		VALUES (:driver_id, :match_id, :status, :started_at)
		RETURNING id, driver_id, match_id, status, started_at, ended_at, created_at, updated_at
	`
	rows, err := r.DB.NamedQueryContext(ctx, query, trip)
	if err != nil {
		return nil, fmt.Errorf("create private trip: %w", err)
	}
	defer rows.Close()

	var created domain.PrivateTrip
	if rows.Next() {
		if err := rows.StructScan(&created); err != nil {
			return nil, fmt.Errorf("scan private trip: %w", err)
		}
	}
	return &created, nil
}

func (r *PrivateDriverRepositoryStruct) GetActivePrivateTripByDriverID(ctx context.Context, driverID uuid.UUID) (*domain.PrivateTrip, error) {
	var trip domain.PrivateTrip
	query := `SELECT id, driver_id, match_id, status, started_at, ended_at, created_at, updated_at FROM private_trips WHERE driver_id = $1 AND status = 'in_progress' LIMIT 1`
	err := r.DB.GetContext(ctx, &trip, query, driverID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrNotFound
		}
		return nil, fmt.Errorf("get active private trip by driver id: %w", err)
	}
	return &trip, nil
}

func (r *PrivateDriverRepositoryStruct) GetPrivateTripByID(ctx context.Context, id uuid.UUID) (*domain.PrivateTrip, error) {
	var trip domain.PrivateTrip
	query := `SELECT id, driver_id, match_id, status, started_at, ended_at, created_at, updated_at FROM private_trips WHERE id = $1`
	err := r.DB.GetContext(ctx, &trip, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrNotFound
		}
		return nil, fmt.Errorf("get private trip by id: %w", err)
	}
	return &trip, nil
}

func (r *PrivateDriverRepositoryStruct) UpdatePrivateTrip(ctx context.Context, trip *domain.PrivateTrip) (*domain.PrivateTrip, error) {
	query := `
		UPDATE private_trips SET status = :status, ended_at = :ended_at, updated_at = now()
		WHERE id = :id
		RETURNING id, driver_id, match_id, status, started_at, ended_at, created_at, updated_at
	`
	rows, err := r.DB.NamedQueryContext(ctx, query, trip)
	if err != nil {
		return nil, fmt.Errorf("update private trip: %w", err)
	}
	defer rows.Close()

	var updated domain.PrivateTrip
	if rows.Next() {
		if err := rows.StructScan(&updated); err != nil {
			return nil, fmt.Errorf("scan updated private trip: %w", err)
		}
	} else {
		return nil, domain.ErrNotFound
	}
	return &updated, nil
}

func (r *PrivateDriverRepositoryStruct) AddPrivateTripChild(ctx context.Context, ptc *domain.PrivateTripChild) (*domain.PrivateTripChild, error) {
	query := `
		INSERT INTO private_trip_children (trip_id, child_id)
		VALUES (:trip_id, :child_id)
		RETURNING id, trip_id, child_id, boarded_at
	`
	rows, err := r.DB.NamedQueryContext(ctx, query, ptc)
	if err != nil {
		return nil, fmt.Errorf("add private trip child: %w", err)
	}
	defer rows.Close()

	var created domain.PrivateTripChild
	if rows.Next() {
		if err := rows.StructScan(&created); err != nil {
			return nil, fmt.Errorf("scan private trip child: %w", err)
		}
	}
	return &created, nil
}

func (r *PrivateDriverRepositoryStruct) GetPrivateTripChildByTripAndChildID(ctx context.Context, tripID, childID uuid.UUID) (*domain.PrivateTripChild, error) {
	var ptc domain.PrivateTripChild
	query := `SELECT id, trip_id, child_id, boarded_at FROM private_trip_children WHERE trip_id = $1 AND child_id = $2`
	err := r.DB.GetContext(ctx, &ptc, query, tripID, childID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrNotFound
		}
		return nil, fmt.Errorf("get private trip child: %w", err)
	}
	return &ptc, nil
}

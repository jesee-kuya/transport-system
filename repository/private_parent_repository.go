package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jesee-kuya/transport-system/domain"
)

// Profile

func (r *PrivateParentRepositoryStruct) CreatePrivateParent(ctx context.Context, parent *domain.PrivateParent) (*domain.PrivateParent, error) {
	query := `
		INSERT INTO private_parents (user_id, first_name, last_name, phone)
		VALUES (:user_id, :first_name, :last_name, :phone)
		RETURNING id, user_id, first_name, last_name, phone, kyc_verified, created_at, updated_at
	`
	rows, err := r.DB.NamedQueryContext(ctx, query, parent)
	if err != nil {
		return nil, fmt.Errorf("create private parent: %w", err)
	}
	defer rows.Close()

	var created domain.PrivateParent
	if rows.Next() {
		if err := rows.StructScan(&created); err != nil {
			return nil, fmt.Errorf("scan private parent: %w", err)
		}
	}
	return &created, nil
}

func (r *PrivateParentRepositoryStruct) GetPrivateParentByUserID(ctx context.Context, userID uuid.UUID) (*domain.PrivateParent, error) {
	var parent domain.PrivateParent
	query := `SELECT id, user_id, first_name, last_name, phone, kyc_verified, created_at, updated_at FROM private_parents WHERE user_id = $1`
	err := r.DB.GetContext(ctx, &parent, query, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrNotFound
		}
		return nil, fmt.Errorf("get private parent by user id: %w", err)
	}
	return &parent, nil
}

func (r *PrivateParentRepositoryStruct) UpdatePrivateParent(ctx context.Context, parent *domain.PrivateParent) (*domain.PrivateParent, error) {
	query := `
		UPDATE private_parents SET first_name = :first_name, last_name = :last_name, phone = :phone, updated_at = now()
		WHERE id = :id
		RETURNING id, user_id, first_name, last_name, phone, kyc_verified, created_at, updated_at
	`
	rows, err := r.DB.NamedQueryContext(ctx, query, parent)
	if err != nil {
		return nil, fmt.Errorf("update private parent: %w", err)
	}
	defer rows.Close()

	var updated domain.PrivateParent
	if rows.Next() {
		if err := rows.StructScan(&updated); err != nil {
			return nil, fmt.Errorf("scan updated private parent: %w", err)
		}
	} else {
		return nil, domain.ErrNotFound
	}
	return &updated, nil
}

func (r *PrivateParentRepositoryStruct) DeleteAccount(ctx context.Context, userID uuid.UUID) error {
	tx, err := r.DB.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("begin transaction: %w", err)
	}
	defer tx.Rollback()

	if _, err := tx.ExecContext(ctx, `DELETE FROM private_parents WHERE user_id = $1`, userID); err != nil {
		return fmt.Errorf("delete private parent: %w", err)
	}
	if _, err := tx.ExecContext(ctx, `UPDATE users SET is_active = false, updated_at = now() WHERE id = $1`, userID); err != nil {
		return fmt.Errorf("deactivate user: %w", err)
	}
	return tx.Commit()
}

// Children

func (r *PrivateParentRepositoryStruct) CreatePrivateChild(ctx context.Context, child *domain.PrivateChild) (*domain.PrivateChild, error) {
	query := `
		INSERT INTO private_children (parent_id, first_name, last_name, grade)
		VALUES (:parent_id, :first_name, :last_name, :grade)
		RETURNING id, parent_id, school_id, first_name, last_name, grade, created_at, updated_at
	`
	rows, err := r.DB.NamedQueryContext(ctx, query, child)
	if err != nil {
		return nil, fmt.Errorf("create private child: %w", err)
	}
	defer rows.Close()

	var created domain.PrivateChild
	if rows.Next() {
		if err := rows.StructScan(&created); err != nil {
			return nil, fmt.Errorf("scan private child: %w", err)
		}
	}
	return &created, nil
}

func (r *PrivateParentRepositoryStruct) GetPrivateChildrenByParentID(ctx context.Context, parentID uuid.UUID) ([]*domain.PrivateChild, error) {
	var children []*domain.PrivateChild
	query := `SELECT id, parent_id, school_id, first_name, last_name, grade, created_at, updated_at FROM private_children WHERE parent_id = $1 ORDER BY last_name, first_name`
	if err := r.DB.SelectContext(ctx, &children, query, parentID); err != nil {
		return nil, fmt.Errorf("get private children: %w", err)
	}
	return children, nil
}

func (r *PrivateParentRepositoryStruct) GetPrivateChildByID(ctx context.Context, id uuid.UUID) (*domain.PrivateChild, error) {
	var child domain.PrivateChild
	query := `SELECT id, parent_id, school_id, first_name, last_name, grade, created_at, updated_at FROM private_children WHERE id = $1`
	err := r.DB.GetContext(ctx, &child, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrNotFound
		}
		return nil, fmt.Errorf("get private child by id: %w", err)
	}
	return &child, nil
}

func (r *PrivateParentRepositoryStruct) UpdatePrivateChild(ctx context.Context, child *domain.PrivateChild) (*domain.PrivateChild, error) {
	query := `
		UPDATE private_children SET first_name = :first_name, last_name = :last_name, grade = :grade, updated_at = now()
		WHERE id = :id
		RETURNING id, parent_id, school_id, first_name, last_name, grade, created_at, updated_at
	`
	rows, err := r.DB.NamedQueryContext(ctx, query, child)
	if err != nil {
		return nil, fmt.Errorf("update private child: %w", err)
	}
	defer rows.Close()

	var updated domain.PrivateChild
	if rows.Next() {
		if err := rows.StructScan(&updated); err != nil {
			return nil, fmt.Errorf("scan updated private child: %w", err)
		}
	} else {
		return nil, domain.ErrNotFound
	}
	return &updated, nil
}

// Driver matching

func (r *PrivateParentRepositoryStruct) CreateDriverParentMatch(ctx context.Context, match *domain.DriverParentMatch) (*domain.DriverParentMatch, error) {
	query := `
		INSERT INTO driver_parent_matches (parent_id, driver_id, status)
		VALUES (:parent_id, :driver_id, :status)
		RETURNING id, parent_id, driver_id, status, created_at, updated_at
	`
	rows, err := r.DB.NamedQueryContext(ctx, query, match)
	if err != nil {
		return nil, fmt.Errorf("create driver parent match: %w", err)
	}
	defer rows.Close()

	var created domain.DriverParentMatch
	if rows.Next() {
		if err := rows.StructScan(&created); err != nil {
			return nil, fmt.Errorf("scan driver parent match: %w", err)
		}
	}
	return &created, nil
}

func (r *PrivateParentRepositoryStruct) GetMatchByParentAndDriverID(ctx context.Context, parentID, driverID uuid.UUID) (*domain.DriverParentMatch, error) {
	var match domain.DriverParentMatch
	query := `SELECT id, parent_id, driver_id, status, created_at, updated_at FROM driver_parent_matches WHERE parent_id = $1 AND driver_id = $2`
	err := r.DB.GetContext(ctx, &match, query, parentID, driverID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrNotFound
		}
		return nil, fmt.Errorf("get match by parent and driver: %w", err)
	}
	return &match, nil
}

// School connections

func (r *PrivateParentRepositoryStruct) CreateParentSchoolConnection(ctx context.Context, conn *domain.ParentSchoolConnection) (*domain.ParentSchoolConnection, error) {
	query := `
		INSERT INTO parent_school_connections (parent_id, school_id)
		VALUES (:parent_id, :school_id)
		RETURNING id, parent_id, school_id, created_at
	`
	rows, err := r.DB.NamedQueryContext(ctx, query, conn)
	if err != nil {
		return nil, fmt.Errorf("create parent school connection: %w", err)
	}
	defer rows.Close()

	var created domain.ParentSchoolConnection
	if rows.Next() {
		if err := rows.StructScan(&created); err != nil {
			return nil, fmt.Errorf("scan parent school connection: %w", err)
		}
	}
	return &created, nil
}

func (r *PrivateParentRepositoryStruct) GetParentSchoolConnection(ctx context.Context, parentID, schoolID uuid.UUID) (*domain.ParentSchoolConnection, error) {
	var conn domain.ParentSchoolConnection
	query := `SELECT id, parent_id, school_id, created_at FROM parent_school_connections WHERE parent_id = $1 AND school_id = $2`
	err := r.DB.GetContext(ctx, &conn, query, parentID, schoolID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrNotFound
		}
		return nil, fmt.Errorf("get parent school connection: %w", err)
	}
	return &conn, nil
}

// School browsing

func (r *PrivateParentRepositoryStruct) GetAllSchools(ctx context.Context) ([]*domain.School, error) {
	var schools []*domain.School
	query := `SELECT id, admin_id, name, address, contact_email, contact_phone, created_at, updated_at FROM schools ORDER BY name`
	if err := r.DB.SelectContext(ctx, &schools, query); err != nil {
		return nil, fmt.Errorf("get all schools: %w", err)
	}
	return schools, nil
}

func (r *PrivateParentRepositoryStruct) SearchSchools(ctx context.Context, query string) ([]*domain.School, error) {
	var schools []*domain.School
	q := `
		SELECT id, admin_id, name, address, contact_email, contact_phone, created_at, updated_at
		FROM schools
		WHERE name ILIKE '%' || $1 || '%'
		ORDER BY name
	`
	if err := r.DB.SelectContext(ctx, &schools, q, query); err != nil {
		return nil, fmt.Errorf("search schools: %w", err)
	}
	return schools, nil
}

func (r *PrivateParentRepositoryStruct) FilterSchools(ctx context.Context, address string) ([]*domain.School, error) {
	var schools []*domain.School
	query := `
		SELECT id, admin_id, name, address, contact_email, contact_phone, created_at, updated_at
		FROM schools
		WHERE ($1 = '' OR address ILIKE '%' || $1 || '%')
		ORDER BY name
	`
	if err := r.DB.SelectContext(ctx, &schools, query, address); err != nil {
		return nil, fmt.Errorf("filter schools: %w", err)
	}
	return schools, nil
}

func (r *PrivateParentRepositoryStruct) GetSchoolByID(ctx context.Context, id uuid.UUID) (*domain.School, error) {
	var school domain.School
	query := `SELECT id, admin_id, name, address, contact_email, contact_phone, created_at, updated_at FROM schools WHERE id = $1`
	err := r.DB.GetContext(ctx, &school, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrNotFound
		}
		return nil, fmt.Errorf("get school by id: %w", err)
	}
	return &school, nil
}

// Trips

func (r *PrivateParentRepositoryStruct) GetPrivateTripsByParentID(ctx context.Context, parentID uuid.UUID) ([]*domain.PrivateTrip, error) {
	var trips []*domain.PrivateTrip
	query := `
		SELECT pt.id, pt.driver_id, pt.match_id, pt.status, pt.started_at, pt.ended_at, pt.created_at, pt.updated_at
		FROM private_trips pt
		JOIN driver_parent_matches m ON m.id = pt.match_id
		WHERE m.parent_id = $1
		ORDER BY pt.created_at DESC
	`
	if err := r.DB.SelectContext(ctx, &trips, query, parentID); err != nil {
		return nil, fmt.Errorf("get private trips by parent: %w", err)
	}
	return trips, nil
}

func (r *PrivateParentRepositoryStruct) GetPrivateTripByIDForParent(ctx context.Context, parentID, tripID uuid.UUID) (*domain.PrivateTrip, error) {
	var trip domain.PrivateTrip
	query := `
		SELECT pt.id, pt.driver_id, pt.match_id, pt.status, pt.started_at, pt.ended_at, pt.created_at, pt.updated_at
		FROM private_trips pt
		JOIN driver_parent_matches m ON m.id = pt.match_id
		WHERE m.parent_id = $1 AND pt.id = $2
	`
	err := r.DB.GetContext(ctx, &trip, query, parentID, tripID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrNotFound
		}
		return nil, fmt.Errorf("get private trip by id for parent: %w", err)
	}
	return &trip, nil
}

func (r *PrivateParentRepositoryStruct) UpdatePrivateTripStatus(ctx context.Context, tripID uuid.UUID, status string) (*domain.PrivateTrip, error) {
	var trip domain.PrivateTrip
	query := `
		UPDATE private_trips SET status = $2, updated_at = now()
		WHERE id = $1
		RETURNING id, driver_id, match_id, status, started_at, ended_at, created_at, updated_at
	`
	err := r.DB.GetContext(ctx, &trip, query, tripID, status)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrNotFound
		}
		return nil, fmt.Errorf("update private trip status: %w", err)
	}
	return &trip, nil
}

func (r *PrivateParentRepositoryStruct) GetPrivateDriverByID(ctx context.Context, id uuid.UUID) (*domain.PrivateDriver, error) {
	var driver domain.PrivateDriver
	query := `SELECT id, user_id, first_name, last_name, phone, license_number, kyc_verified, created_at, updated_at FROM private_drivers WHERE id = $1`
	err := r.DB.GetContext(ctx, &driver, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrNotFound
		}
		return nil, fmt.Errorf("get private driver by id: %w", err)
	}
	return &driver, nil
}

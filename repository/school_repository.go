package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jesee-kuya/transport-system/domain"
)

// School

func (r *SchoolRepositoryStruct) CreateSchool(ctx context.Context, school *domain.School) (*domain.School, error) {
	query := `
		INSERT INTO schools (admin_id, name, address, contact_email, contact_phone)
		VALUES (:admin_id, :name, :address, :contact_email, :contact_phone)
		RETURNING id, admin_id, name, address, contact_email, contact_phone, created_at, updated_at
	`
	rows, err := r.DB.NamedQueryContext(ctx, query, school)
	if err != nil {
		return nil, fmt.Errorf("create school: %w", err)
	}
	defer rows.Close()

	var created domain.School
	if rows.Next() {
		if err := rows.StructScan(&created); err != nil {
			return nil, fmt.Errorf("scan school: %w", err)
		}
	}
	return &created, nil
}

func (r *SchoolRepositoryStruct) GetSchoolByAdminID(ctx context.Context, adminID uuid.UUID) (*domain.School, error) {
	var school domain.School
	query := `SELECT id, admin_id, name, address, contact_email, contact_phone, created_at, updated_at FROM schools WHERE admin_id = $1`
	err := r.DB.GetContext(ctx, &school, query, adminID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrNotFound
		}
		return nil, fmt.Errorf("get school by admin id: %w", err)
	}
	return &school, nil
}

// Students

func (r *SchoolRepositoryStruct) CreateStudent(ctx context.Context, student *domain.Student) (*domain.Student, error) {
	query := `
		INSERT INTO students (school_id, first_name, last_name, grade)
		VALUES (:school_id, :first_name, :last_name, :grade)
		RETURNING id, school_id, first_name, last_name, grade, is_active, created_at, updated_at
	`
	rows, err := r.DB.NamedQueryContext(ctx, query, student)
	if err != nil {
		return nil, fmt.Errorf("create student: %w", err)
	}
	defer rows.Close()

	var created domain.Student
	if rows.Next() {
		if err := rows.StructScan(&created); err != nil {
			return nil, fmt.Errorf("scan student: %w", err)
		}
	}
	return &created, nil
}

func (r *SchoolRepositoryStruct) GetStudentsBySchoolID(ctx context.Context, schoolID uuid.UUID) ([]*domain.Student, error) {
	var students []*domain.Student
	query := `SELECT id, school_id, first_name, last_name, grade, is_active, created_at, updated_at FROM students WHERE school_id = $1 ORDER BY last_name, first_name`
	if err := r.DB.SelectContext(ctx, &students, query, schoolID); err != nil {
		return nil, fmt.Errorf("get students: %w", err)
	}
	return students, nil
}

func (r *SchoolRepositoryStruct) GetStudentByID(ctx context.Context, id uuid.UUID) (*domain.Student, error) {
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

func (r *SchoolRepositoryStruct) UpdateStudent(ctx context.Context, student *domain.Student) (*domain.Student, error) {
	query := `
		UPDATE students SET first_name = :first_name, last_name = :last_name, grade = :grade, updated_at = now()
		WHERE id = :id
		RETURNING id, school_id, first_name, last_name, grade, is_active, created_at, updated_at
	`
	rows, err := r.DB.NamedQueryContext(ctx, query, student)
	if err != nil {
		return nil, fmt.Errorf("update student: %w", err)
	}
	defer rows.Close()

	var updated domain.Student
	if rows.Next() {
		if err := rows.StructScan(&updated); err != nil {
			return nil, fmt.Errorf("scan updated student: %w", err)
		}
	}
	return &updated, nil
}

func (r *SchoolRepositoryStruct) DeactivateStudent(ctx context.Context, id uuid.UUID) error {
	_, err := r.DB.ExecContext(ctx, `UPDATE students SET is_active = false, updated_at = now() WHERE id = $1`, id)
	if err != nil {
		return fmt.Errorf("deactivate student: %w", err)
	}
	return nil
}

func (r *SchoolRepositoryStruct) FilterStudents(ctx context.Context, schoolID uuid.UUID, grade string, isActive *bool) ([]*domain.Student, error) {
	var students []*domain.Student
	query := `
		SELECT id, school_id, first_name, last_name, grade, is_active, created_at, updated_at
		FROM students
		WHERE school_id = $1
		  AND ($2 = '' OR grade = $2)
		  AND ($3::boolean IS NULL OR is_active = $3)
		ORDER BY last_name, first_name
	`
	if err := r.DB.SelectContext(ctx, &students, query, schoolID, grade, isActive); err != nil {
		return nil, fmt.Errorf("filter students: %w", err)
	}
	return students, nil
}

func (r *SchoolRepositoryStruct) SearchStudents(ctx context.Context, schoolID uuid.UUID, query string) ([]*domain.Student, error) {
	var students []*domain.Student
	q := `
		SELECT id, school_id, first_name, last_name, grade, is_active, created_at, updated_at
		FROM students
		WHERE school_id = $1 AND (first_name ILIKE '%' || $2 || '%' OR last_name ILIKE '%' || $2 || '%')
		ORDER BY last_name, first_name
	`
	if err := r.DB.SelectContext(ctx, &students, q, schoolID, query); err != nil {
		return nil, fmt.Errorf("search students: %w", err)
	}
	return students, nil
}

// Guardians

func (r *SchoolRepositoryStruct) GetGuardianByID(ctx context.Context, id uuid.UUID) (*domain.Guardian, error) {
	var guardian domain.Guardian
	query := `SELECT id, user_id, first_name, last_name, phone, created_at, updated_at FROM guardians WHERE id = $1`
	err := r.DB.GetContext(ctx, &guardian, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrNotFound
		}
		return nil, fmt.Errorf("get guardian by id: %w", err)
	}
	return &guardian, nil
}

func (r *SchoolRepositoryStruct) GetStudentGuardians(ctx context.Context, studentID uuid.UUID) ([]*domain.Guardian, error) {
	var guardians []*domain.Guardian
	query := `
		SELECT g.id, g.user_id, g.first_name, g.last_name, g.phone, g.created_at, g.updated_at
		FROM guardians g
		JOIN student_guardians sg ON sg.guardian_id = g.id
		WHERE sg.student_id = $1
	`
	if err := r.DB.SelectContext(ctx, &guardians, query, studentID); err != nil {
		return nil, fmt.Errorf("get student guardians: %w", err)
	}
	return guardians, nil
}

func (r *SchoolRepositoryStruct) AddStudentGuardian(ctx context.Context, sg *domain.StudentGuardian) (*domain.StudentGuardian, error) {
	query := `
		INSERT INTO student_guardians (student_id, guardian_id, relationship, is_primary)
		VALUES (:student_id, :guardian_id, :relationship, :is_primary)
		RETURNING id, student_id, guardian_id, relationship, is_primary, created_at
	`
	rows, err := r.DB.NamedQueryContext(ctx, query, sg)
	if err != nil {
		return nil, fmt.Errorf("add student guardian: %w", err)
	}
	defer rows.Close()

	var created domain.StudentGuardian
	if rows.Next() {
		if err := rows.StructScan(&created); err != nil {
			return nil, fmt.Errorf("scan student guardian: %w", err)
		}
	}
	return &created, nil
}

// Buses

func (r *SchoolRepositoryStruct) CreateBus(ctx context.Context, bus *domain.Bus) (*domain.Bus, error) {
	query := `
		INSERT INTO buses (school_id, plate_number, model, capacity)
		VALUES (:school_id, :plate_number, :model, :capacity)
		RETURNING id, school_id, plate_number, model, capacity, is_active, created_at, updated_at
	`
	rows, err := r.DB.NamedQueryContext(ctx, query, bus)
	if err != nil {
		return nil, fmt.Errorf("create bus: %w", err)
	}
	defer rows.Close()

	var created domain.Bus
	if rows.Next() {
		if err := rows.StructScan(&created); err != nil {
			return nil, fmt.Errorf("scan bus: %w", err)
		}
	}
	return &created, nil
}

func (r *SchoolRepositoryStruct) GetBusesBySchoolID(ctx context.Context, schoolID uuid.UUID) ([]*domain.Bus, error) {
	var buses []*domain.Bus
	query := `SELECT id, school_id, plate_number, model, capacity, is_active, created_at, updated_at FROM buses WHERE school_id = $1 ORDER BY plate_number`
	if err := r.DB.SelectContext(ctx, &buses, query, schoolID); err != nil {
		return nil, fmt.Errorf("get buses: %w", err)
	}
	return buses, nil
}

func (r *SchoolRepositoryStruct) GetBusByID(ctx context.Context, id uuid.UUID) (*domain.Bus, error) {
	var bus domain.Bus
	query := `SELECT id, school_id, plate_number, model, capacity, is_active, created_at, updated_at FROM buses WHERE id = $1`
	err := r.DB.GetContext(ctx, &bus, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrNotFound
		}
		return nil, fmt.Errorf("get bus by id: %w", err)
	}
	return &bus, nil
}

func (r *SchoolRepositoryStruct) UpdateBus(ctx context.Context, bus *domain.Bus) (*domain.Bus, error) {
	query := `
		UPDATE buses SET plate_number = :plate_number, model = :model, capacity = :capacity, is_active = :is_active, updated_at = now()
		WHERE id = :id
		RETURNING id, school_id, plate_number, model, capacity, is_active, created_at, updated_at
	`
	rows, err := r.DB.NamedQueryContext(ctx, query, bus)
	if err != nil {
		return nil, fmt.Errorf("update bus: %w", err)
	}
	defer rows.Close()

	var updated domain.Bus
	if rows.Next() {
		if err := rows.StructScan(&updated); err != nil {
			return nil, fmt.Errorf("scan updated bus: %w", err)
		}
	}
	return &updated, nil
}

func (r *SchoolRepositoryStruct) DeactivateBus(ctx context.Context, id uuid.UUID) error {
	_, err := r.DB.ExecContext(ctx, `UPDATE buses SET is_active = false, updated_at = now() WHERE id = $1`, id)
	if err != nil {
		return fmt.Errorf("deactivate bus: %w", err)
	}
	return nil
}

func (r *SchoolRepositoryStruct) FilterBuses(ctx context.Context, schoolID uuid.UUID, isActive *bool) ([]*domain.Bus, error) {
	var buses []*domain.Bus
	query := `
		SELECT id, school_id, plate_number, model, capacity, is_active, created_at, updated_at
		FROM buses
		WHERE school_id = $1 AND ($2::boolean IS NULL OR is_active = $2)
		ORDER BY plate_number
	`
	if err := r.DB.SelectContext(ctx, &buses, query, schoolID, isActive); err != nil {
		return nil, fmt.Errorf("filter buses: %w", err)
	}
	return buses, nil
}

func (r *SchoolRepositoryStruct) SearchBuses(ctx context.Context, schoolID uuid.UUID, query string) ([]*domain.Bus, error) {
	var buses []*domain.Bus
	q := `
		SELECT id, school_id, plate_number, model, capacity, is_active, created_at, updated_at
		FROM buses
		WHERE school_id = $1 AND (plate_number ILIKE '%' || $2 || '%' OR model ILIKE '%' || $2 || '%')
		ORDER BY plate_number
	`
	if err := r.DB.SelectContext(ctx, &buses, q, schoolID, query); err != nil {
		return nil, fmt.Errorf("search buses: %w", err)
	}
	return buses, nil
}

func (r *SchoolRepositoryStruct) GetActiveTripByBusID(ctx context.Context, busID uuid.UUID) (*domain.Trip, error) {
	var trip domain.Trip
	query := `SELECT id, school_id, driver_id, bus_id, trip_type, status, started_at, ended_at, created_at, updated_at FROM trips WHERE bus_id = $1 AND status = 'in_progress' LIMIT 1`
	err := r.DB.GetContext(ctx, &trip, query, busID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrNotFound
		}
		return nil, fmt.Errorf("get active trip by bus id: %w", err)
	}
	return &trip, nil
}

// School drivers

func (r *SchoolRepositoryStruct) CreateSchoolDriver(ctx context.Context, driver *domain.SchoolDriver) (*domain.SchoolDriver, error) {
	query := `
		INSERT INTO school_drivers (user_id, school_id, license_number)
		VALUES (:user_id, :school_id, :license_number)
		RETURNING id, user_id, school_id, license_number, is_active, created_at, updated_at
	`
	rows, err := r.DB.NamedQueryContext(ctx, query, driver)
	if err != nil {
		return nil, fmt.Errorf("create school driver: %w", err)
	}
	defer rows.Close()

	var created domain.SchoolDriver
	if rows.Next() {
		if err := rows.StructScan(&created); err != nil {
			return nil, fmt.Errorf("scan school driver: %w", err)
		}
	}
	return &created, nil
}

func (r *SchoolRepositoryStruct) GetSchoolDriversBySchoolID(ctx context.Context, schoolID uuid.UUID) ([]*domain.SchoolDriver, error) {
	var drivers []*domain.SchoolDriver
	query := `SELECT id, user_id, school_id, license_number, is_active, created_at, updated_at FROM school_drivers WHERE school_id = $1 ORDER BY created_at DESC`
	if err := r.DB.SelectContext(ctx, &drivers, query, schoolID); err != nil {
		return nil, fmt.Errorf("get school drivers: %w", err)
	}
	return drivers, nil
}

func (r *SchoolRepositoryStruct) GetSchoolDriverByID(ctx context.Context, id uuid.UUID) (*domain.SchoolDriver, error) {
	var driver domain.SchoolDriver
	query := `SELECT id, user_id, school_id, license_number, is_active, created_at, updated_at FROM school_drivers WHERE id = $1`
	err := r.DB.GetContext(ctx, &driver, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrNotFound
		}
		return nil, fmt.Errorf("get school driver by id: %w", err)
	}
	return &driver, nil
}

func (r *SchoolRepositoryStruct) UpdateSchoolDriver(ctx context.Context, driver *domain.SchoolDriver) (*domain.SchoolDriver, error) {
	query := `
		UPDATE school_drivers SET license_number = :license_number, is_active = :is_active, updated_at = now()
		WHERE id = :id
		RETURNING id, user_id, school_id, license_number, is_active, created_at, updated_at
	`
	rows, err := r.DB.NamedQueryContext(ctx, query, driver)
	if err != nil {
		return nil, fmt.Errorf("update school driver: %w", err)
	}
	defer rows.Close()

	var updated domain.SchoolDriver
	if rows.Next() {
		if err := rows.StructScan(&updated); err != nil {
			return nil, fmt.Errorf("scan updated school driver: %w", err)
		}
	}
	return &updated, nil
}

func (r *SchoolRepositoryStruct) DeactivateSchoolDriver(ctx context.Context, id uuid.UUID) error {
	_, err := r.DB.ExecContext(ctx, `UPDATE school_drivers SET is_active = false, updated_at = now() WHERE id = $1`, id)
	if err != nil {
		return fmt.Errorf("deactivate school driver: %w", err)
	}
	return nil
}

func (r *SchoolRepositoryStruct) FilterSchoolDrivers(ctx context.Context, schoolID uuid.UUID, isActive *bool) ([]*domain.SchoolDriver, error) {
	var drivers []*domain.SchoolDriver
	query := `
		SELECT id, user_id, school_id, license_number, is_active, created_at, updated_at
		FROM school_drivers
		WHERE school_id = $1 AND ($2::boolean IS NULL OR is_active = $2)
		ORDER BY created_at DESC
	`
	if err := r.DB.SelectContext(ctx, &drivers, query, schoolID, isActive); err != nil {
		return nil, fmt.Errorf("filter school drivers: %w", err)
	}
	return drivers, nil
}

func (r *SchoolRepositoryStruct) SearchSchoolDrivers(ctx context.Context, schoolID uuid.UUID, query string) ([]*domain.SchoolDriver, error) {
	var drivers []*domain.SchoolDriver
	q := `
		SELECT sd.id, sd.user_id, sd.school_id, sd.license_number, sd.is_active, sd.created_at, sd.updated_at
		FROM school_drivers sd
		JOIN users u ON u.id = sd.user_id
		WHERE sd.school_id = $1 AND (u.username ILIKE '%' || $2 || '%' OR sd.license_number ILIKE '%' || $2 || '%')
		ORDER BY sd.created_at DESC
	`
	if err := r.DB.SelectContext(ctx, &drivers, q, schoolID, query); err != nil {
		return nil, fmt.Errorf("search school drivers: %w", err)
	}
	return drivers, nil
}

// Private driver connections

func (r *SchoolRepositoryStruct) ConnectPrivateDriver(ctx context.Context, schoolID, driverID uuid.UUID) error {
	query := `INSERT INTO school_private_driver_connections (school_id, driver_id) VALUES ($1, $2)`
	_, err := r.DB.ExecContext(ctx, query, schoolID, driverID)
	if err != nil {
		return fmt.Errorf("connect private driver: %w", err)
	}
	return nil
}

func (r *SchoolRepositoryStruct) GetPrivateDriversBySchoolID(ctx context.Context, schoolID uuid.UUID) ([]*domain.PrivateDriver, error) {
	var drivers []*domain.PrivateDriver
	query := `
		SELECT pd.id, pd.user_id, pd.first_name, pd.last_name, pd.phone, pd.license_number, pd.kyc_verified, pd.created_at, pd.updated_at
		FROM private_drivers pd
		JOIN school_private_driver_connections c ON c.driver_id = pd.id
		WHERE c.school_id = $1
		ORDER BY pd.last_name, pd.first_name
	`
	if err := r.DB.SelectContext(ctx, &drivers, query, schoolID); err != nil {
		return nil, fmt.Errorf("get private drivers by school: %w", err)
	}
	return drivers, nil
}

func (r *SchoolRepositoryStruct) GetPrivateDriverByID(ctx context.Context, id uuid.UUID) (*domain.PrivateDriver, error) {
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

// Private trips

func (r *SchoolRepositoryStruct) GetActivePrivateTripsBySchoolID(ctx context.Context, schoolID uuid.UUID) ([]*domain.PrivateTrip, error) {
	var trips []*domain.PrivateTrip
	query := `
		SELECT pt.id, pt.driver_id, pt.match_id, pt.status, pt.started_at, pt.ended_at, pt.created_at, pt.updated_at
		FROM private_trips pt
		JOIN school_private_driver_connections c ON c.driver_id = pt.driver_id
		WHERE c.school_id = $1 AND pt.status = 'in_progress'
		ORDER BY pt.created_at DESC
	`
	if err := r.DB.SelectContext(ctx, &trips, query, schoolID); err != nil {
		return nil, fmt.Errorf("get active private trips: %w", err)
	}
	return trips, nil
}

func (r *SchoolRepositoryStruct) GetPrivateTripByID(ctx context.Context, id uuid.UUID) (*domain.PrivateTrip, error) {
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

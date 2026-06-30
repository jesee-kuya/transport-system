package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/jesee-kuya/transport-system/domain"
)

type AuthRepository interface {
	CreateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
	GetUserByUsername(ctx context.Context, username string) (*domain.User, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (*domain.User, error)
	UpdatePassword(ctx context.Context, id uuid.UUID, passwordHash string) error
}

type SchoolRepository interface {
	// School
	CreateSchool(ctx context.Context, school *domain.School) (*domain.School, error)
	GetSchoolByAdminID(ctx context.Context, adminID uuid.UUID) (*domain.School, error)

	// Students
	CreateStudent(ctx context.Context, student *domain.Student) (*domain.Student, error)
	GetStudentsBySchoolID(ctx context.Context, schoolID uuid.UUID) ([]*domain.Student, error)
	GetStudentByID(ctx context.Context, id uuid.UUID) (*domain.Student, error)
	UpdateStudent(ctx context.Context, student *domain.Student) (*domain.Student, error)
	DeactivateStudent(ctx context.Context, id uuid.UUID) error
	FilterStudents(ctx context.Context, schoolID uuid.UUID, grade string, isActive *bool) ([]*domain.Student, error)
	SearchStudents(ctx context.Context, schoolID uuid.UUID, query string) ([]*domain.Student, error)

	// Guardians
	GetGuardianByID(ctx context.Context, id uuid.UUID) (*domain.Guardian, error)
	GetStudentGuardians(ctx context.Context, studentID uuid.UUID) ([]*domain.Guardian, error)
	AddStudentGuardian(ctx context.Context, sg *domain.StudentGuardian) (*domain.StudentGuardian, error)

	// Buses
	CreateBus(ctx context.Context, bus *domain.Bus) (*domain.Bus, error)
	GetBusesBySchoolID(ctx context.Context, schoolID uuid.UUID) ([]*domain.Bus, error)
	GetBusByID(ctx context.Context, id uuid.UUID) (*domain.Bus, error)
	UpdateBus(ctx context.Context, bus *domain.Bus) (*domain.Bus, error)
	DeactivateBus(ctx context.Context, id uuid.UUID) error
	FilterBuses(ctx context.Context, schoolID uuid.UUID, isActive *bool) ([]*domain.Bus, error)
	SearchBuses(ctx context.Context, schoolID uuid.UUID, query string) ([]*domain.Bus, error)
	GetActiveTripByBusID(ctx context.Context, busID uuid.UUID) (*domain.Trip, error)

	// School drivers
	CreateSchoolDriver(ctx context.Context, driver *domain.SchoolDriver) (*domain.SchoolDriver, error)
	GetSchoolDriversBySchoolID(ctx context.Context, schoolID uuid.UUID) ([]*domain.SchoolDriver, error)
	GetSchoolDriverByID(ctx context.Context, id uuid.UUID) (*domain.SchoolDriver, error)
	UpdateSchoolDriver(ctx context.Context, driver *domain.SchoolDriver) (*domain.SchoolDriver, error)
	DeactivateSchoolDriver(ctx context.Context, id uuid.UUID) error
	FilterSchoolDrivers(ctx context.Context, schoolID uuid.UUID, isActive *bool) ([]*domain.SchoolDriver, error)
	SearchSchoolDrivers(ctx context.Context, schoolID uuid.UUID, query string) ([]*domain.SchoolDriver, error)

	// Private driver connections
	ConnectPrivateDriver(ctx context.Context, schoolID, driverID uuid.UUID) error
	GetPrivateDriversBySchoolID(ctx context.Context, schoolID uuid.UUID) ([]*domain.PrivateDriver, error)
	GetPrivateDriverByID(ctx context.Context, id uuid.UUID) (*domain.PrivateDriver, error)

	// Private trips (for school's connected drivers)
	GetActivePrivateTripsBySchoolID(ctx context.Context, schoolID uuid.UUID) ([]*domain.PrivateTrip, error)
	GetPrivateTripByID(ctx context.Context, id uuid.UUID) (*domain.PrivateTrip, error)
}

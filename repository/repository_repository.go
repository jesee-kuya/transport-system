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

type PrivateDriverRepository interface {
	// Profile/KYC
	CreatePrivateDriver(ctx context.Context, driver *domain.PrivateDriver) (*domain.PrivateDriver, error)
	GetPrivateDriverByUserID(ctx context.Context, userID uuid.UUID) (*domain.PrivateDriver, error)

	// Matches
	GetMatchByID(ctx context.Context, matchID uuid.UUID) (*domain.DriverParentMatch, error)
	UpdateMatchStatus(ctx context.Context, matchID uuid.UUID, status string) (*domain.DriverParentMatch, error)

	// Private trips
	CreatePrivateTrip(ctx context.Context, trip *domain.PrivateTrip) (*domain.PrivateTrip, error)
	GetActivePrivateTripByDriverID(ctx context.Context, driverID uuid.UUID) (*domain.PrivateTrip, error)
	GetPrivateTripByID(ctx context.Context, id uuid.UUID) (*domain.PrivateTrip, error)
	UpdatePrivateTrip(ctx context.Context, trip *domain.PrivateTrip) (*domain.PrivateTrip, error)

	// Private trip children
	AddPrivateTripChild(ctx context.Context, ptc *domain.PrivateTripChild) (*domain.PrivateTripChild, error)
	GetPrivateTripChildByTripAndChildID(ctx context.Context, tripID, childID uuid.UUID) (*domain.PrivateTripChild, error)
}

type SchoolDriverRepository interface {
	GetSchoolDriverByUserID(ctx context.Context, userID uuid.UUID) (*domain.SchoolDriver, error)
	SearchStudentsBySchoolID(ctx context.Context, schoolID uuid.UUID, query string) ([]*domain.Student, error)
	GetStudentByID(ctx context.Context, id uuid.UUID) (*domain.Student, error)
	CreateTrip(ctx context.Context, trip *domain.Trip) (*domain.Trip, error)
	GetTripByID(ctx context.Context, id uuid.UUID) (*domain.Trip, error)
	GetActiveTripByDriverID(ctx context.Context, driverID uuid.UUID) (*domain.Trip, error)
	UpdateTrip(ctx context.Context, trip *domain.Trip) (*domain.Trip, error)
	AddTripStudent(ctx context.Context, ts *domain.TripStudent) (*domain.TripStudent, error)
	GetTripStudentByTripAndStudentID(ctx context.Context, tripID, studentID uuid.UUID) (*domain.TripStudent, error)
	GetBoardedStudents(ctx context.Context, tripID uuid.UUID) ([]*domain.Student, error)
}

type PrivateParentRepository interface {
	// Profile
	CreatePrivateParent(ctx context.Context, parent *domain.PrivateParent) (*domain.PrivateParent, error)
	GetPrivateParentByUserID(ctx context.Context, userID uuid.UUID) (*domain.PrivateParent, error)
	UpdatePrivateParent(ctx context.Context, parent *domain.PrivateParent) (*domain.PrivateParent, error)
	DeleteAccount(ctx context.Context, userID uuid.UUID) error

	// Children
	CreatePrivateChild(ctx context.Context, child *domain.PrivateChild) (*domain.PrivateChild, error)
	GetPrivateChildrenByParentID(ctx context.Context, parentID uuid.UUID) ([]*domain.PrivateChild, error)
	GetPrivateChildByID(ctx context.Context, id uuid.UUID) (*domain.PrivateChild, error)
	UpdatePrivateChild(ctx context.Context, child *domain.PrivateChild) (*domain.PrivateChild, error)

	// Driver matching
	CreateDriverParentMatch(ctx context.Context, match *domain.DriverParentMatch) (*domain.DriverParentMatch, error)
	GetMatchByParentAndDriverID(ctx context.Context, parentID, driverID uuid.UUID) (*domain.DriverParentMatch, error)

	// School connections
	CreateParentSchoolConnection(ctx context.Context, conn *domain.ParentSchoolConnection) (*domain.ParentSchoolConnection, error)
	GetParentSchoolConnection(ctx context.Context, parentID, schoolID uuid.UUID) (*domain.ParentSchoolConnection, error)

	// School browsing
	GetAllSchools(ctx context.Context) ([]*domain.School, error)
	SearchSchools(ctx context.Context, query string) ([]*domain.School, error)
	FilterSchools(ctx context.Context, address string) ([]*domain.School, error)
	GetSchoolByID(ctx context.Context, id uuid.UUID) (*domain.School, error)

	// Driver lookup (for verifying match targets)
	GetPrivateDriverByID(ctx context.Context, id uuid.UUID) (*domain.PrivateDriver, error)

	// Trips
	GetPrivateTripsByParentID(ctx context.Context, parentID uuid.UUID) ([]*domain.PrivateTrip, error)
	GetPrivateTripByIDForParent(ctx context.Context, parentID, tripID uuid.UUID) (*domain.PrivateTrip, error)
	UpdatePrivateTripStatus(ctx context.Context, tripID uuid.UUID, status string) (*domain.PrivateTrip, error)
}

type GuardianRepository interface {
	GetGuardianByUserID(ctx context.Context, userID uuid.UUID) (*domain.Guardian, error)
	GetStudentsByGuardianID(ctx context.Context, guardianID uuid.UUID) ([]*domain.Student, error)
	GetStudentByIDForGuardian(ctx context.Context, guardianID uuid.UUID, studentID uuid.UUID) (*domain.Student, error)
	UpdateGuardian(ctx context.Context, guardian *domain.Guardian) (*domain.Guardian, error)
	GetActiveSchoolTripForStudent(ctx context.Context, studentID uuid.UUID) (*domain.Trip, error)
}

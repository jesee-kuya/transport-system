package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/jesee-kuya/transport-system/domain"
)

type Authentication interface {
	SignUp(ctx context.Context, req *domain.SignUpRequest) (*domain.SignUpResponse, error)
	Login(ctx context.Context, req *domain.LoginRequest) (*domain.LoginResponse, error)
	ChangePassword(ctx context.Context, userID uuid.UUID, req *domain.ChangePasswordRequest) error
	ForgotPassword(ctx context.Context, req *domain.ForgotPasswordRequest) (*domain.ForgotPasswordResponse, error)
	ResetPassword(ctx context.Context, req *domain.ResetPasswordRequest) error
}

type SchoolService interface {
	AddMySchool(ctx context.Context, adminID uuid.UUID, req *domain.AddSchoolRequest) (*domain.School, error)

	EnrollStudent(ctx context.Context, adminID uuid.UUID, req *domain.EnrollStudentRequest) (*domain.Student, error)
	GetStudents(ctx context.Context, adminID uuid.UUID) ([]*domain.Student, error)
	GetStudent(ctx context.Context, adminID uuid.UUID, studentID uuid.UUID) (*domain.Student, error)
	GetStudentGuardians(ctx context.Context, adminID uuid.UUID, studentID uuid.UUID) ([]*domain.Guardian, error)
	AddStudentGuardian(ctx context.Context, adminID uuid.UUID, studentID uuid.UUID, req *domain.AddStudentGuardianRequest) (*domain.StudentGuardian, error)
	ManageStudent(ctx context.Context, adminID uuid.UUID, studentID uuid.UUID, req *domain.ManageStudentRequest) (*domain.Student, error)
	FilterStudents(ctx context.Context, adminID uuid.UUID, grade string, isActive *bool) ([]*domain.Student, error)
	SearchStudents(ctx context.Context, adminID uuid.UUID, query string) ([]*domain.Student, error)
	RemoveStudent(ctx context.Context, adminID uuid.UUID, studentID uuid.UUID) error
	EditStudent(ctx context.Context, adminID uuid.UUID, studentID uuid.UUID, req *domain.EditStudentRequest) (*domain.Student, error)

	AddBus(ctx context.Context, adminID uuid.UUID, req *domain.AddBusRequest) (*domain.Bus, error)
	GetBuses(ctx context.Context, adminID uuid.UUID) ([]*domain.Bus, error)
	GetBus(ctx context.Context, adminID uuid.UUID, busID uuid.UUID) (*domain.Bus, error)
	EditBus(ctx context.Context, adminID uuid.UUID, busID uuid.UUID, req *domain.EditBusRequest) (*domain.Bus, error)
	FilterBuses(ctx context.Context, adminID uuid.UUID, isActive *bool) ([]*domain.Bus, error)
	SearchBuses(ctx context.Context, adminID uuid.UUID, query string) ([]*domain.Bus, error)
	RemoveBus(ctx context.Context, adminID uuid.UUID, busID uuid.UUID) error
	TrackBus(ctx context.Context, adminID uuid.UUID, busID uuid.UUID) (*domain.Trip, error)

	AddDriver(ctx context.Context, adminID uuid.UUID, req *domain.AddDriverRequest) (*domain.SchoolDriver, error)
	GetDrivers(ctx context.Context, adminID uuid.UUID) ([]*domain.SchoolDriver, error)
	GetDriver(ctx context.Context, adminID uuid.UUID, driverID uuid.UUID) (*domain.SchoolDriver, error)
	EditDriver(ctx context.Context, adminID uuid.UUID, driverID uuid.UUID, req *domain.EditDriverRequest) (*domain.SchoolDriver, error)
	FilterDrivers(ctx context.Context, adminID uuid.UUID, isActive *bool) ([]*domain.SchoolDriver, error)
	SearchDrivers(ctx context.Context, adminID uuid.UUID, query string) ([]*domain.SchoolDriver, error)
	RemoveDriver(ctx context.Context, adminID uuid.UUID, driverID uuid.UUID) error
	ConnectPrivateDriver(ctx context.Context, adminID uuid.UUID, req *domain.ConnectPrivateDriverRequest) error
	GetPrivateDrivers(ctx context.Context, adminID uuid.UUID) ([]*domain.PrivateDriver, error)
	GetPrivateDriver(ctx context.Context, adminID uuid.UUID, driverID uuid.UUID) (*domain.PrivateDriver, error)
	TrackPrivateTrips(ctx context.Context, adminID uuid.UUID) ([]*domain.PrivateTrip, error)
	GetPrivateTripData(ctx context.Context, adminID uuid.UUID, tripID uuid.UUID) (*domain.PrivateTrip, error)
}

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

type PrivateDriverService interface {
	KYCDriver(ctx context.Context, userID uuid.UUID, req *domain.KYCDriverRequest) (*domain.PrivateDriver, error)
	MatchWithParent(ctx context.Context, userID uuid.UUID, matchID uuid.UUID, req *domain.RespondToMatchRequest) (*domain.DriverParentMatch, error)
	OnboardPrivateStudent(ctx context.Context, userID uuid.UUID, tripID uuid.UUID, req *domain.OnboardPrivateStudentRequest) (*domain.PrivateTripChild, error)
	StartTrip(ctx context.Context, userID uuid.UUID, req *domain.StartPrivateTripRequest) (*domain.PrivateTrip, error)
	EndTrip(ctx context.Context, userID uuid.UUID, tripID uuid.UUID) (*domain.PrivateTrip, error)
	UpdatePrivateTripStatus(ctx context.Context, userID uuid.UUID, tripID uuid.UUID, req *domain.UpdatePrivateTripStatusRequest) (*domain.PrivateTrip, error)
}

type SchoolDriverService interface {
	SearchStudents(ctx context.Context, userID uuid.UUID, query string) ([]*domain.Student, error)
	OnboardStudent(ctx context.Context, userID uuid.UUID, tripID uuid.UUID, studentID uuid.UUID) (*domain.TripStudent, error)
	StartTrip(ctx context.Context, userID uuid.UUID, req *domain.StartTripRequest) (*domain.Trip, error)
	EndTrip(ctx context.Context, userID uuid.UUID, tripID uuid.UUID) (*domain.Trip, error)
	UpdateTripStatus(ctx context.Context, userID uuid.UUID, tripID uuid.UUID, req *domain.UpdateTripStatusRequest) (*domain.Trip, error)
	ViewBoardedStudents(ctx context.Context, userID uuid.UUID, tripID uuid.UUID) ([]*domain.Student, error)
}

type PrivateParentService interface {
	CollectKYC(ctx context.Context, userID uuid.UUID, req *domain.CollectKYCRequest) (*domain.PrivateParent, error)
	GetPrivateProfile(ctx context.Context, userID uuid.UUID) (*domain.PrivateParent, error)
	AddMyChild(ctx context.Context, userID uuid.UUID, req *domain.AddPrivateChildRequest) (*domain.PrivateChild, error)
	GetMyChildren(ctx context.Context, userID uuid.UUID) ([]*domain.PrivateChild, error)
	EditMyPrivateProfile(ctx context.Context, userID uuid.UUID, req *domain.EditPrivateProfileRequest) (*domain.PrivateParent, error)
	EditMyChild(ctx context.Context, userID uuid.UUID, childID uuid.UUID, req *domain.EditPrivateChildRequest) (*domain.PrivateChild, error)
	DeleteMyAccount(ctx context.Context, userID uuid.UUID) error
	MatchWithDriver(ctx context.Context, userID uuid.UUID, req *domain.MatchWithDriverRequest) (*domain.DriverParentMatch, error)
	ConnectWithSchool(ctx context.Context, userID uuid.UUID, req *domain.ConnectWithSchoolRequest) (*domain.ParentSchoolConnection, error)
	GetSchools(ctx context.Context) ([]*domain.School, error)
	SearchSchools(ctx context.Context, query string) ([]*domain.School, error)
	FilterSchools(ctx context.Context, address string) ([]*domain.School, error)
	GetSchool(ctx context.Context, schoolID uuid.UUID) (*domain.School, error)
	ReceiveStudent(ctx context.Context, userID uuid.UUID, tripID uuid.UUID) (*domain.PrivateTrip, error)
	ConfirmStudentBoarding(ctx context.Context, userID uuid.UUID, tripID uuid.UUID) (*domain.PrivateTrip, error)
	GetTrips(ctx context.Context, userID uuid.UUID) ([]*domain.PrivateTrip, error)
	TrackTrip(ctx context.Context, userID uuid.UUID, tripID uuid.UUID) (*domain.PrivateTrip, error)
}

type GuardianService interface {
	GetMyStudents(ctx context.Context, userID uuid.UUID) ([]*domain.Student, error)
	GetMyStudent(ctx context.Context, userID uuid.UUID, studentID uuid.UUID) (*domain.Student, error)
	TrackMyStudent(ctx context.Context, userID uuid.UUID, studentID uuid.UUID) (*domain.Trip, error)
	GetMyProfile(ctx context.Context, userID uuid.UUID) (*domain.Guardian, error)
	EditMyProfile(ctx context.Context, userID uuid.UUID, req *domain.EditGuardianProfileRequest) (*domain.Guardian, error)
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

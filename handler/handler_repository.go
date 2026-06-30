package handler

import "github.com/gin-gonic/gin"

type Authentication interface {
	Login(*gin.Context)
	SignUp(*gin.Context)
	ChangePassword(*gin.Context)
	ResetPassword(*gin.Context)
	ForgotPassword(*gin.Context)
}

type School interface {
	AddMySchool(*gin.Context)
	EnrollStudent(*gin.Context)
	GetStudents(*gin.Context)
	GetStudent(*gin.Context)
	GetStudentGuardians(*gin.Context)
	AddStudentGuardian(*gin.Context)
	ManageStudents(*gin.Context)
	FilterStudents(*gin.Context)
	SearchStudents(*gin.Context)
	RemoveStudent(*gin.Context)
	EditStudent(*gin.Context)

	AddBus(*gin.Context)
	GetBuses(*gin.Context)
	GetBus(*gin.Context)
	EditBus(*gin.Context)
	FilterBuses(*gin.Context)
	SearchBuses(*gin.Context)
	RemoveBus(*gin.Context)
	TrackBus(*gin.Context)

	AddDriver(*gin.Context)
	GetDrivers(*gin.Context)
	GetDriver(*gin.Context)
	EditDriver(*gin.Context)
	FilterDrivers(*gin.Context)
	SearchDrivers(*gin.Context)
	RemoveDriver(*gin.Context)
	ConnectPrivateDriver(*gin.Context)
	GetPrivateDrivers(*gin.Context) //only the ones delivering and picking students from the said school
	GetPrivateDriver(*gin.Context)
	TrackPrivateTrips(*gin.Context)
	GetPrivateTripData(*gin.Context)
}

type Guardian interface {
	GetMyStudents(*gin.Context)
	GetMyStudent(*gin.Context)
	TrackMyStudent(*gin.Context)
	GetMyProfile(*gin.Context)
	EditMyProfile(*gin.Context)
}

type PrivateParent interface {
	CollectKYC(*gin.Context)
	GetPrivateProfile(gin.Context)
	AddMyChild(*gin.Context)
	GetMyChildren(*gin.Context)
	EditMyPrivateProfile(*gin.Context)
	EditMyChild(*gin.Context)
	DeleteMyAccount(*gin.Context)
	MatchWithDriver(*gin.Context)
	ConnectWithSchool(*gin.Context)
	GetSchools(*gin.Context)
	SearchSchools(*gin.Context)
	FilterSchools(*gin.Context)
	GetSchool(*gin.Context)
	ReceiveStudent(*gin.Context)
	ConfirmStudentBoarding(*gin.Context)
	GetTrips(*gin.Context)
	TrackTrip(*gin.Context)
}

type SchoolDriver interface {
	SearchStudents(*gin.Context)
	OnboardStudent(*gin.Context)
	StartTrip(*gin.Context)
	EndTrip(*gin.Context)
	UpdateTripStatus(*gin.Context)
	ViewBoardedStudents(*gin.Context)
}

type PrivateDriver interface {
	KYCDriver(*gin.Context)
	MatchWithParent(*gin.Context)
	OnboardPrivateStudent(*gin.Context)
	StartTrip(*gin.Context)
	EndTrip(*gin.Context)
	UpdatePrivateTripStatus(*gin.Context)
}

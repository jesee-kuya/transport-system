package domain

type SignUpRequest struct {
	Username string `json:"username" binding:"required,min=3,max=100"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
	Role     string `json:"role" binding:"required,oneof=school_admin guardian private_parent school_driver private_driver"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type ChangePasswordRequest struct {
	CurrentPassword string `json:"current_password" binding:"required"`
	NewPassword     string `json:"new_password" binding:"required,min=8"`
}

type ForgotPasswordRequest struct {
	Email string `json:"email" binding:"required,email"`
}

type ResetPasswordRequest struct {
	Token       string `json:"token" binding:"required"`
	NewPassword string `json:"new_password" binding:"required,min=8"`
}

type AddSchoolRequest struct {
	Name         string `json:"name" binding:"required"`
	Address      string `json:"address"`
	ContactEmail string `json:"contact_email" binding:"omitempty,email"`
	ContactPhone string `json:"contact_phone"`
}

type EnrollStudentRequest struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Grade     string `json:"grade"`
}

type EditStudentRequest struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Grade     string `json:"grade"`
}

type ManageStudentRequest struct {
	IsActive bool `json:"is_active"`
}

type AddStudentGuardianRequest struct {
	GuardianID   string `json:"guardian_id" binding:"required,uuid"`
	Relationship string `json:"relationship"`
	IsPrimary    bool   `json:"is_primary"`
}

type AddBusRequest struct {
	PlateNumber string `json:"plate_number" binding:"required"`
	Model       string `json:"model"`
	Capacity    int    `json:"capacity" binding:"required,min=1"`
}

type EditBusRequest struct {
	PlateNumber string `json:"plate_number" binding:"required"`
	Model       string `json:"model"`
	Capacity    int    `json:"capacity" binding:"required,min=1"`
	IsActive    bool   `json:"is_active"`
}

type AddDriverRequest struct {
	UserID        string `json:"user_id" binding:"required,uuid"`
	LicenseNumber string `json:"license_number" binding:"required"`
}

type EditDriverRequest struct {
	LicenseNumber string `json:"license_number" binding:"required"`
	IsActive      bool   `json:"is_active"`
}

type ConnectPrivateDriverRequest struct {
	DriverID string `json:"driver_id" binding:"required,uuid"`
}

type EditGuardianProfileRequest struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Phone     string `json:"phone"`
}

type CollectKYCRequest struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Phone     string `json:"phone" binding:"required"`
}

type EditPrivateProfileRequest struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Phone     string `json:"phone" binding:"required"`
}

type AddPrivateChildRequest struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Grade     string `json:"grade"`
}

type EditPrivateChildRequest struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Grade     string `json:"grade"`
}

type MatchWithDriverRequest struct {
	DriverID string `json:"driver_id" binding:"required,uuid"`
}

type ConnectWithSchoolRequest struct {
	SchoolID string `json:"school_id" binding:"required,uuid"`
}

type StartTripRequest struct {
	BusID    string `json:"bus_id" binding:"required,uuid"`
	TripType string `json:"trip_type" binding:"required"`
}

type OnboardStudentRequest struct {
	StudentID string `json:"student_id" binding:"required,uuid"`
}

type UpdateTripStatusRequest struct {
	Status string `json:"status" binding:"required"`
}

type KYCDriverRequest struct {
	FirstName     string `json:"first_name" binding:"required"`
	LastName      string `json:"last_name" binding:"required"`
	Phone         string `json:"phone" binding:"required"`
	LicenseNumber string `json:"license_number" binding:"required"`
}

type RespondToMatchRequest struct {
	Status string `json:"status" binding:"required,oneof=accepted rejected"`
}

type OnboardPrivateStudentRequest struct {
	ChildID string `json:"child_id" binding:"required,uuid"`
}

type StartPrivateTripRequest struct {
	MatchID string `json:"match_id" binding:"required,uuid"`
}

type UpdatePrivateTripStatusRequest struct {
	Status string `json:"status" binding:"required"`
}

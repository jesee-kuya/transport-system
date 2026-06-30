package domain

import (
	"time"

	"github.com/google/uuid"
)

type SchoolDriver struct {
	ID            uuid.UUID `db:"id" json:"id"`
	UserID        uuid.UUID `db:"user_id" json:"user_id"`
	SchoolID      uuid.UUID `db:"school_id" json:"school_id"`
	LicenseNumber string    `db:"license_number" json:"license_number"`
	IsActive      bool      `db:"is_active" json:"is_active"`
	CreatedAt     time.Time `db:"created_at" json:"created_at"`
	UpdatedAt     time.Time `db:"updated_at" json:"updated_at"`
}

type PrivateDriver struct {
	ID            uuid.UUID `db:"id" json:"id"`
	UserID        uuid.UUID `db:"user_id" json:"user_id"`
	FirstName     string    `db:"first_name" json:"first_name"`
	LastName      string    `db:"last_name" json:"last_name"`
	Phone         string    `db:"phone" json:"phone"`
	LicenseNumber string    `db:"license_number" json:"license_number"`
	KYCVerified   bool      `db:"kyc_verified" json:"kyc_verified"`
	CreatedAt     time.Time `db:"created_at" json:"created_at"`
	UpdatedAt     time.Time `db:"updated_at" json:"updated_at"`
}

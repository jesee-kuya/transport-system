package domain

import (
	"time"

	"github.com/google/uuid"
)

type PrivateParent struct {
	ID          uuid.UUID `db:"id" json:"id"`
	UserID      uuid.UUID `db:"user_id" json:"user_id"`
	FirstName   string    `db:"first_name" json:"first_name"`
	LastName    string    `db:"last_name" json:"last_name"`
	Phone       string    `db:"phone" json:"phone"`
	KYCVerified bool      `db:"kyc_verified" json:"kyc_verified"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}

type PrivateChild struct {
	ID        uuid.UUID  `db:"id" json:"id"`
	ParentID  uuid.UUID  `db:"parent_id" json:"parent_id"`
	SchoolID  *uuid.UUID `db:"school_id" json:"school_id"`
	FirstName string     `db:"first_name" json:"first_name"`
	LastName  string     `db:"last_name" json:"last_name"`
	Grade     string     `db:"grade" json:"grade"`
	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt time.Time  `db:"updated_at" json:"updated_at"`
}

type DriverParentMatch struct {
	ID        uuid.UUID `db:"id" json:"id"`
	ParentID  uuid.UUID `db:"parent_id" json:"parent_id"`
	DriverID  uuid.UUID `db:"driver_id" json:"driver_id"`
	Status    string    `db:"status" json:"status"` // pending, accepted, rejected
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

type ParentSchoolConnection struct {
	ID        uuid.UUID `db:"id" json:"id"`
	ParentID  uuid.UUID `db:"parent_id" json:"parent_id"`
	SchoolID  uuid.UUID `db:"school_id" json:"school_id"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

type PrivateTripChild struct {
	ID        uuid.UUID `db:"id" json:"id"`
	TripID    uuid.UUID `db:"trip_id" json:"trip_id"`
	ChildID   uuid.UUID `db:"child_id" json:"child_id"`
	BoardedAt time.Time `db:"boarded_at" json:"boarded_at"`
}

package domain

import (
	"time"

	"github.com/google/uuid"
)

type Student struct {
	ID        uuid.UUID `db:"id" json:"id"`
	SchoolID  uuid.UUID `db:"school_id" json:"school_id"`
	FirstName string    `db:"first_name" json:"first_name"`
	LastName  string    `db:"last_name" json:"last_name"`
	Grade     string    `db:"grade" json:"grade"`
	IsActive  bool      `db:"is_active" json:"is_active"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

type Guardian struct {
	ID        uuid.UUID `db:"id" json:"id"`
	UserID    uuid.UUID `db:"user_id" json:"user_id"`
	FirstName string    `db:"first_name" json:"first_name"`
	LastName  string    `db:"last_name" json:"last_name"`
	Phone     string    `db:"phone" json:"phone"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

type StudentGuardian struct {
	ID           uuid.UUID `db:"id" json:"id"`
	StudentID    uuid.UUID `db:"student_id" json:"student_id"`
	GuardianID   uuid.UUID `db:"guardian_id" json:"guardian_id"`
	Relationship string    `db:"relationship" json:"relationship"`
	IsPrimary    bool      `db:"is_primary" json:"is_primary"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
}

package domain

import (
	"time"

	"github.com/google/uuid"
)

type Bus struct {
	ID          uuid.UUID `db:"id" json:"id"`
	SchoolID    uuid.UUID `db:"school_id" json:"school_id"`
	PlateNumber string    `db:"plate_number" json:"plate_number"`
	Model       string    `db:"model" json:"model"`
	Capacity    int       `db:"capacity" json:"capacity"`
	IsActive    bool      `db:"is_active" json:"is_active"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}

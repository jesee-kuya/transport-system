package domain

import (
	"time"

	"github.com/google/uuid"
)

type Trip struct {
	ID        uuid.UUID  `db:"id" json:"id"`
	SchoolID  uuid.UUID  `db:"school_id" json:"school_id"`
	DriverID  uuid.UUID  `db:"driver_id" json:"driver_id"`
	BusID     uuid.UUID  `db:"bus_id" json:"bus_id"`
	TripType  string     `db:"trip_type" json:"trip_type"`
	Status    string     `db:"status" json:"status"`
	StartedAt *time.Time `db:"started_at" json:"started_at"`
	EndedAt   *time.Time `db:"ended_at" json:"ended_at"`
	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt time.Time  `db:"updated_at" json:"updated_at"`
}

type TripStudent struct {
	ID        uuid.UUID `db:"id" json:"id"`
	TripID    uuid.UUID `db:"trip_id" json:"trip_id"`
	StudentID uuid.UUID `db:"student_id" json:"student_id"`
	BoardedAt time.Time `db:"boarded_at" json:"boarded_at"`
}

type PrivateTrip struct {
	ID        uuid.UUID  `db:"id" json:"id"`
	DriverID  uuid.UUID  `db:"driver_id" json:"driver_id"`
	MatchID   uuid.UUID  `db:"match_id" json:"match_id"`
	Status    string     `db:"status" json:"status"`
	StartedAt *time.Time `db:"started_at" json:"started_at"`
	EndedAt   *time.Time `db:"ended_at" json:"ended_at"`
	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt time.Time  `db:"updated_at" json:"updated_at"`
}

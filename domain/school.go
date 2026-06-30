package domain

import (
	"time"

	"github.com/google/uuid"
)

type School struct {
	ID           uuid.UUID `db:"id" json:"id"`
	AdminID      uuid.UUID `db:"admin_id" json:"admin_id"`
	Name         string    `db:"name" json:"name"`
	Address      string    `db:"address" json:"address"`
	ContactEmail string    `db:"contact_email" json:"contact_email"`
	ContactPhone string    `db:"contact_phone" json:"contact_phone"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
	UpdatedAt    time.Time `db:"updated_at" json:"updated_at"`
}

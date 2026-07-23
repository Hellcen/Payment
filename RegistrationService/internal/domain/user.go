package domain

import (
	"time"

	"github.com/gofrs/uuid/v5"
)

type User struct {
	ID           uuid.UUID
	PublicID     uuid.UUID
	Username     string
	Email        string
	Number       string
	PasswordHash string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

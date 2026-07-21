package models

import (
	"time"

	"github.com/gofrs/uuid/v5"
)

type User struct {
	Id           uuid.UUID
	Username     string
	Number       string
	PasswordHash string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

package port

import (
	"context"
	"payment/RegistrationService/internal/domain"
)

type UserRepository interface {
	Register(ctx context.Context, username, email, number, password string) (*domain.User, error)
}
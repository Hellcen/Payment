package service

import (
	"context"
	"payment/RegistrationService/internal/domain"
	"payment/RegistrationService/internal/port"
	"time"

	"github.com/gofrs/uuid/v5"
)

type AuthService struct {
	User port.DatabaseRepository
}

func NewAuthService(ur port.DatabaseRepository) *AuthService {
	return &AuthService{
		User: ur,
	}
}

func (as *AuthService) Register(ctx context.Context, username, email, number, password string) (*domain.User, error) {
	//Проверка на email
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	publicId, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	user := &domain.User{
		ID:       id,
		PublicID: publicId,
		Username: username,
		Email: email,
		Number: number,
		PasswordHash: password, //TODO: Написать функцию хеширования для пароля
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	return user, nil
}

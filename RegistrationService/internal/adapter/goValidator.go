package adapter

import (
	"payment/RegistrationService/internal/port"

	v "github.com/go-playground/validator/v10"
)

type GoValidator struct {
	validator port.Validator
}

func New() *GoValidator{
	return &GoValidator{validator: v.New(v.WithRequiredStructEnabled())}
}

func (gv *GoValidator) Struct(z any) error {
	return gv.validator.Struct(z)
}
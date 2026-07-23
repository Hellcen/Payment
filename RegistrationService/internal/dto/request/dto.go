package request

type UserRegistrationDTO struct {
	Username string `json:"username" validate:"required,min=1,max=30"`
	Email    string `json:"email" validate:"required,email"`
	Number   string `json:"number" validate:"required,e164"`
	Password string `json:"password" validate:"required,min=10"`
}

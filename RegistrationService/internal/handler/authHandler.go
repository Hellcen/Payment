package handler

import (
	"encoding/json"
	"net/http"
	h "payment/RegistrationService/internal/adapter/http/response"
	"payment/RegistrationService/internal/dto/request"
	"payment/RegistrationService/internal/port"
	"payment/pkg/logger"

	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type AuthHandler struct {
	AuthService port.UserRepository
	Logger      logger.Logger
	Validator   port.Validator
}

func NewAuthHandler(ur port.UserRepository, logger logger.Logger, val port.Validator) *AuthHandler {
	return &AuthHandler{
		AuthService: ur,
		Logger:      logger,
		Validator:   val,
	}
}

func (ah *AuthHandler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var req request.UserRegistrationDTO

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ah.Logger.Zaplogger.Error("JSON not decoder",
			zap.Error(err),
		)
	}

	if err := ah.Validator.Struct(req); err != nil {
		errs := make(map[string]string)
		for _, e := range err.(validator.ValidationErrors) {
			errs[e.Field()] = e.Tag()
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"errors": errs,
		})

		return
	}

	user, err := ah.AuthService.Register(r.Context(), req.Username, req.Email, req.Number, req.Password)
	if err != nil {
		ah.Logger.Zaplogger.Error("The user was unable to register",
			zap.Error(err),
		)
	}

	// TODO: Нужен jwt для id

	res := &h.RegisterResponseDTO{

	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

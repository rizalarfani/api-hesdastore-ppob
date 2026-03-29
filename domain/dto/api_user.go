package dto

type ApiUserRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	ApiKey   string `json:"key" validate:"required"`
}

package dto

type ApiUserRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	ApiKey   string `json:"key" validate:"required"`
}

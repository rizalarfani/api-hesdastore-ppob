package model

type ApiUser struct {
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
	APIKey   string `json:"key" db:"key"`
}

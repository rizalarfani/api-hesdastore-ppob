package model

type ApiUser struct {
	UserID   int    `db:"id"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
	APIKey   string `json:"key" db:"key"`
	Role     int    `db:"level"`
}

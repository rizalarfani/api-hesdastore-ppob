package model

type Account struct {
	Name    string `db:"nama"`
	Balance int    `db:"saldo"`
}

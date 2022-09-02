package models

type Response struct {
	ID   string `db:"id" json:"id"`
	Data any    `db:"data" json:"data"`
}

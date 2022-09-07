package models

type Response struct {
	Data any `db:"data" json:"data"`
}

type User struct {
	ID        int    `json:"id"`
	Email     string `json:"email" binding:"required,email"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Active    int    `json:"active"`
}

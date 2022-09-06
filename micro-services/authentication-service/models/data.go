package models

type Response struct {
	Data any `db:"data" json:"data"`
}

type User struct {
	ID        int    `json:"id"`
	Email     string `json:"emial"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Password  string `json:"password"`
	Active    int    `json:"active"`
}

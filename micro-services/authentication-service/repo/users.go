package repo

import (
	"auth/models"
	"database/sql"
	"fmt"
)

type UserInterface interface {
	GetUserById(id int) (*models.User, error)
	CreateUser(user models.User) (*models.Response, error)
}

type userRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) UserInterface {
	return &userRepo{db}
}

func (r *userRepo) GetUserById(id int) (*models.User, error) {
	stmnt := `select id, email, first_name, last_name, active from users where id = $1 `
	qry, err := r.db.Prepare(stmnt)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	defer qry.Close()

	var user models.User
	qErr := qry.QueryRow(id).Scan(&user.ID, &user.Email, &user.FirstName, &user.LastName, &user.Active)
	if qErr != nil {
		if qErr == sql.ErrNoRows {
			return nil, fmt.Errorf("repo error: user not found")
		}
		return nil, fmt.Errorf("repo error: %s", qErr.Error())
	}
	fmt.Printf("user: %v\n", user)

	return &user, nil
}

func (r userRepo) CreateUser(user models.User) (*models.Response, error) {
	var err error
	stmnt := `insert into users (email, first_name, last_name, password, active)
			values($1, $2, $3, $4) returning ID`

	qry, err := r.db.Prepare(stmnt)
	if err != nil {
		fmt.Errorf(err.Error())
		return nil, err
	}
	defer qry.Close()

	_, err = qry.Exec(user.Email, user.FirstName, user.LastName, user.Password, user.Active)
	if err != nil {
		fmt.Errorf(err.Error())
		return nil, err
	}

	res := models.Response{
		Data: "user added",
	}

	return &res, nil
}

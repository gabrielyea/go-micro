package repo

import (
	"auth/models"
	"database/sql"
	"fmt"
)

type UserInterface interface {
	GetUserById(int) (*models.User, error)
	GetUserByEmail(string) (*models.User, error)
	GetAll() (*[]models.User, error)
	CreateUser(*models.User) (*models.Response, error)
	DeleteUserById(int) (*models.Response, error)
}

type userRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) UserInterface {
	return &userRepo{db}
}

func (r *userRepo) GetUserById(id int) (*models.User, error) {
	var err error
	stmnt := `select id, email, first_name, last_name, role from users where id = $1 `
	qry, err := r.db.Prepare(stmnt)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	defer qry.Close()

	var user models.User
	row := qry.QueryRow(id)
	err = row.Scan(&user.ID, &user.Email, &user.FirstName, &user.LastName, &user.Role)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("id not found: %s", err.Error())
	} else if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepo) GetUserByEmail(email string) (*models.User, error) {
	stmnt := `select id, email, first_name, last_name, password_hash, role from users where email = $1`
	var err error
	qry, err := r.db.Prepare(stmnt)
	if err != nil {
		return nil, err
	}
	defer qry.Close()

	var user models.User
	row := qry.QueryRow(email)
	err = row.Scan(&user.ID, &user.Email, &user.FirstName, &user.LastName, &user.PasswordHash, &user.Role)
	if err == sql.ErrNoRows {
		return nil, err
	} else if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepo) CreateUser(user *models.User) (*models.Response, error) {
	var err error
	stmnt := `insert into users (email, first_name, last_name, password_hash, role)
			values($1, $2, $3, $4, $5)`

	qry, err := r.db.Prepare(stmnt)
	if err != nil {
		return nil, fmt.Errorf("repo error: %s", err.Error())
	}
	defer qry.Close()

	if err != nil {
		return nil, fmt.Errorf("repo error: %s", err.Error())
	}
	_, err = qry.Exec(user.Email, user.FirstName, user.LastName, user.PasswordHash, user.Role)
	if err != nil {
		return nil, fmt.Errorf("repo error: %s", err.Error())
	}

	res := models.Response{
		Data: "user added",
	}

	return &res, nil
}

func (r *userRepo) DeleteUserById(id int) (*models.Response, error) {
	var err error
	stmnt := `delete from users where id = $1`
	qry, err := r.db.Prepare(stmnt)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	defer qry.Close()

	_, err = qry.Exec(id)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	var res models.Response
	res.Data = fmt.Sprintf("user with id: %d removed", id)
	return &res, nil
}

func (r *userRepo) GetAll() (*[]models.User, error) {
	var userList []models.User
	stmnt := `select id, email, first_name, last_name FROM users`

	rows, err := r.db.Query(stmnt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var usr models.User
		err = rows.Scan(&usr.ID, &usr.Email, &usr.FirstName, &usr.LastName)
		if err != nil {
			return nil, err
		}
		userList = append(userList, usr)
	}

	fmt.Printf("userList: %v\n", userList)
	return &userList, nil
}

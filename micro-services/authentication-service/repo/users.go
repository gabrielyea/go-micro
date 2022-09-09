package repo

import (
	"auth/models"
	"database/sql"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type UserInterface interface {
	GetUserById(int) (*models.User, error)
	GetUserByEmail(string) (*models.User, error)
	CreateUser(*models.User) (*models.Response, error)
	DeleteUserById(int) (*models.Response, error)
	CheckPasswordHash(string, string) bool
}

type userRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) UserInterface {
	return &userRepo{db}
}

func (r *userRepo) GetUserById(id int) (*models.User, error) {
	var err error
	stmnt := `select id, email, first_name, last_name, active from users where id = $1 `
	qry, err := r.db.Prepare(stmnt)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	defer qry.Close()

	var user models.User
	row := qry.QueryRow(id)
	err = row.Scan(&user.ID, &user.Email, &user.FirstName, &user.LastName, &user.Active, &user.Password)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("id not found: %s", err.Error())
	} else if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepo) GetUserByEmail(email string) (*models.User, error) {
	stmnt := `select id, email, first_name, last_name, active, password from users where email = $1`
	var err error
	qry, err := r.db.Prepare(stmnt)
	if err != nil {
		fmt.Printf("HELLO email: %v\n", email)
		return nil, fmt.Errorf(err.Error())
	}
	defer qry.Close()

	var user models.User
	row := qry.QueryRow(email)
	err = row.Scan(&user.ID, &user.Email, &user.FirstName, &user.LastName, &user.Active, &user.Password)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("email not found: %s", err.Error())
	} else if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepo) CreateUser(user *models.User) (*models.Response, error) {
	var err error
	stmnt := `insert into users (email, first_name, last_name, password, active)
			values($1, $2, $3, $4, $5)`

	qry, err := r.db.Prepare(stmnt)
	if err != nil {
		return nil, fmt.Errorf("repo error: %s", err.Error())
	}
	defer qry.Close()

	hash, err := hashPassowrd(user.Password)
	if err != nil {
		return nil, fmt.Errorf("repo error: %s", err.Error())
	}
	_, err = qry.Exec(user.Email, user.FirstName, user.LastName, hash, user.Active)
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

func (r *userRepo) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func hashPassowrd(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", fmt.Errorf(err.Error())
	}
	return string(bytes), err
}

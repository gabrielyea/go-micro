package services

import (
	"auth/models"
	"auth/repo"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type UserServiceInterface interface {
	GetUserById(id int) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	GetAll() (*[]models.User, error)
	DeleteUserById(id int) (*models.Response, error)
	CreateUser(*models.User) (*models.Response, error)
	Authenticate(string, string) (*models.Response, error)
}

type userService struct {
	r repo.UserInterface
}

func NewUserService(repo repo.UserInterface) UserServiceInterface {
	return &userService{repo}
}

func (s *userService) GetUserById(id int) (*models.User, error) {
	user, err := s.r.GetUserById(id)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	return user, nil
}

func (s *userService) DeleteUserById(id int) (*models.Response, error) {
	res, err := s.r.DeleteUserById(id)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	return res, nil
}

func (s *userService) CreateUser(user *models.User) (*models.Response, error) {
	res, err := s.r.CreateUser(user)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	return res, nil
}

func (s *userService) GetUserByEmail(email string) (*models.User, error) {
	usr, err := s.r.GetUserByEmail(email)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	return usr, nil
}

func (s *userService) Authenticate(email, password string) (*models.Response, error) {
	usr, err := s.r.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	if checkPasswordHash(password, usr.PasswordHash) {
		var res models.Response
		res.Data = "valid user"
		return &res, nil
	}

	credErr := errors.New("invalid credentials")
	return nil, credErr
}

func (s *userService) GetAll() (*[]models.User, error) {
	res, err := s.r.GetAll()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

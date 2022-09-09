package services

import (
	"auth/models"
	"auth/repo"
	"errors"
	"fmt"
)

type UserServiceInterface interface {
	GetUserById(id int) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
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
		return nil, fmt.Errorf("invalid credentials: %s", err.Error())
	}

	if s.r.CheckPasswordHash(password, usr.Password) {
		var res models.Response
		res.Data = "valid user"
		return &res, nil
	}

	credErr := errors.New("invalid credentials")
	return nil, credErr
}

package services

import (
	"auth/models"
	"auth/repo"
	"fmt"
)

type UserServiceInterface interface {
	GetUserById(id int) (*models.User, error)
	DeleteUserById(id int) (*models.Response, error)
	CreateUser(*models.User) (*models.Response, error)
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

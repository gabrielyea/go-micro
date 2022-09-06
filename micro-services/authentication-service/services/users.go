package services

import (
	"auth/models"
	"auth/repo"
	"fmt"
)

type UserServiceInterface interface {
	GetUserById(id int) (*models.User, error)
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

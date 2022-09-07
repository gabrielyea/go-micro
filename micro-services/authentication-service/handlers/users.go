package handlers

import (
	"auth/models"
	"auth/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandlerInterface interface {
	GetUserById(c *gin.Context)
	DeleteUserById(c *gin.Context)
	CreateUser(c *gin.Context)
}

type userHandler struct {
	s services.UserServiceInterface
}

func NewUserHandler(service services.UserServiceInterface) UserHandlerInterface {
	return &userHandler{service}
}

func (h *userHandler) GetUserById(c *gin.Context) {
	id := c.Param("id")
	intId, _ := strconv.Atoi(id)
	user, err := h.s.GetUserById(intId)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *userHandler) DeleteUserById(c *gin.Context) {
	id := c.Param("id")
	intId, _ := strconv.Atoi(id)
	res, err := h.s.DeleteUserById(intId)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *userHandler) CreateUser(c *gin.Context) {
	body := models.User{}
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	res, err := h.s.CreateUser(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, res)

}

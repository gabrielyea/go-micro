package handlers

import (
	"auth/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandlerInterface interface {
	GetUserById(c *gin.Context)
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

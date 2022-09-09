package handlers

import (
	"auth/models"
	"auth/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type AuthenticationHandlerInterface interface {
	GetUserById(c *gin.Context)
	DeleteUserById(c *gin.Context)
	CreateUser(c *gin.Context)
	Authenticate(c *gin.Context)
	Test(c *gin.Context)
}

type authenticationHandler struct {
	s services.UserServiceInterface
}

func NewAuthHandler(service services.UserServiceInterface) AuthenticationHandlerInterface {
	binding.EnableDecoderDisallowUnknownFields = true //jsons should be exactly the same as binding struct
	return &authenticationHandler{service}
}

func (h *authenticationHandler) GetUserById(c *gin.Context) {
	id := c.Param("id")
	intId, _ := strconv.Atoi(id)
	user, err := h.s.GetUserById(intId)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *authenticationHandler) DeleteUserById(c *gin.Context) {
	id := c.Param("id")
	intId, _ := strconv.Atoi(id)
	res, err := h.s.DeleteUserById(intId)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *authenticationHandler) CreateUser(c *gin.Context) {
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

func (h *authenticationHandler) Authenticate(c *gin.Context) {
	login := models.Login{}
	if err := c.BindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err.Error(),
		})
		return
	}

	_, err := h.s.Authenticate(login.Email, login.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": err.Error(),
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"status": "valid user",
	})
}

func (h *authenticationHandler) Test(c *gin.Context) {
	c.JSON(http.StatusAccepted, gin.H{
		"message": "just a test",
	})
}

package handlers

import (
	"auth/jwt"
	"auth/models"
	"auth/services"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"golang.org/x/crypto/bcrypt"
)

type AuthenticationHandlerInterface interface {
	GetUserById(c *gin.Context)
	Index(*gin.Context)
	DeleteUserById(c *gin.Context)
	SignUp(c *gin.Context)
	SignIn(c *gin.Context)
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

func (h *authenticationHandler) SignUp(c *gin.Context) {
	var userReq models.User

	err := c.BindJSON(&userReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	// Hash the password
	hashPass, err := hashPassowrd(userReq.PasswordHash)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	userReq.PasswordHash = hashPass
	res, err := h.s.CreateUser(&userReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Valid credentials!",
		"payload": res,
	})
}

func (h *authenticationHandler) SignIn(c *gin.Context) {
	var loginData models.Login
	err := c.BindJSON(&loginData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	_, err = h.s.Authenticate(loginData.Email, loginData.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})
		return
	}

	tokenString, err := getSignedToken()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusAccepted, tokenString)
}

func (h *authenticationHandler) Index(c *gin.Context) {
	res, err := h.s.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, res)
}

func hashPassowrd(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", fmt.Errorf(err.Error())
	}
	return string(bytes), err
}

func getSignedToken() (string, error) {
	// we make a JWT Token here with signing method of ES256 and claims.
	// claims are attributes.
	// aud - audience
	// iss - issuer
	// exp - expiration of the Token
	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	// 	"aud": "frontend.knowsearch.ml",
	// 	"iss": "knowsearch.ml",
	// 	"exp": string(time.Now().Add(time.Minute * 1).Unix()),
	// })
	claimsMap := map[string]string{
		// "aud": "audience.example",
		// "iss": "issuer.example,
		"exp": fmt.Sprint(time.Now().Add(time.Second * 30).Unix()),
	}
	// here we provide the shared secret. It should be very complex.\
	// Aslo, it should be passed as a System Environment variable

	secret := "V3ry_S3cr3t_C0D3"
	header := "HS256"
	tokenString, err := jwt.GenerateToken(header, claimsMap, secret)
	if err != nil {
		return tokenString, err
	}
	return tokenString, nil
}

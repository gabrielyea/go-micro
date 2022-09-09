package handlers

import (
	"broker/models"
	"broker/services"
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type BrokerHandlerInterface interface {
	SubmissionHandler(*gin.Context)
	Authenticate(models.AuthPayload, *gin.Context)
	Test(*gin.Context)
}

type brokerHandler struct {
	s services.BrokerServiceInterface
}

func NewBrokerHandler(s services.BrokerServiceInterface) BrokerHandlerInterface {
	binding.EnableDecoderDisallowUnknownFields = true
	return &brokerHandler{s}
}

func (h *brokerHandler) SubmissionHandler(c *gin.Context) {
	request := models.RequestPayload{}

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	switch request.Action {
	case "auth":
		h.Authenticate(request.Auth, c)
	}
}

func (h *brokerHandler) Authenticate(req models.AuthPayload, c *gin.Context) {
	jsonData, _ := json.Marshal(req)

	request, err := http.NewRequest("GET", "http://auth-service/v1/authenticate", bytes.NewBuffer(jsonData))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	client := http.Client{}
	res, err := client.Do(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if res.StatusCode != 202 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": res.Status,
		})
	}
}

func (h *brokerHandler) Test(c *gin.Context) {
	c.JSON(http.StatusAccepted, gin.H{
		"message": "just a test",
	})
}

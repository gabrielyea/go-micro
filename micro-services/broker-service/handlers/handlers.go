package handlers

import (
	"broker/models"
	"broker/services"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type BrokerHandlerInterface interface {
	SubmissionHandler(*gin.Context)
	Authenticate(models.AuthPayload, *gin.Context)
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
	case "log":
		h.Log(request.Log)
	case "mail":
		h.Mail(request.Mail)
	}
}

func (h *brokerHandler) Authenticate(req models.AuthPayload, c *gin.Context) {
	jsonData, _ := json.Marshal(req)
	var logData models.LogEntry

	request, err := http.NewRequest("POST", "http://auth-service/v1/auth", bytes.NewBuffer(jsonData))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	client := http.Client{}
	res, err := client.Do(request)
	if err != nil {
		logData.Name = "authorization"
		logData.Data = fmt.Sprintf("failed authorization attempt: %d", res.StatusCode)
		h.Log(logData)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if res.StatusCode != 202 {
		logData.Name = "authorization"
		logData.Data = fmt.Sprintf("failed authorization attempt: %d", res.StatusCode)
		h.Log(logData)
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": res.Status,
		})
		return
	}

	logData.Name = "authorization"
	logData.Data = fmt.Sprintf("succesfull authorization: %d", res.StatusCode)
	h.Log(logData)
	c.JSON(http.StatusAccepted, gin.H{
		"message": "logged in!",
	})
}

func (h *brokerHandler) Log(req models.LogEntry) {
	jsonData, _ := json.Marshal(req)
	request, err := http.NewRequest("POST", "http://logger-service/v1/logs/new", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("err.Error(): %v\n", err.Error())
		return
	}

	client := http.Client{}
	res, err := client.Do(request)
	if err != nil {
		fmt.Printf("err.Error(): %v\n", err.Error())
		return
	}

	fmt.Printf("res: %v\n", res.StatusCode)
}

func (h *brokerHandler) Mail(payload models.MailPayload) {
	jsonData, _ := json.Marshal(payload)
	request, err := http.NewRequest("POST", "http://mail-service/v1/mail", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("err.Error(): %v\n", err.Error())
		return
	}

	client := http.Client{}
	res, err := client.Do(request)
	if err != nil {
		fmt.Printf("err: %v\n", err.Error())
		return
	}

	fmt.Printf("res.StatusCode: %v\n", res.StatusCode)
}

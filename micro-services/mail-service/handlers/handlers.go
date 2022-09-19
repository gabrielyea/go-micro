package handlers

import (
	"mail-service/mailer"
	"mail-service/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MailHandlerInterface interface {
	SendMail(*gin.Context)
	GoMail(*gin.Context)
}

type mailHandler struct {
	mailer mailer.MailerInterface
}

func NewMailHandler(m mailer.MailerInterface) MailHandlerInterface {
	return &mailHandler{m}
}

func (h *mailHandler) SendMail(c *gin.Context) {
	var req models.MailMessage

	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	msg := models.Message{
		From:    req.From,
		To:      req.To,
		Subject: req.Subject,
		Data:    req.Message,
	}

	err = h.mailer.SendSMTPMessage(msg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"message": "Sent to: " + req.To,
	})
}

func (h *mailHandler) GoMail(c *gin.Context) {
	var req models.MailMessage

	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}

	msg := models.Message{
		From:    req.From,
		To:      req.To,
		Subject: req.Subject,
		Data:    req.Message,
	}

	err = h.mailer.SendGoMail(msg)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "mail sent to: " + msg.To,
	})
}

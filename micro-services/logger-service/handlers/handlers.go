package handlers

import (
	"fmt"
	"logger-service/models"
	"logger-service/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoggerHandInterface interface {
	Insert(*gin.Context)
	All(*gin.Context)
	GetById(*gin.Context)
	DropCollection(*gin.Context)
	UpdateCollection(*gin.Context)
}

type loggerHandler struct {
	s services.LoggerServInterface
}

func NewLoggerHandlers(s services.LoggerServInterface) LoggerHandInterface {
	return &loggerHandler{s}
}

func (h *loggerHandler) Insert(c *gin.Context) {
	var body models.LogEntry

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := h.s.Insert(body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	res := fmt.Sprintf("log added: %s", body.Name)
	c.JSON(http.StatusAccepted, gin.H{
		"message": res,
	})
}

func (h *loggerHandler) All(c *gin.Context) {
	res, err := h.s.All()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusAccepted, res)
}

func (h *loggerHandler) GetById(c *gin.Context) {
	logId := c.Param("id")

	fmt.Printf("logId: %v\n", logId)
	res, err := h.s.GetById(logId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	if res != nil {
		c.JSON(http.StatusOK, res)
	}
}

func (h *loggerHandler) DropCollection(c *gin.Context) {
	if err := h.s.DropCollection(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"message": "collection dropped",
	})
}

func (h *loggerHandler) UpdateCollection(c *gin.Context) {
	var body models.LogEntry
	err := c.BindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	res, err := h.s.UpdateCollection(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusAccepted, res)
}

package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type MyResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func GetData(c *gin.Context) {
	var res MyResponse
	res.Message = "hello"
	c.JSON(http.StatusOK, res)
}

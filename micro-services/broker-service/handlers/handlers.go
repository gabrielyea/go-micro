package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type MyResponse struct {
	Msg string `json:"msg"`
}

func GetData(c *gin.Context) {
	var res MyResponse
	res.Msg = "hello"
	c.JSON(http.StatusOK, res)
}

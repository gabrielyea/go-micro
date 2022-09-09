package routes

import (
	"broker/handlers"

	"github.com/gin-gonic/gin"
)

func Public(g *gin.Engine, h handlers.BrokerHandlerInterface) *gin.RouterGroup {

	v1 := g.Group("/v1")
	{
		v1.POST("/auth", h.SubmissionHandler)
		v1.GET("/", h.Test)
	}
	return v1
}

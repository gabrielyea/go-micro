package routes

import (
	"broker/handlers"

	"github.com/gin-gonic/gin"
)

func SetRoutes(g *gin.Engine, h handlers.BrokerHandlerInterface) *gin.RouterGroup {

	v1 := g.Group("/v1")
	publicRoutes(v1, h)
	return v1
}

func publicRoutes(g *gin.RouterGroup, h handlers.BrokerHandlerInterface) {
	g.Group("/")
	{
		g.POST("/auth", h.SubmissionHandler)
		g.POST("/log-grpc", h.LogWithGRPC)
	}
}

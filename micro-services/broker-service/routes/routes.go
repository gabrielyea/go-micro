package routes

import (
	"broker/handlers"

	"github.com/gin-gonic/gin"
)

func SetRoutes(g *gin.Engine, h handlers.BrokerHandlerInterface) *gin.RouterGroup {

	v1 := g.Group("/v1")
	{
		public(g, h)
	}
	return v1
}

func public(g *gin.Engine, h handlers.BrokerHandlerInterface) *gin.RouterGroup {
	pbRoutes := g.Group("/public")
	{
		pbRoutes.POST("/auth", h.SubmissionHandler)
	}
	return pbRoutes
}

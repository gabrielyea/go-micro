package routes

import (
	"auth/handlers"

	"github.com/gin-gonic/gin"
)

func SetRoutes(g *gin.Engine, h handlers.AuthenticationHandlerInterface) *gin.RouterGroup {

	v1 := g.Group("/v1")
	publicRoutes(v1, h)
	return v1
}

func publicRoutes(g *gin.RouterGroup, h handlers.AuthenticationHandlerInterface) {
	g.Group("/")
	{
		g.POST("/auth", h.Authenticate)
	}
}

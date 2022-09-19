package routes

import (
	"mail-service/handlers"

	"github.com/gin-gonic/gin"
)

func SetRoutes(g *gin.Engine, h handlers.MailHandlerInterface) *gin.RouterGroup {
	v1 := g.Group("/v1")
	publicRoutes(v1, h)
	return v1
}

func publicRoutes(g *gin.RouterGroup, h handlers.MailHandlerInterface) {
	g.Group("/")
	{
		g.POST("/mail", h.GoMail)
	}
}

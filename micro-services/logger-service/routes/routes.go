package routes

import (
	"logger-service/handlers"

	"github.com/gin-gonic/gin"
)

func SetRoutes(g *gin.Engine, h handlers.LoggerHandInterface) *gin.RouterGroup {
	v1 := g.Group("/v1")
	publicRoutes(v1, h)
	return v1
}

func publicRoutes(g *gin.RouterGroup, h handlers.LoggerHandInterface) {
	g.Group("/")
	{
		g.GET("/index", h.All)
		g.GET("/logs/:id", h.GetById)
		g.POST("/logs/new", h.Insert)
		g.DELETE("/logs/delete/all", h.DropCollection)
	}
}

package routes

import (
	"logger-service/handlers"

	"github.com/gin-gonic/gin"
)

func SetRoutes(g *gin.Engine, h handlers.LoggerHandInterface) *gin.RouterGroup {
	v1 := g.Group("/v1")
	{
		public(g, h)
		private(g)
	}
	return v1
}

func public(g *gin.Engine, h handlers.LoggerHandInterface) *gin.RouterGroup {
	pbGroup := g.Group("/public")
	{
		pbGroup.GET("/index", h.All)
		pbGroup.GET("/logs/:id", h.GetById)
		pbGroup.POST("/logs/new", h.Insert)
		pbGroup.DELETE("/logs/delete/all", h.DropCollection)
	}
	return pbGroup
}

func private(g *gin.Engine) *gin.RouterGroup {
	pvGroup := g.Group("/private")
	{
		pvGroup.GET("/test1", nil)
	}
	return pvGroup
}

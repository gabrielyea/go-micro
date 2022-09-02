package routes

import (
	"broker/handlers"

	"github.com/gin-gonic/gin"
)

func Public(g *gin.RouterGroup) {
	g.GET("/", handlers.GetData)
}

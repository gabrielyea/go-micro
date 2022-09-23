package routes

import (
	"auth/handlers"
	"auth/middelware"

	"github.com/gin-gonic/gin"
)

func SetRoutes(g *gin.Engine, h handlers.AuthenticationHandlerInterface) *gin.RouterGroup {

	v1 := g.Group("/v1")
	{
		publicGroup := v1.Group("/")
		{
			publicRoutes(publicGroup, h)
		}
		privateGroup := v1.Group("/admin", middelware.TokenValidation())
		{
			privateRoutes(privateGroup, h)
		}
	}
	return v1
}

func publicRoutes(g *gin.RouterGroup, h handlers.AuthenticationHandlerInterface) {
	{
		g.POST("/sign-up", h.SignUp)
		g.POST("/sign-in", h.SignIn)
	}
}
func privateRoutes(g *gin.RouterGroup, h handlers.AuthenticationHandlerInterface) {
	{
		g.GET("/users", h.Index)
	}
}

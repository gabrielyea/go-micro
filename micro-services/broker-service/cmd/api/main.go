package main

import (
	"broker/handlers"
	"broker/routes"
	"broker/services"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.Use(routes.CorsConfig())
	// router.Use(cors.Default())

	s := services.NewBrokerService()
	h := handlers.NewBrokerHandler(s)

	router.RouterGroup = *routes.SetRoutes(router, h)

	router.Run(":80")
}

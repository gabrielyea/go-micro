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

	s := services.NewBrokerService()
	h := handlers.NewBrokerHandler(s)

	routes.SetRoutes(router, h)

	router.Run(":80")
}

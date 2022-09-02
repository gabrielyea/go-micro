package main

import (
	"broker/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.Use(routes.CorsConfig())

	v1 := router.Group("/v1")
	{
		routes.Public(v1.Group("/"))
	}

	router.Run(":8081")
}

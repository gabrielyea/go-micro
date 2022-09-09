package main

import (
	"auth/db"
	"auth/handlers"
	"auth/repo"
	"auth/routes"
	"auth/services"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := db.ConnectDB()
	if err != nil {
		fmt.Errorf("something went wrong with db %s", err)
		return
	}
	defer db.Close()

	router := gin.Default()
	router.Use(routes.CorsConfig())
	// router.Use(cors.Default())

	repo := repo.NewUserRepo(db)
	service := services.NewUserService(repo)
	handlers := handlers.NewAuthHandler(service)

	router.POST("/v1/authenticate", handlers.Authenticate)
	router.GET("/v1/", handlers.Test)

	router.Run(":80")
}

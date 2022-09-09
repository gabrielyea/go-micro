package main

import (
	"auth/db"
	"auth/handlers"
	"auth/repo"
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

	repo := repo.NewUserRepo(db)
	service := services.NewUserService(repo)
	handlers := handlers.NewAuthHandler(service)

	router.GET("/v1/user/:id", handlers.GetUserById)
	router.DELETE("/v1/user/:id", handlers.DeleteUserById)
	router.POST("/v1/user/new", handlers.CreateUser)
	router.GET("/v1/authenticate", handlers.Authenticate)

	router.Run(":8082")
}

package main

import (
	"auth/db"
	"auth/handlers"
	"auth/middelware"
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
	router.GET("/testing", middelware.TokenValidation())
	router.Use(routes.CorsConfig())

	repo := repo.NewUserRepo(db)
	service := services.NewUserService(repo)
	handlers := handlers.NewAuthHandler(service)

	routes.SetRoutes(router, handlers)

	router.Run(":80")
}

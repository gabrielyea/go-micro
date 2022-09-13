package main

import (
	"context"
	"fmt"
	"logger-service/config"
	"logger-service/db"
	"logger-service/handlers"
	"logger-service/repo"
	"logger-service/routes"
	"logger-service/services"

	"github.com/gin-gonic/gin"
)

func main() {
	newConf, err := config.NewConfig("config.yml")
	if err != nil {
		fmt.Printf("err: %v\n", err.Error())
		return
	}

	db, err := db.ConnectDB(*newConf)
	if err != nil {
		fmt.Printf("err: %v\n", err.Error())
		return
	}
	defer db.Disconnect(context.TODO())

	router := gin.Default()

	r := repo.NewLoggerRepo(db)
	s := services.NewLoggerServices(r)
	h := handlers.NewLoggerHandlers(s)

	router.RouterGroup = *routes.SetRoutes(router, h)

	router.Run(newConf.InternalPort)

}

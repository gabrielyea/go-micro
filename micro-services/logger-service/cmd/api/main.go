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
	"net/rpc"

	"github.com/gin-gonic/gin"
)

func main() {
	newConf, err := config.NewConfig("config.yml")
	if err != nil {
		fmt.Printf("err: %v\n", err.Error())
		return
	}
	// db := mongo.Client{}
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

	rpcServer := services.NewRpcServer(r)
	gRpcServer := services.NewGrpcServer(r)

	err = rpc.Register(rpcServer)

	if err != nil {
		fmt.Printf("err: %v\n", err.Error())
		return
	}

	go rpcServer.Listen(*newConf)
	go gRpcServer.Listen(*newConf)

	routes.SetRoutes(router, h)
	router.Run(newConf.InternalPort)

}

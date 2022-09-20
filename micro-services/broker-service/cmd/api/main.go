package main

import (
	"broker/handlers"
	"broker/listener"
	"broker/routes"
	"broker/services"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.Use(routes.CorsConfig())

	rConn, err := listener.NewRabbitConn("amqp://guest:guest@rabbitmq")
	if err != nil {
		fmt.Printf("err: %v\n", err.Error())
		return
	}
	ch, q, err := listener.SetRabbitConf(rConn)
	if err != nil {
		fmt.Printf("err: %v\n", err.Error())
		return
	}

	rabbit := listener.NewRabbit(rConn, ch, q)

	s := services.NewBrokerService(rabbit)
	h := handlers.NewBrokerHandler(s)

	routes.SetRoutes(router, h)

	router.Run(":80")
}

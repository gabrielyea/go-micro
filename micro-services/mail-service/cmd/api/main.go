package main

import (
	"mail-service/config"
	"mail-service/handlers"
	"mail-service/mailer"
	"mail-service/models"
	"mail-service/routes"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	conf, _ := config.NewConfig("config.yml")
	router := gin.Default()
	router.Use(config.CorsConfig())

	port, _ := strconv.Atoi(conf.MailPort)
	mail := models.Mail{
		Domain:      conf.MailDomain,
		Host:        conf.MailHost,
		Port:        port,
		Username:    conf.UserName,
		Password:    conf.Password,
		Encryption:  conf.Encryption,
		FromName:    conf.FromName,
		FromAddress: conf.FromAddress,
	}

	s := mailer.NewMailer(mail)
	h := handlers.NewMailHandler(s)

	routes.SetRoutes(router, h)

	router.Run(conf.WebPort)
}

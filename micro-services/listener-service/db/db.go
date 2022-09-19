package db

import (
	"fmt"
	"listener/config"

	amqp "github.com/rabbitmq/amqp091-go"
)

var db *amqp.Connection

func NewRabbitConn(conf config.Config) (*amqp.Connection, error) {
	var err error
	db, err = amqp.Dial(conf.RemoteURL)
	if err != nil {
		fmt.Printf("err: %v\n", err.Error())
		return nil, err
	}
	return db, nil
}

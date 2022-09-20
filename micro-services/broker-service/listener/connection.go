package listener

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

var db *amqp.Connection

func NewRabbitConn(conString string) (*amqp.Connection, error) {
	var err error
	db, err = amqp.Dial(conString)
	if err != nil {
		fmt.Printf("err: %v\n", err.Error())
		return nil, err
	}
	return db, nil
}

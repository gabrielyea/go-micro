package db

import (
	"fmt"
	"listener/config"
	"log"
	"math"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

var db *amqp.Connection

func NewRabbitConn(conf config.Config) (*amqp.Connection, error) {
	var err error
	db, err = expBackoffConnection(conf.RemoteURL)
	if err != nil {
		fmt.Printf("err: %v\n", err.Error())
		return nil, err
	}
	return db, nil
}

func expBackoffConnection(conStr string) (*amqp.Connection, error) {
	var counts int64
	var backOff = 1 * time.Second
	// var err error

	// don't continue until rabbit is ready
	for {
		c, err := amqp.Dial(conStr)
		if err != nil {
			fmt.Println("RabbitMQ not yet ready...")
			counts++
		} else {
			log.Println("Connected to RabbitMQ!")
			db = c
			break
		}

		if counts > 5 {
			fmt.Println(err)
			return nil, err
		}

		backOff = time.Duration(math.Pow(float64(counts), 2)) * time.Second
		log.Println("backing off...")
		time.Sleep(backOff)
		continue
	}

	return db, nil
}

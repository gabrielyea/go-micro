package listener

import (
	"broker/models"
	"context"
	"encoding/json"
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitInterface interface {
	Push(string, string, string, context.Context) error
}

type rabbit struct {
	db *amqp.Connection
	ch *amqp.Channel
	q  *amqp.Queue
}

func NewRabbit(db *amqp.Connection, ch *amqp.Channel, q *amqp.Queue) RabbitInterface {
	return &rabbit{db, ch, q}
}

func (r *rabbit) Push(name string, data string, key string, ctx context.Context) error {
	var err error
	fmt.Printf("name %s data %s key %s \n", name, data, key)
	err = r.ch.QueueBind(
		r.q.Name,
		key,
		"logs_topic",
		false,
		nil,
	)
	if err != nil {
		return err
	}

	var event models.MqPayload
	event.Name = name
	event.Data = data

	jsonData, _ := json.Marshal(&event)

	err = r.ch.PublishWithContext(ctx,
		"logs_topic",
		key,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        jsonData,
		})
	if err != nil {
		return err
	}

	log.Printf("Sent %s\n", name)
	return nil
}

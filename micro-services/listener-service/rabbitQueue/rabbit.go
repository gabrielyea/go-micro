package rabbitQueue

import (
	"context"
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitInterface interface {
	Listen([]string) error
	Send(string, string, context.Context)
}

type rabbit struct {
	db *amqp.Connection
	ch *amqp.Channel
	q  *amqp.Queue
}

func NewRabbit(db *amqp.Connection, ch *amqp.Channel, q *amqp.Queue) RabbitInterface {
	return &rabbit{db, ch, q}
}

func (r *rabbit) Listen(topics []string) error {
	for _, topic := range topics {
		r.ch.QueueBind(
			r.q.Name,
			topic,
			"logs_topic",
			false,
			nil,
		)
	}

	msgs, err := r.ch.Consume(
		r.q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	var forever chan struct{}

	go func() {
		for d := range msgs {
			fmt.Printf("Got a message: %s\n", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
	return nil
}

func (r *rabbit) Send(msg string, key string, ctx context.Context) {
	var err error
	err = r.ch.QueueBind(
		r.q.Name,
		key,
		"logs_topic",
		false,
		nil,
	)
	if err != nil {
		fmt.Printf("err: %v\n", err.Error())
		return
	}

	err = r.ch.PublishWithContext(ctx,
		"",
		r.q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		})
	if err != nil {
		fmt.Printf("err: %v\n", err.Error())
		return
	}
	log.Printf("Sent %s\n", msg)
}

package event

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

func declarExchange(ch *amqp.Channel) error {
	err := ch.ExchangeDeclare(
		"logs_topic",
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}
	return nil
}

func declarRandomQueue(ch *amqp.Channel) (amqp.Queue, error) {
	return ch.QueueDeclare(
		"",
		false,
		false,
		false,
		false,
		nil,
	)
}

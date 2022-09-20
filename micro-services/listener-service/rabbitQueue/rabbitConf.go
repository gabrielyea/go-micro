package rabbitQueue

import "github.com/rabbitmq/amqp091-go"

func SetRabbitConf(conn *amqp091.Connection) (*amqp091.Channel, *amqp091.Queue, error) {
	c, err := conn.Channel()
	if err != nil {
		return nil, nil, err
	}

	err = c.ExchangeDeclare(
		"logs_topic",
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, nil, err
	}

	q, err := c.QueueDeclare(
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, nil, err
	}

	return c, &q, nil
}

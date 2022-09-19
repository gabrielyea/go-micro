package event

import (
	"bytes"
	"encoding/json"
	"fmt"
	"listener/models"
	"log"
	"net/http"

	amqp "github.com/rabbitmq/amqp091-go"
)

type consumer struct {
	conn      *amqp.Connection
	queueName string
}

func NewConsumer(conn *amqp.Connection) (*consumer, error) {
	consumer := consumer{
		conn: conn,
	}

	err := consumer.setup()
	if err != nil {
		return nil, err
	}
	return &consumer, nil
}

func (cmr *consumer) Listen(topics []string) error {
	ch, err := cmr.conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	q, err := declarRandomQueue(ch)
	if err != nil {
		return err
	}

	for _, topic := range topics {
		err := ch.QueueBind(
			q.Name,
			topic,
			"logs_topic",
			false,
			nil,
		)
		if err != nil {
			return err
		}
	}

	messages, err := ch.Consume(q.Name, "", false, false, false, false, nil)
	if err != nil {
		return err
	}

	forever := make(chan bool)
	go func() {
		for d := range messages {
			var p models.Payload
			_ = json.Unmarshal(d.Body, &p)

			go handlePayload(p)
		}
	}()

	fmt.Printf("Waiting for message [Exange, Queue] [logst_topic, %s]\n", q.Name)
	<-forever
	return nil
}

func handlePayload(p models.Payload) {
	switch p.Name {
	case "log", "event":
		err := logEvent(p)
		if err != nil {
			log.Println(err)
		}
	}
}

func (cmr *consumer) setup() error {
	ch, err := cmr.conn.Channel()
	if err != nil {
		return err
	}

	err = declarExchange(ch)
	if err != nil {
		return err
	}

	return nil
}

func logEvent(p models.Payload) error {
	jsonData, _ := json.Marshal(p)

	request, err := http.NewRequest("POST", "http://logger-service/v1/logs/new", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	client := http.Client{}
	res, err := client.Do(request)
	if err != nil {
		return err
	}
	fmt.Printf("res.StatusCode: %v\n", res.StatusCode)
	return nil
}

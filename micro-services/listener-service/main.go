package main

import (
	"fmt"
	"listener/config"
	"listener/db"
	"listener/rabbitQueue"
)

func main() {
	conf, err := config.NewConfig("config.yml")
	if err != nil {
		fmt.Printf("err: %v\n", err.Error())
	}

	conn, err := db.NewRabbitConn(*conf)
	if err != nil {
		fmt.Printf("err: %v\n", err.Error())
		return
	}
	defer conn.Close()

	ch, q, err := rabbitQueue.SetRabbitConf(conn)
	if err != nil {
		fmt.Printf("err: %v\n", err.Error())
		return
	}
	defer ch.Close()

	rabbit := rabbitQueue.NewRabbit(conn, ch, q)

	topics := []string{"alert", "auth", "logs"}

	// ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	// for i, key := range topics {
	// 	s := strconv.Itoa(i)
	// 	msg := fmt.Sprintf("message number: %s, key: %s", s, key)
	// 	rabbit.Send(msg, key, ctx)
	// }
	// defer cancel()

	err = rabbit.Listen(topics)
	if err != nil {
		fmt.Printf("err: %v\n", err.Error())
		return
	}
}

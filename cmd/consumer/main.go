package main

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/souluanf/lfutils/pkg/rabbitmq"
)

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	messages := make(chan amqp.Delivery)

	go rabbitmq.Consume(ch, messages, "go-consumer", "orders")

	for msg := range messages {
		fmt.Println(string(msg.Body))
		msg.Ack(false)
	}

}

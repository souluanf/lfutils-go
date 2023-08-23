package main

import "github.com/souluanf/lfutils/pkg/rabbitmq"

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	err = rabbitmq.Publish(ch, "Hello World!", "amq.direct")
}

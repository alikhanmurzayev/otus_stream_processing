package main

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"github.com/streadway/amqp"
	"log"
)

type Config struct {
	RabbitLogin    string `envconfig:"RABBIT_LOGIN" default:"guest"`
	RabbitPassword string `envconfig:"RABBIT_PASSWORD" default:"guest"`
	RabbitHost     string `envconfig:"RABBIT_HOST" default:"localhost"`
	RabbitPort     string `envconfig:"RABBIT_PORT" default:"5672"`
}

func main() {
	var config Config
	err := envconfig.Process("", &config)
	if err != nil {
		log.Fatalf("failed to process config: %s", err)
	}

	conn, err := amqp.Dial(
		fmt.Sprintf(
			"amqp://%s:%s@%s:%s/",
			config.RabbitLogin,
			config.RabbitPassword,
			config.RabbitHost,
			config.RabbitPort,
		),
	)
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	channel, err := conn.Channel()
	if err != nil {
		log.Fatalf("could not allocate channel: %s", err)
	}
	defer channel.Close()
	if err != nil {
		log.Fatalf("could not close channel")
	}
	_, err = channel.QueueDeclare("", true, false, true, false, nil)
	if err != nil {
		log.Fatalf("could not declare queue: %s", err)
	}
	log.Println("connected to rabbit")
}

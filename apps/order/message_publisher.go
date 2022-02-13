package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"strconv"
)

type messagePublisher struct {
	channel *amqp.Channel
	queue   amqp.Queue
}

func NewMessagePublisher(conn *amqp.Connection, queueName string) (*messagePublisher, error) {
	channel, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("could not get channel: %w", err)
	}
	queue, err := channel.QueueDeclare(queueName, true, false, false, false, nil)
	if err != nil {
		return nil, fmt.Errorf("QueueDeclare: %w", err)
	}
	publisher := &messagePublisher{
		channel: channel,
		queue:   queue,
	}
	return publisher, nil
}

func (publisher *messagePublisher) Publish(ctx context.Context, event OrderEvent) error {
	body, err := json.Marshal(event)
	if err != nil {
		return err
	}
	log.Printf("Published message: %#v", event)
	return publisher.channel.Publish("", publisher.queue.Name, false, false, amqp.Publishing{
		ContentType:   "application/json",
		DeliveryMode:  amqp.Persistent,
		CorrelationId: strconv.FormatInt(event.UserID, 10),
		Body:          body,
	})
}

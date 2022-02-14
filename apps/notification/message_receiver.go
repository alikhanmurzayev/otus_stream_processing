package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

type messageReceiver struct {
	channel *amqp.Channel
	queue   amqp.Queue
}

func NewMessageReceiver(conn *amqp.Connection, queueName string) (*messageReceiver, error) {
	channel, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("could not get channel: %w", err)
	}
	queue, err := channel.QueueDeclare(queueName, true, false, false, false, nil)
	if err != nil {
		return nil, fmt.Errorf("QueueDeclare: %w", err)
	}
	receiver := &messageReceiver{
		channel: channel,
		queue:   queue,
	}
	return receiver, nil
}

func (receiver *messageReceiver) Receive(ctx context.Context) (<-chan OrderEvent, error) {
	delivery, err := receiver.channel.Consume(
		receiver.queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("could not consume: %w", err)
	}
	messages := make(chan OrderEvent)
	go func() {
		defer close(messages)

		for {
			select {
			case <-ctx.Done():
				return
			case msg := <-delivery:
				var message OrderEvent
				err := json.Unmarshal(msg.Body, &message)
				if err != nil {
					log.Printf("failed to Unmarshal: %s. Body: %s", err, string(msg.Body))
					continue
				}
				log.Printf("got message: %#v", message)
				messages <- message
			}
		}

	}()
	return messages, nil
}

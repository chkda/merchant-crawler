package queue

import (
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
)

func (c *RabbitMQClient) ReceiveMessage(ctx context.Context) (<-chan amqp.Delivery, error) {
	channel, err := c.Conn.Channel()
	if err != nil {
		return nil, err
	}
	// defer channel.Close()

	messageChan, err := channel.ConsumeWithContext(
		ctx,
		c.QueueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}
	return messageChan, nil

}

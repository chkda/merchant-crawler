package queue

import (
	"context"
	"encoding/json"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Message struct {
	ID                 string `json:"id"`
	Query              string `json:"query"`
	NormalisedMerchant string `json:"normalised_merchant"`
	MerchantLink       string `json:"merchant_link"`
	Description        string `json:"description"`
	Title              string `json:"title"`
}

func (c *RabbitMQClient) SendMessage(ctx context.Context, msg *Message) error {
	body, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	channel, err := c.Conn.Channel()
	if err != nil {
		return err
	}
	defer channel.Close()
	err = channel.PublishWithContext(
		ctx,
		"",
		c.QueueName,
		false,
		false,
		amqp.Publishing{
			Body:        body,
			ContentType: "application/json",
		},
	)
	return err
}

package queue

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQClient struct {
	Conn      *amqp.Connection
	QueueName string
}

func New(cfg *RabbitMQConfig) (*RabbitMQClient, error) {
	host := fmt.Sprintf("amqp://%s:%s@%s/", cfg.Username, cfg.Password, cfg.Host)
	conn, err := amqp.Dial(host)
	if err != nil {
		return nil, err
	}
	channel, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	defer channel.Close()
	channel.QueueDeclare(
		cfg.QueueName,
		false,
		false,
		false,
		false,
		nil,
	)
	return &RabbitMQClient{
		Conn:      conn,
		QueueName: cfg.QueueName,
	}, nil
}

package queue

type RabbitMQConfig struct {
	Host      string `json:"host"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	QueueName string `json:"queue_name"`
}

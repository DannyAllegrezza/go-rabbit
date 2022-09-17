package internal

import "github.com/rabbitmq/amqp091-go"

type Connection struct {
	Channel *amqp091.Channel
}

func GetConnection(rabbitUrl string) (Connection, error) {
	conn, err := amqp091.Dial(rabbitUrl)
	if err != nil {
		return Connection{}, err
	}

	channel, err := conn.Channel()

	return Connection{
		Channel: channel,
	}, err
}

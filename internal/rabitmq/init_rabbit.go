package rabitmq

import (
	"github.com/go_geofetch/cmd/models"
	"github.com/streadway/amqp"
)

func InitRabbitMQ(env *models.EnvModel) (*amqp.Connection, error) {
	rabbitConn, err := amqp.Dial(env.RabbitMQURL)
	if err != nil {
		return nil, err
	}
	return rabbitConn, nil
}

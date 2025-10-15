package event

import (
	"encoding/json"
	"fmt"

	"github.com/go_geofetch/cmd/models"
	"github.com/go_geofetch/internal/rabitmq/helper"
	"github.com/streadway/amqp"
)

// Consumer for receiving AMPQ events
type Consumer struct {
	conn      *amqp.Connection
	queueName string
}

func (consumer *Consumer) setup() error {
	channel, err := consumer.conn.Channel()
	if err != nil {
		return fmt.Errorf("failed to open channel: %w", err)
	}
	defer channel.Close()
	
	return helper.DeclareExchange(channel)
}

// NewConsumer returns a new Consumer
func NewConsumer(conn *amqp.Connection) (Consumer, error) {
	consumer := Consumer{
		conn: conn,
	}
	err := consumer.setup()
	if err != nil {
		return Consumer{}, err
	}

	return consumer, nil
}

// Listen will listen for all new Queue publications
// and print them to the console.
func (consumer *Consumer) Listen() error {
	ch, err := consumer.conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	q, err := helper.DeclareRandomQueue(ch)
	if err != nil {
		return err
	}

	err = ch.QueueBind(
		q.Name,
		"vehicle.*.geofence_entry",
		"fleet.events",
		false,
		nil,
	)

	if err != nil {
		return err
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",         // consumer
		true,       // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	if err != nil {
		return fmt.Errorf("failed to register consumer: %w", err)
	}

	for msg := range msgs {
		var event models.GeoFenceEvent
		if err := json.Unmarshal(msg.Body, &event); err != nil {
			fmt.Println("Error unmarshalling message:", err)
			continue
		}
		fmt.Printf("Received geofence event: %+v\n", event)
	}

	return nil
}

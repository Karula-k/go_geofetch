package event

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/go_geofetch/cmd/models"
	"github.com/go_geofetch/internal/rabitmq/helper"
	"github.com/streadway/amqp"
)

// Emitter for publishing AMQP events
type Emitter struct {
	connection *amqp.Connection
}

func (e *Emitter) setup() error {
	channel, err := e.connection.Channel()
	if err != nil {
		return fmt.Errorf("failed to open channel: %w", err)
	}

	defer channel.Close()
	return helper.DeclareExchange(channel)
}

// Push (Publish) a specified message to the AMQP exchange
func (e *Emitter) Push(event models.GeoFenceEvent) error {
	channel, err := e.connection.Channel()
	if err != nil {
		return err
	}

	body, err := json.Marshal(event)

	if err != nil {
		return fmt.Errorf("failed to marshal event: %w", err)
	}
	defer channel.Close()

	routingKey := fmt.Sprintf("vehicle.%s.%s", event.VehicleID, event.Event)

	err = channel.Publish(
		"fleet.events",
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	log.Printf("Sending message:  -> %s", routingKey)
	return nil
}

// NewEventEmitter returns a new event.Emitter object
// ensuring that the object is initialised, without error
func NewEventEmitter(conn *amqp.Connection) (Emitter, error) {
	emitter := Emitter{
		connection: conn,
	}

	err := emitter.setup()
	if err != nil {
		return Emitter{}, err
	}

	return emitter, nil
}

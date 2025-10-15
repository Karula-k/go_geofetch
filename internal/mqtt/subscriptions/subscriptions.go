package subscriptions

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/go_geofetch/cmd/models"
	"github.com/go_geofetch/generated"
	"github.com/go_geofetch/internal/mqtt/subscriptions/service"
	"github.com/go_geofetch/internal/rabitmq/event"
	"github.com/streadway/amqp"
)

const geofenceRadius = 50.0

func MQTTSubscription(ctx context.Context, queries *generated.Queries, env *models.EnvModel, client mqtt.Client, rabbitConn *amqp.Connection) {
	// Create emitter for publishing events
	emitter, err := event.NewEventEmitter(rabbitConn)
	if err != nil {
		fmt.Println("Error creating event emitter:", err)
		return
	}

	if token := client.Subscribe("/fleet/vehicle/+/location", 0, func(client mqtt.Client, msg mqtt.Message) {
		payload := string(msg.Payload())

		payload = strings.Trim(payload, "'")

		var vehicleLocation generated.VehicleLocation
		err := json.Unmarshal([]byte(payload), &vehicleLocation)
		if err != nil {
			fmt.Println("Error unmarshalling message:", err, payload)
			return
		}

		// Save to database
		_, err = queries.CreateVehicleLocation(ctx, generated.CreateVehicleLocationParams{
			VehicleID: vehicleLocation.VehicleID,
			Latitude:  vehicleLocation.Latitude,
			Longitude: vehicleLocation.Longitude,
			Timestamp: vehicleLocation.Timestamp,
		})
		if err != nil {
			fmt.Println("Error creating vehicle location:", err, vehicleLocation)
		}

		// Check geofence
		go service.GeoFenceTrigger(models.Location{
			Latitude:  -6.2088,
			Longitude: 106.8456,
		}, vehicleLocation, &emitter, geofenceRadius)

		fmt.Printf("Received message on topic %s: %+v\n", msg.Topic(), vehicleLocation)
	}); token.Wait() && token.Error() != nil {
		fmt.Println("Subscription error:", token.Error())
	}

}

package subscriptions

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/go_geofetch/cmd/models"
	"github.com/go_geofetch/generated"
)

func MQTTSubscription(ctx context.Context, queries *generated.Queries, env *models.EnvModel, client mqtt.Client) {
	if token := client.Subscribe("/fleet/vehicle/+/location", 0, func(client mqtt.Client, msg mqtt.Message) {
		payload := string(msg.Payload())

		payload = strings.Trim(payload, "'")

		var vehicleLocation generated.VehicleLocation
		err := json.Unmarshal([]byte(payload), &vehicleLocation)
		if err != nil {
			fmt.Println("Error unmarshalling message:", err, payload)
			return
		}

		_, err = queries.CreateVehicleLocation(ctx, generated.CreateVehicleLocationParams{
			VehicleID: vehicleLocation.VehicleID,
			Latitude:  vehicleLocation.Latitude,
			Longitude: vehicleLocation.Longitude,
			Timestamp: vehicleLocation.Timestamp,
		})
		if err != nil {
			fmt.Println("Error creating vehicle location:", err, vehicleLocation)
		}
		fmt.Printf("Received message on topic %s: %+v\n", msg.Topic(), vehicleLocation)
	}); token.Wait() && token.Error() != nil {
		fmt.Println("Subscription error:", token.Error())
	}

}

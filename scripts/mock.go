package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type VehicleLocation struct {
	VehicleID string  `json:"vehicle_id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Timestamp int64   `json:"timestamp"`
}

func main() {
	vehicleID := "MOCK123"
	opts := mqtt.NewClientOptions().AddBroker("tcp://localhost:1883")
	opts.SetClientID(os.Getenv("MOCK_CLIENT_ID"))

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		fmt.Println("Failed to connect to MQTT broker:", token.Error())
	}
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	geoFenceLat := -6.2088
	geoFenceLon := 106.8456

	jitter := 0.0005

	for {
		select {
		case <-ticker.C:
			data := VehicleLocation{
				VehicleID: vehicleID,
				Latitude:  geoFenceLat + (rand.Float64()-0.5)*2*jitter,
				Longitude: geoFenceLon + (rand.Float64()-0.5)*2*jitter,
				Timestamp: time.Now().Unix(),
			}
			payload, _ := json.Marshal(data)
			topic := fmt.Sprintf("/fleet/vehicle/%s/location", vehicleID)
			client.Publish(topic, 0, false, payload)

		case <-sigChan:
			fmt.Println("Disconnect")
			client.Disconnect(250)
			return
		}
	}
}

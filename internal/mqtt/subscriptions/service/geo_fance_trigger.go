package service

import (
	"fmt"
	"math"

	"github.com/go_geofetch/cmd/models"
	"github.com/go_geofetch/generated"
	"github.com/go_geofetch/internal/rabitmq/event"
)

func GeoFenceTrigger(point models.Location, vehicleLocation generated.VehicleLocation, emitter *event.Emitter, geoFenceRadius float64) {
	distance := haversineDistance(
		vehicleLocation.Latitude, vehicleLocation.Longitude,
		point.Latitude, point.Longitude,
	)

	if distance <= geoFenceRadius {
		geofenceEvent := models.GeoFenceEvent{
			VehicleID: vehicleLocation.VehicleID,
			Event:     "geofence_entry",
			Location: models.Location{
				Latitude:  vehicleLocation.Latitude,
				Longitude: vehicleLocation.Longitude,
			},
			Timestamp: vehicleLocation.Timestamp,
		}

		err := emitter.Push(geofenceEvent)
		if err != nil {
			fmt.Println("Error publishing geofence event:", err)
		} else {
			fmt.Printf("Geofence event published for vehicle %s (distance: %.2f meters)\n", vehicleLocation.VehicleID, distance)
		}
	}

}

func haversineDistance(lat1, lon1, lat2, lon2 float64) float64 {
	const earthRadius = 6371000 // Earth radius in meters

	phi1 := lat1 * math.Pi / 180
	phi2 := lat2 * math.Pi / 180
	deltaPhi := (lat2 - lat1) * math.Pi / 180
	deltaLambda := (lon2 - lon1) * math.Pi / 180

	a := math.Sin(deltaPhi/2)*math.Sin(deltaPhi/2) +
		math.Cos(phi1)*math.Cos(phi2)*
			math.Sin(deltaLambda/2)*math.Sin(deltaLambda/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return earthRadius * c
}

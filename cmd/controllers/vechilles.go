package controllers

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/go_geofetch/cmd/models"
	"github.com/go_geofetch/generated"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgtype"
)

// DummyVehicleController simulates a dummy endpoint for vehicle location.
// It listens for incoming requests, creates a location message, and publishes it to MQTT.
//
//	@Summary		Simulate vehicle location update
//	@Description	Receives a request, creates a dummy vehicle location, and publishes it to MQTT
//	@Tags			Vehicle
//	@Accept			json
//	@Produce		json
//	@Param			request	body		models.VehiclesDummyRequest	true	"Dummy vehicle location data"
//	@Success		200		{string}	string						"Vehicle location published"
//	@Router			/vehicles/dummy [post]
func DummyVehicleController(ctx context.Context, client mqtt.Client, env *models.EnvModel) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req models.VehiclesDummyRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
		}
		vehicleID := "B1234XYZ"
		if req.VehicleID != nil {
			vehicleID = *req.VehicleID
		}

		dummyLocation := models.VehicleLocationHelper{
			VehicleID: vehicleID,
			Latitude:  -6.2088,
			Longitude: 106.8456,
			Timestamp: 1715003456,
		}

		payload, err := json.Marshal(dummyLocation)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to marshal json"})
		}

		client.Publish(`/fleet/vehicle/`+dummyLocation.VehicleID+`/location`, 0, false, payload)

		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "dummy vehicle sended"})
	}
}

// LastLocationController retrieves the latest location of a vehicle.
// It listens for incoming requests, queries the database for the latest location,
// and returns the location information as a JSON response.
//
//	@Summary		Retrieve latest vehicle location
//	@Description	Queries the database for the latest location of a vehicle
//	@Tags			Vehicle
//	@Accept			json
//	@Produce		json
//	@Param			vehicle_id	path		string	true	"Vehicle ID"
//	@Success		200			{object}	generated.VehicleLocation
//	@Router			/vehicles/{vehicle_id}/location [get]
func LastLocationController(ctx context.Context, queries *generated.Queries) fiber.Handler {
	return func(c *fiber.Ctx) error {
		vehicleID := c.Params("vehicle_id")
		if vehicleID == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "vehicle_id is required"})
		}
		location, err := queries.GetVehicleLocation(ctx, vehicleID)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error":      "location not found",
				"vehicle_id": vehicleID,
				"db_error":   err.Error()})
		}
		return c.Status(fiber.StatusOK).JSON(location)
	}
}

// HistoryLocationController retrieves the history of locations for a vehicle.
// It listens for incoming requests, queries the database for the location history,
// and returns the location information as a JSON response.
//
//	@Summary		Retrieve vehicle location history
//	@Description	Queries the database for the location history of a vehicle
//	@Tags			Vehicle
//	@Accept			json
//	@Produce		json
//	@Param			vehicle_id	path	string	true	"Vehicle ID"
//	@Param			start		query	int64	false	"Start timestamp (Unix timestamp)"
//	@Param			end			query	int64	false	"End timestamp (Unix timestamp)"
//	@Success		200			{array}	generated.VehicleLocation
//	@Router			/vehicles/{vehicle_id}/history [get]
func HistoryLocationController(ctx context.Context, queries *generated.Queries) fiber.Handler {
	return func(c *fiber.Ctx) error {
		vehicleID := c.Params("vehicle_id")
		start := c.Query("start")
		end := c.Query("end")
		if vehicleID == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "vehicle_id is required"})
		}

		var parsedStart, parsedEnd *time.Time

		if start != "" {
			startUnix, err := strconv.ParseInt(start, 10, 64)
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid start timestamp"})
			}
			startTime := time.Unix(startUnix, 0)
			parsedStart = &startTime
		}

		if end != "" {
			endUnix, err := strconv.ParseInt(end, 10, 64)
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid end timestamp"})
			}
			endTime := time.Unix(endUnix, 0)
			parsedEnd = &endTime
		}

		var startTimestamp, endTimestamp pgtype.Int8

		if parsedStart != nil {
			startTimestamp = pgtype.Int8{Int64: parsedStart.Unix(), Valid: true}
		}

		if parsedEnd != nil {
			endTimestamp = pgtype.Int8{Int64: parsedEnd.Unix(), Valid: true}
		}

		locations, err := queries.GetVehicleHistory(ctx, generated.GetVehicleHistoryParams{
			VehicleID: vehicleID,
			StartDate: startTimestamp,
			EndDate:   endTimestamp,
		})
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error":      "location not found",
				"vehicle_id": vehicleID,
				"db_error":   err.Error(),
			})
		}

		if locations == nil {
			locations = []generated.VehicleLocation{}
		}
		return c.Status(fiber.StatusOK).JSON(locations)
	}
}

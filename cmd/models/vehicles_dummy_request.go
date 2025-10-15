package models

type VehiclesDummyRequest struct {
	VehicleID *string  `json:"vehicle_id,omitempty"`
	Latitude  *float64 `json:"latitude,omitempty"`
	Longitude *float64 `json:"longitude,omitempty"`
	Timestamp *int64   `json:"timestamp,omitempty"`
}

-- name: ListVehicleLocation :many
SELECT * FROM vehicle_location
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: CreateVehicleLocation :one
INSERT INTO vehicle_location(
  vehicle_id,
  latitude,
  longitude,
  timestamp
) VALUES (
  $1, $2, $3, $4
)RETURNING *;

-- name: DeleteVehicleLocation :exec
DELETE FROM vehicle_location
WHERE id = $1;

-- name: UpdateVehicleLocation :exec
UPDATE vehicle_location
SET vehicle_id = $2,
    latitude = $3,
    longitude = $4,
    timestamp = $5
WHERE id =$1;

-- name: GetVehicleHistory :many
SELECT * FROM vehicle_location
WHERE vehicle_id = $1
AND (sqlc.narg('start_date')::BIGINT IS NULL OR timestamp >= sqlc.narg('start_date'))
AND (sqlc.narg('end_date')::BIGINT IS NULL OR timestamp <= sqlc.narg('end_date'))
ORDER BY timestamp DESC;

-- name: GetVehicleLocation :one
SELECT * FROM vehicle_location
WHERE vehicle_id = $1
ORDER BY timestamp DESC
LIMIT 1;
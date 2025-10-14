CREATE TABLE vehicle_location (
  id         BIGSERIAL PRIMARY KEY,
  vehicle_id text NOT NULL UNIQUE,
  latitude   FLOAT NOT NULL,
  longitude  FLOAT NOT NULL,
  timestamp  TIMESTAMPTZ
);
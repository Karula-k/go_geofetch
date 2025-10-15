package config

import (
	"errors"
	"os"

	"github.com/go_geofetch/cmd/models"
)

func EnvConfig() (models.EnvModel, error) {
	jwtSecret := os.Getenv("JWT_SECRET")
	port := os.Getenv("PORT")
	postgresUser := os.Getenv("POSTGRES_USER")
	postgresPassword := os.Getenv("POSTGRES_PASSWORD")
	postgresDB := os.Getenv("POSTGRES_DB")
	databaseUrl := os.Getenv("DATABASE_URL")
	mqttBroker := os.Getenv("MQTT_BROKER")
	mqttClientID := os.Getenv("MQTT_CLIENT_ID")
	rabbitMQURL := os.Getenv("RABBITMQ_URL")

	if databaseUrl == "" {
		return models.EnvModel{}, errors.New("DATABASE_URL is not set")
	}
	if jwtSecret == "" {
		return models.EnvModel{}, errors.New("JWT_SECRET is not set")
	}
	if port == "" {
		return models.EnvModel{}, errors.New("PORT is not set")
	}
	if postgresUser == "" {
		return models.EnvModel{}, errors.New("POSTGRES_USER is not set")
	}
	if postgresPassword == "" {
		return models.EnvModel{}, errors.New("POSTGRES_PASSWORD is not set")
	}
	if postgresDB == "" {
		return models.EnvModel{}, errors.New("POSTGRES_DB is not set")
	}
	if mqttBroker == "" {
		return models.EnvModel{}, errors.New("MQTT_BROKER is not set")
	}
	if mqttClientID == "" {
		return models.EnvModel{}, errors.New("MQTT_CLIENT_ID is not set")
	}
	if rabbitMQURL == "" {
		return models.EnvModel{}, errors.New("RABBITMQ_URL is not set")
	}

	model := models.EnvModel{
		DatabaseUrl:      databaseUrl,
		JwtSecret:        jwtSecret,
		Port:             port,
		MQTTBroker:       mqttBroker,
		MQTTClientID:     mqttClientID,
		PostgresUser:     postgresUser,
		PostgresPassword: postgresPassword,
		PostgresDB:       postgresDB,
		RabbitMQURL:      rabbitMQURL,
	}

	return model, nil
}

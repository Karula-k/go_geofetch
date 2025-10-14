package mqtt

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/go_geofetch/cmd/models"
)

func InitMQTT(env *models.EnvModel) (mqtt.Client, error) {
	opts := mqtt.NewClientOptions().AddBroker(env.MQTTBroker)
	opts.SetClientID(env.MQTTClientID)

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}

	return client, nil
}

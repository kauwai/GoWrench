package connection_settings

import (
	"wrench/app/manifest/validation"
)

type NatsConnectionSettings struct {
	Id            string `yaml:"id"`
	ServerAddress string `yaml:"serverAddress"`
}

func (setting *NatsConnectionSettings) GetId() string {
	return setting.Id
}

func (settings NatsConnectionSettings) Valid() validation.ValidateResult {
	var result validation.ValidateResult

	if len(settings.Id) == 0 {
		result.AddError("connections.nats.id is required")
	}

	if len(settings.ServerAddress) == 0 {
		result.AddError("the connections.nats.serverAddress is required")
	}

	return result
}

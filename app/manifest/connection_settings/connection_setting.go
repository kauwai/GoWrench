package connection_settings

import (
	"wrench/app/manifest/validation"
)

type ConnectionSettings struct {
	Nats  []*NatsConnectionSettings  `yaml:"nats"`
	Kafka []*KafkaConnectionSettings `yaml:"kafka"`
	Redis []*RedisConnectionSettings `yaml:"redis"`
}

func (settings ConnectionSettings) Valid() validation.ValidateResult {
	var result validation.ValidateResult

	if len(settings.Nats) > 0 {
		for _, validable := range settings.Nats {
			result.AppendValidable(validable)
		}
	}

	if len(settings.Kafka) > 0 {
		for _, validable := range settings.Kafka {
			result.AppendValidable(validable)
		}
	}

	if len(settings.Redis) > 0 {
		for _, validable := range settings.Redis {
			result.AppendValidable(validable)
		}
	}

	return result
}

package config

import (
	"os"

	"github.com/spf13/cast"
)

type ProducerConfig struct {
	KafkaBootstrapServers string
	KafkaTopic string
	MessageSendingRate int
}

func LoadProducerConfig() (cfg ProducerConfig) {

	cfg.KafkaBootstrapServers = cast.ToString(os.Getenv("KAFKA_BOOTSTRAP_SERVERS"))
	cfg.KafkaTopic= cast.ToString(os.Getenv("KAFKA_TOPIC"))
	cfg.MessageSendingRate= cast.ToInt(os.Getenv("MESSAGE_SENDING_RATE"))

	return cfg
}

package config

import (
	"log"
	"os"

	"github.com/spf13/cast"
)

type ProducerConfig struct {
	KafkaBootstrapServers string
	KafkaTopic string
}

func LoadProducerConfig() (cfg ProducerConfig) {

	cfg.KafkaBootstrapServers = cast.ToString(os.Getenv("KAFKA_BOOTSTRAP_SERVERS"))
	cfg.KafkaTopic= cast.ToString(os.Getenv("KAFKA_TOPIC"))

	return cfg
}

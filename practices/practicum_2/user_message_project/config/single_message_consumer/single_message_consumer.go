package config

import (
	"os"

	"github.com/spf13/cast"
)

type SingleMessageConfig struct {
	KafkaBootstrapServers string
	KafkaTopic string
	KafkaSessionTimeout int
	KafkaConsumerTimeout int
	KafkaEnableAutoCommit bool
	KafkaGroupId int
}

func LoadSingleMessageConfig() (cfg SingleMessageConfig) {

	cfg.KafkaBootstrapServers = cast.ToString(os.Getenv("KAFKA_BOOTSTRAP_SERVERS"))
	cfg.KafkaTopic= cast.ToString(os.Getenv("KAFKA_TOPIC"))
	cfg.KafkaSessionTimeout= cast.ToInt(os.Getenv("KAFKA_SESSION_TIMEOUT"))
	cfg.KafkaConsumerTimeout= cast.ToInt(os.Getenv("KAFKA_CONSUMER_TIMEOUT"))
	cfg.KafkaGroupId= cast.ToInt(os.Getenv("KAFKA_GROUP_ID"))

	return cfg
}

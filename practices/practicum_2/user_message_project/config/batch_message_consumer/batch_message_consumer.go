package config

import (
	"os"

	"github.com/spf13/cast"
)

type BatchMessageConfig struct {
	KafkaBootstrapServers string
	KafkaTopic string
	KafkaSessionTimeout int
	KafkaConsumerTimeout int
	KafkaEnableAutoCommit bool
	KafkaGroupId int
	KafkaBatchSize int
}

func LoadBatchMessageConfig() (cfg BatchMessageConfig) {

	cfg.KafkaBootstrapServers = cast.ToString(os.Getenv("KAFKA_BOOTSTRAP_SERVERS"))
	cfg.KafkaTopic= cast.ToString(os.Getenv("KAFKA_TOPIC"))
	cfg.KafkaSessionTimeout= cast.ToInt(os.Getenv("KAFKA_SESSION_TIMEOUT"))
	cfg.KafkaConsumerTimeout= cast.ToInt(os.Getenv("KAFKA_CONSUMER_TIMEOUT"))
	cfg.KafkaEnableAutoCommit= cast.ToBool(os.Getenv("KAFKA_ENABLE_AUTO_COMMIT"))
	cfg.KafkaGroupId= cast.ToInt(os.Getenv("KAFKA_GROUP_ID"))
	cfg.KafkaBatchSize= cast.ToInt(os.Getenv("KAFKA_BATCH_SIZE"))

	return cfg
}

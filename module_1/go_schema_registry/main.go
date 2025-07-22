package main

import (
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/confluentinc/confluent-kafka-go/schemaregistry"
	"github.com/confluentinc/confluent-kafka-go/schemaregistry/serde"
	"github.com/confluentinc/confluent-kafka-go/schemaregistry/serde/jsonschema"
)

// Определение структуры сообщения
type Product struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	// Конфигурация для Schema Registry
	srClient, err := schemaregistry.NewClient(schemaregistry.NewConfig("http://localhost:8081"))
	if err != nil {
		log.Fatalf("Ошибка при создании клиента Schema Registry: %v", err)
	}

	// Создание сериализатора JSON
	serializer, err := jsonschema.NewSerializer(srClient, serde.ValueSerde, jsonschema.NewSerializerConfig())
	if err != nil {
		log.Fatalf("Ошибка при создании сериализатора: %v", err)
	}

	// Конфигурация для Kafka Producer
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9094",
	})
	if err != nil {
		log.Fatalf("Ошибка при создании продюсера: %v", err)
	}
	defer producer.Close()

	// Тема Kafka
	topic := "your-topic"

	// Сообщение для отправки
	Product := Product{
		Name: "Product",
		Id:   30,
	}

	// Сериализация сообщения
	payload, err := serializer.Serialize(topic, &Product)
	if err != nil {
		log.Fatalf("Ошибка при сериализации сообщения: %v", err)
	}

	// Отправка сообщения
	deliveryChan := make(chan kafka.Event)
	err = producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          payload,
		Key:            []byte("Product_key"),
	}, deliveryChan)

	if err != nil {
		log.Fatalf("Ошибка при отправке сообщения: %v", err)
	}

	// Ожидание подтверждения доставки
	e := <-deliveryChan
	m := e.(*kafka.Message)

	if m.TopicPartition.Error != nil {
		fmt.Printf("Ошибка доставки: %v\n", m.TopicPartition.Error)
	} else {
		fmt.Printf("Сообщение доставлено в %v\n", m.TopicPartition)
	}

	close(deliveryChan)
}

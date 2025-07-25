// main.go
package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"

	"github.com/svirskey/kafka-practicum/practices/practicum_2/user_message_project/config/producer"
	"github.com/svirskey/kafka-practicum/practices/practicum_2/user_message_project/internal/model"
	"github.com/svirskey/kafka-practicum/practices/practicum_2/user_message_project/pkg/random"
)


func main() {

	producerCfg := config.LoadProducerConfig()

	// Создаём продюсера
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": producerCfg.KafkaBootstrapServers,
		"acks": "1", // at least once guarantee
	})
	if err != nil {
		log.Fatalf("Невозможно создать продюсера: %s\n", err)
		return
	}
	
	log.Printf("Продюсер создан %v\n", p)

	// Канал доставки событий (информации об отправленном сообщении)
	deliveryChan := make(chan kafka.Event)
	duration := time.Duration(float64(time.Second) / float64(producerCfg.MessageSendingRate))
	for {
		// Создаём сообщение
		value := &model.UserMessage{
			UserId: random.RandRange(0, 1000),
			Message:  random.RandStringBytes(32),
			Timestamp: time.Now(),
		}

		// Сериализуем заказ в массив
		payload, err := json.Marshal(value)
		if err != nil {
			log.Fatalf("Невозможно сериализовать заказ: %s\n", err)
		}

		// Отправляем сообщение в брокер
		err = p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &producerCfg.KafkaTopic, Partition: kafka.PartitionAny},
			Value:          payload,
			Headers:        nil,
		}, deliveryChan)
		if err != nil {
			log.Fatalf("Ошибка при отправке сообщения: %v\n", err)
		}

		// Ждём информацию об отправленном сообщении. (https://docs.confluent.io/kafka-clients/go/current/overview.html#synchronous-writes)
		e := <-deliveryChan

		// Приводим Events к типу *kafka.Message
		m := e.(*kafka.Message)

		// Если возникла ошибка доставки сообщения
		if m.TopicPartition.Error != nil {
			log.Printf("Ошибка доставки сообщения: %v\n", m.TopicPartition.Error)
		} else {
			log.Printf("Сообщение %s отправлено в топик %s [%d] оффсет %v\n", payload,
				*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
		}
		time.Sleep(duration)
	}
}

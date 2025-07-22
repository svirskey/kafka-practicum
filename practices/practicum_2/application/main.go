// main.go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

// UserMessage – структура, описывающая сообщение пользователя
type UserMessage struct {
	UserId    string    `json:"user_id"`
	Message   string    `json:"message"`
	timestamp time.Time `json:"timestamp"`
}

func main() {

	// Проверяем, что количество параметров при запуске программы ровно 3
	if len(os.Args) != 3 {
		log.Fatalf("Пример использования: %s <bootstrap-servers> <topic>\n", os.Args[0])
	}

	// Парсим параметры и получаем адрес брокера и имя топика
	bootstrapServers := os.Args[1]
	topic := os.Args[2]

	// Создаём продюсера
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": bootstrapServers})
	if err != nil {
		log.Fatalf("Невозможно создать продюсера: %s\n", err)
	}

	log.Printf("Продюсер создан %v\n", p)

	// Канал доставки событий (информации об отправленном сообщении)
	deliveryChan := make(chan kafka.Event)

	// Создаём заказ
	value := &Order{
		OrderID: "0001",
		UserID:  "00001",
		Items: []Item{
			{ProductID: "535", Quantity: 1, Price: 300},
			{ProductID: "125", Quantity: 2, Price: 100},
		},
		TotalPrice: 500.00,
	}

	// Сериализуем заказ в массив
	payload, err := json.Marshal(value)
	if err != nil {
		log.Fatalf("Невозможно сериализовать заказ: %s\n", err)
	}

	// Отправляем сообщение в брокер
	err = p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          payload,
		Headers:        []kafka.Header{{Key: "myTestHeader", Value: []byte("header values are binary")}},
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
		fmt.Printf("Ошибка доставки сообщения: %v\n", m.TopicPartition.Error)
	} else {
		fmt.Printf("Сообщение отправлено в топик %s [%d] офсет %v\n",
			*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
	}
	// Закрываем продюсера
	p.Close()
	// Закрываем канал доставки событий
	close(deliveryChan)
}

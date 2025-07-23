// main.go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	cfg "github.com/svirskey/kafka-practicum/practices/practicum_2/user_message_project/config/single_message_consumer"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)


func main() {

	config := cfg.LoadSingleMessageConfig()

	// Перехватываем сигналы syscall.SIGINT и syscall.SIGTERM для graceful shutdown
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	// Создаём консьюмера
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":  config.KafkaBootstrapServers,
		"group.id":           config.KafkaGroupId,
		"session.timeout.ms": config.KafkaSessionTimeout,
		"auto.offset.reset":  "earliest",
		"enable-auto-commit": config.KafkaEnableAutoCommit})

	if err != nil {
		log.Fatalf("Невозможно создать консьюмера: %s\n", err)
	}

	fmt.Printf("Консьюмер создан %v\n", c)

	// Подписываемся на топик
	err = c.SubscribeTopics(config.KafkaTopic, nil)

	if err != nil {
		log.Fatalf("Невозможно подписаться на топик: %s\n", err)
	}

	run := true
	// Запускаем бесконечный цикл
	for run {
		select {
		// Для выхода нажмите ctrl+C
		case sig := <-sigchan:
			fmt.Printf("Передан сигнал %v: приложение останавливается\n", sig)
			run = false
		default:

			// Делаем запрос на считывание сообщения из брокера
			ev := c.Poll(timeconfig.KafkaConsumerTimeout)
			if ev == nil {
				continue
			}

			//     Приводим Events к
			switch e := ev.(type) {
			// типу *kafka.Message,
			case *kafka.Message:
				value := Order{}
				err := json.Unmarshal(e.Value, &value)
				if err != nil {
					fmt.Printf("Ошибка десериализации: %s\n", err)
				} else {
					fmt.Printf("%% Получено сообщение в топик %s:\n%+v\n", e.TopicPartition, value)
				}
				if e.Headers != nil {
					fmt.Printf("%% Заголовки: %v\n", e.Headers)
				}
			// типу Ошибки брокера
			case kafka.Error:
				// Ошибки обычно следует считать
				// информационными, клиент попытается
				// автоматически их восстановить.
				fmt.Fprintf(os.Stderr, "%% Error: %v: %v\n", e.Code(), e)
			default:
				fmt.Printf("Другие события %v\n", e)
			}
		}
	}
	// Закрываем консьюмера
	c.Close()
}

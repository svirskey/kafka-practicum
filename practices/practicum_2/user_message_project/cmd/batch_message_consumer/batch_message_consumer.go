// main.go
package main

import (
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	cfg "github.com/svirskey/kafka-practicum/practices/practicum_2/user_message_project/config/batch_message_consumer"
	"github.com/svirskey/kafka-practicum/practices/practicum_2/user_message_project/internal/model"
)

func main() {

	config := cfg.LoadBatchMessageConfig()

	// Перехватываем сигналы syscall.SIGINT и syscall.SIGTERM для graceful shutdown
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	// Создаём консьюмера
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":  config.KafkaBootstrapServers,
		"group.id":           config.KafkaGroupId,
		"session.timeout.ms": config.KafkaSessionTimeout,
		"fetch.min.bytes":    1040,
		"enable.auto.commit": config.KafkaEnableAutoCommit})

	if err != nil {
		log.Fatalf("Невозможно создать консьюмера: %s\n", err)
	}

	log.Printf("Консьюмер создан %v\n", c)

	// Подписываемся на топик
	err = c.Subscribe(config.KafkaTopic, nil)

	if err != nil {
		log.Fatalf("Невозможно подписаться на топик: %s\n", err)
	}

	run := true
	messageCount := 0
	var messages []*kafka.Message
	// Запускаем бесконечный цикл
	for run {
		select {
		// Для выхода нажмите ctrl+C
		case sig := <-sigchan:
			log.Printf("Передан сигнал %v: приложение останавливается\n", sig)
			run = false
		default:

			// Делаем запрос на считывание сообщения из брокера
			msg, err := c.ReadMessage(time.Second * time.Duration(config.KafkaConsumerTimeout))
			if err != nil {
				log.Printf("Ошибка чтения сообщения: %s\n", err)
				continue
			}

			messages = append(messages, msg)
			messageCount++

			if messageCount%config.KafkaBatchSize == 0 {

				log.Printf("Обработка пачки сообщений")

				log.Printf(strconv.Itoa(len(messages)))
				for _, msg := range messages {
					value := model.UserMessage{}
					err := json.Unmarshal(msg.Value, &value)
					if err != nil {
						log.Printf("Ошибка десериализации: %s\n", err)
					} else {
						log.Printf("%% Обработано сообщение в топике %s:\n%+v\n", msg.TopicPartition, value)
					}
				}
				_, err = c.Commit()

				if err != nil {
					log.Printf("Ошибка коммита: %s\n", err)
				}
				log.Printf("Закоммичено")

				messages = []*kafka.Message{}
			}
		}
	}
	// Закрываем консьюмера
	err = c.Close()
	if err != nil {
		log.Printf("Ошибка закрытия консьюмера: %s\n", err)
	}
}

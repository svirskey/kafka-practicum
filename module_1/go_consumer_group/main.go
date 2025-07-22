// main.go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

// Item – структура, описывающая продукт в заказе
type Item struct {
	ProductID string  `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}

// Order – структура, описывающая заказ с продуктами
type Order struct {
	Offset     int     `json:"offset"`
	OrderID    string  `json:"order_id"`
	UserID     string  `json:"user_id"`
	Items      []Item  `json:"items"`
	TotalPrice float64 `json:"total_price"`
}

// Таймаут для запроса
const timeoutMs = 100

func main() {

	if len(os.Args) < 4 {
		log.Fatalf("Пример использования: %s <bootstrap-servers> <group> <topics..>\n",
			os.Args[0])
	}
	// Парсим параметры и получаем адрес брокера, группу и имя топиков
	bootstrapServers := os.Args[1]
	group := os.Args[2]
	topics := os.Args[3:]

	// Перехватываем сигналы syscall.SIGINT и syscall.SIGTERM для graceful shutdown
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	// Создаём консьюмера
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":  bootstrapServers,
		"group.id":           group,       // уникальный идентификатор группы
		"session.timeout.ms": 6000,        // время, в течение которого Kafka ожидает активности от консьюмера до того, как объявит его «мёртвым» и начнёт ребалансировку
		"enable.auto.commit": true,        // автоматически сохранять смещения (offsets) после получения сообщений
		"auto.offset.reset":  "earliest"}) // начинать чтение с самого старого доступного сообщения.

	if err != nil {
		log.Fatalf("Невозможно создать консьюмера: %s\n", err)
	}

	fmt.Printf("Консьюмер создан %v\n", c)

	// Подписываемся на топики, в примере он один
	err = c.SubscribeTopics(topics, nil)

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
			ev := c.Poll(timeoutMs)
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
				// автоматически их восстановить
				fmt.Fprintf(os.Stderr, "%% Error: %v: %v\n", e.Code(), e)
			default:
				fmt.Printf("Другие события %v\n", e)
			}
		}
	}
	// Закрываем потребителя
	c.Close()
}

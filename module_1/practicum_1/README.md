## Установка и запуск с использованием Zookeeper

Запуск Kafka с Zookeeper ― классический способ развёртывания. Быстро установить и настроить кластер Apache Kafka помогает Docker. 
Для этого нужно:

- Установить и запустить Docker.
- Установить Docker Compose.
- Выделить в Docker как минимум 8 Гб оперативной памяти.
- Иметь доступ в Интернет.
- Создадим docker-compose-файл для запуска.

```
version: '3'
services:
  zookeeper:
    image: confluentinc/cp-zookeeper:7.0.1
    environment:
      ZOOKEEPER_CLIENT_PORT: 2181
      ZOOKEEPER_TICK_TIME: 2000
    ports:
      - "2181:2181"
  kafka:
    image: confluentinc/cp-kafka:7.0.1
    depends_on:
      - zookeeper
    environment:
      KAFKA_BROKER_ID: 1
      KAFKA_ZOOKEEPER_CONNECT: zookeeper:2181
      KAFKA_ADVERTISED_LISTENERS: PLAINTEXT://localhost:9092
      KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR: 1
    ports:
      - "9092:9092"
```

В докер-файле используется образ confluentinc/cp-zookeeper:7.0.1 ― контейнер, в котором запускается Apache Zookeeper. Это обязательный компонент для старых версий Kafka. 
Образ confluentinc/cp-zookeeper принадлежит компании Confluent, которая разработала расширения и инструменты для работы с Apache Kafka. Использование этого образа гарантирует, что вы получите версию Zookeeper, которая проверена, поддерживается и совместима с Kafka. 
Вот объяснение ключевых переменных окружения для настройки Kafka. Объяснение справочное, более детально вы разберётесь с переменными в следующих уроках.

- ZOOKEEPER_CLIENT_PORT указывает порт, на котором Zookeeper будет слушать входящие подключения от клиентов. Порт 2181 ― стандартный для Zookeeper. Он используется для взаимодействия с приложениями, которые требуют доступа к Zookeeper.
- ZOOKEEPER_TICK_TIME определяет базовый временной интервал (в миллисекундах) для управления сессиями. Это время используется для тайм-аутов и периодов обновления.
- KAFKA_BROKER_ID ― это уникальный идентификатор для брокера в кластере Kafka. Каждый брокер имеет уникальный ID, чтобы другие брокеры и клиенты отличали его от других.
- KAFKA_ZOOKEEPER_CONNECT указывает адрес и порт Zookeeper, к которому брокер Kafka должен подключаться, чтобы получить информацию о состоянии кластера и управлять им. Формат zookeeper:2181 указывает на имя сервиса Zookeeper и порт.
- KAFKA_ADVERTISED_LISTENERS определяет, как брокер будет сообщать другим компонентам (например, клиентам), как подключаться к нему. Здесь используется простой текстовый протокол (PLAINTEXT) и адрес localhost с портом 9092.
- KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR указывает количество реплик для топика смещений (__consumer_offsets), которая хранит информацию о смещениях потребителей. Значение 1 означает, что не будет резервных копий смещения.

Запускаем Docker Compose, чтобы создать и запустить контейнеры:

```
docker-compose up -d 
```

- Проверяем успешность запуска. Убедитесь, что контейнеры работают:

```
docker ps 
```
Вы должны увидеть два работающих контейнера: один для Zookeeper и один для Kafka:

```
CONTAINER ID   IMAGE                              COMMAND                  CREATED          STATUS          PORTS                    NAMES
1a2b3c4d5e6f   confluentinc/cp-kafka:7.0.1       "/etc/confluent/dock…"   10 seconds ago   Up 9 seconds    0.0.0.0:9092->9092/tcp   your_project_kafka_1
7g8h9i0j1k2l   confluentinc/cp-zookeeper:7.0.1   "/etc/confluent/dock…"   15 seconds ago   Up 14 seconds   0.0.0.0:2181->2181/tcp   your_project_zookeeper_1 
```

Для взаимодействия с Kafka можно использовать команды внутри контейнера Kafka. Подключитесь к контейнеру:

```
docker exec -it <container_id_kafka> /bin/sh 
```

Ссылка :
https://practicum.yandex.ru/learn/kafka/courses/83ca50bc-5dce-42cd-b629-997c971f1765/sprints/611653/topics/0c44bbdf-06e0-4ad4-9c5f-cf1969a67c2e/lessons/c6bd2abc-6684-44a3-b3e1-6bfb7554a80d/
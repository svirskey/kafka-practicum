# Установка и запуск Kafka кластера с использованием Kraft

### Описание работы приложения 

- Один инстанс приложения продюсера пишет 2 раза в секунду сгенерированное сообщение в топик для имитации работы.
- При поднятии композа топик автоматически создаётся с нужным количеством партиций и replication-factor
- Single consumer подписывается на топик и считывает каждое сообщение. Коммит автоматический
- Batch consumer подписывается на топик и пока не прочитает пачку в 10 сообщений не будет коммитить оффсет. Автокоммит выключен

### Разворачивание Kafka-кластера вместе с приложениями, имитирующими работу.

```
make build && make
```

### Вывод информации о созданном топике

```
docker exec -it infra-kafka-0-1 kafka-topics.sh --describe --topic user-msg  --bootstrap-server kafka-0:9092
```

### Просмотр логов 

- Просмотр логов инстансов продюсеров

```
docker logs infra-app-producer-1
docker logs infra-app-producer-2
```

- Просмотр логов инстансов консьюмеров, читающих логи по одному

```
docker logs infra-app-single-message-consumer-1
docker logs infra-app-single-message-consumer-1
```

- Просмотр логов инстансов консьюмеров, читающих логи пачкой

```
docker logs infra-app-batch-message-consumer-1
docker logs infra-app-batch-message-consumer-2
```


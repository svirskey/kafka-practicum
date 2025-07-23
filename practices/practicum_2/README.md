## Установка и запуск Kafka кластера с использованием Kraft

1. Как развернуть Kafka-кластер.

```
docker compose up -d
```

2. Вывод информации о созданном топике

```
kafka-topics.sh --describe --topic my-topic --bootstrap-server localhost:9092
```

3.


```
docker logs infra-app-producer-1
docker logs infra-app-producer-2
```
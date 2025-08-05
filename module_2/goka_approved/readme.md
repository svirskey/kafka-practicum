
## Блокировка

```
go run cmd/block-user/main.go -user user-0 -stream blocker-stream 
```

## Разблокировка

```
go run cmd/block-user/main.go -unblock -user user-0 -stream blocker-stream
```


## Создаем топики

```
docker exec -it infra-kafka-0-1 /opt/bitnami/kafka/bin/kafka-topics.sh --create --topic emitter2filter-stream --bootstrap-server localhost:9092 --partitions 3 --replication-factor 2
```

```
docker exec -it infra-kafka-0-1 /opt/bitnami/kafka/bin/kafka-topics.sh --create --topic user-like-table --bootstrap-server localhost:9092 --partitions 3 --replication-factor 2 --config cleanup.policy=compact
```


```
docker exec -it infra-kafka-0-1 /opt/bitnami/kafka/bin/kafka-topics.sh --create --topic filter2userprocessor-stream --bootstrap-server localhost:9092 --partitions 3 --replication-factor 2
```


```
docker exec -it infra-kafka-0-1 /opt/bitnami/kafka/bin/kafka-topics.sh --create --topic filter-table --bootstrap-server localhost:9092 --partitions 3 --replication-factor 2 --config cleanup.policy=compact
```


```
docker exec -it infra-kafka-0-1 /opt/bitnami/kafka/bin/kafka-topics.sh --create --topic blocker-stream --bootstrap-server localhost:9092 --partitions 3 --replication-factor 2 --config cleanup.policy=compact
```

```
docker exec -it infra-kafka-0-1 /opt/bitnami/kafka/bin/kafka-topics.sh --create --topic blocker-table --bootstrap-server localhost:9092 --partitions 3 --replication-factor 2 --config cleanup.policy=compact
```
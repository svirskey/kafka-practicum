
```
kafka-topics.sh --create --topic my-topic --partitions 5 --replication-factor 2 --bootstrap-server localhost:9092
```

Здесь --partitions ― количество партиций, а --replication-factor ― количество реплик для каждой партиции.

Запустить пример на Golang можно так: 

`go run main.go localhost:9094 example-topic 5 2` 

Этот код создаёт топик с именем example-topic, пятью партициями и фактором репликации, равным двум.
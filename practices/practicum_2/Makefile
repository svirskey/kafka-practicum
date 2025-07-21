
.PHONY: all
all:
	go build ./src/cmd/BatchMessageConsumer/main.go && \
	go build ./src/cmd/SingleMessageConsumer/main.go && \
	go build ./src/cmd/Producer/main.go && \
	docker compose up -d
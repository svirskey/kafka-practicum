
.PHONY: all
all:
	@docker compose -f ./infra/docker-compose.yml up -d

.PHONY: build
build:
	@docker compose -f ./infra/docker-compose.yml build

.PHONY: run
run: all

.PHONY: ps
ps:
	@docker compose -f ./infra/docker-compose.yml ps

.PHONY: down
down:
	@docker compose -f ./infra/docker-compose.yml down

.PHONY: rm
rm:
	@echo 'y' | docker compose -f ./infra/docker-compose.yml rm
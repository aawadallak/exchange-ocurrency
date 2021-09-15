#!make
include .env
export $(shell sed 's/=.*//' .env)

Build:
	@echo "Building..."
	go build -o out/main cmd/main.go

#define vars
DOCKER_NETWORK="$$NETWORK"
MIGRATIONS_PATH="$(shell pwd)/database/migrations"
DB_DSN="$$DB_TYPE://$$DB_USER:$$DB_PASS@$$DB_HOST:$$DB_PORT/$$DB_NAME"

migrations:
	docker run --rm -v $(MIGRATIONS_PATH):/migrations --network $(DOCKER_NETWORK) migrate/migrate -path=/migrations -database $(DB_DSN) up

clean:
	docker run --rm -it -v $(MIGRATIONS_PATH):/migrations --network $(DOCKER_NETWORK) migrate/migrate -path=/migrations -database $(DB_DSN) down

#Creates the migration file already named
create_migration:
	docker run --rm -v $(MIGRATIONS_PATH):/migrations --user $(shell id -u):$(shell id -g) \
					--network $(DOCKER_NETWORK) migrate/migrate create -ext sql -dir ./migrations -seq $(name)

v:
	docker run --rm -v $(MIGRATIONS_PATH):/migrations --user $(shell id -u):$(shell id -g) \
					--network $(DOCKER_NETWORK) migrate/migrate $(version)
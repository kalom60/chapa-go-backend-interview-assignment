include .env
.PHONY: build
build:
	go build -ldflags="-s" -o ./tmp/main ./cmd && chmod +x ./tmp/main
.PHONY: run
run:
	@echo "Loading environment variables from .env file"
	@set -o allexport; source ./load_env.sh; set +o allexport; \
	echo "Running Go application"; \
	go run ./cmd/main.go
.PHONY: air
watch:
	@echo "Loading environment variables from .env file"
	@set -o allexport; source ./load_env.sh; set +o allexport; \
	echo "Running air"; \
	air -c .air.toml

.PHONY: migrations/up
migrations/new:
	@echo 'Creating migration files for DB_URL'
	migrate create -seq -ext=.sql -dir=./db/migrations $(name)
.PHONY: migrations/up
migrations/up:
	@echo 'Running up migrations...'
	migrate -path ./db/migrations -database $(DB_URL) up

.PHONY: swagger
swagger:
	swag init -g cmd/main.go
migrate-create:
	- migrate create -ext sql -dir db/query -tz "UTC" $(args)

sqlc:
	- sqlc generate -f ./sqlc.yaml

up:
	@echo "Starting Docker images..."
	docker-compose -f docker-compose.yaml up --build -d
	@echo "Docker images started!"

down:
	@echo "Stopping docker compose..."
	docker-compose -f docker-compose.yaml down
	@echo "Done!"


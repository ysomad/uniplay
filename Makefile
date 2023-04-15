include .env
export

MIGRATE := migrate -path migrations -database "$(PG_URL)?sslmode=disable"

.PHONY: compose-up
compose-up:
	docker-compose up --build -d postgres jaeger prometheus grafana && docker compose logs -f

.PHONY: compose-all
compose-all:
	docker-compose up --build -d && docker compose logs -f

.PHONY: compose-down
compose-down:
	docker-compose down --remove-orphans

PHONY: run
run:
	go mod tidy && go mod download && \
	go run ./cmd/app

.PHONY: run-migrate
run-migrate:
	go mod tidy && go mod download && \
	go run -tags migrate ./cmd/app

.PHONY: lint
lint:
	golangci-lint run

.PHONY: test
test:
	go test -v -cover -race -count 1 ./internal/...

.PHONY: migrate-new
migrate-new:
	@read -p "Enter the name of the new migration: " name; \
	$(MIGRATE) create -ext sql -dir migrations $${name// /_}

.PHONY: migrate-up
migrate-up:
	@echo "Running all new database migrations..."
	@$(MIGRATE) up

.PHONY: migrate-down
migrate-down:
	@echo "Running all down database migrations..."
	@$(MIGRATE) down

.PHONY: migrate-drop
migrate-drop:
	@echo "Dropping everything in database..."
	@$(MIGRATE) drop

.PHONY: dry-run
dry-run: migrate-drop run-migrate

.PHONY: gen-swagger
gen-swagger:
	rm -rf ./internal/gen/swagger2/*
	swagger generate server \
	-t internal/gen/swagger2 \
	-f swagger2.json \
	-A uniplay \
	--exclude-main \
	--strict-responders \
	go mod tidy

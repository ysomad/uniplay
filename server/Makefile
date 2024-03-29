include .env
export

export GOOSE_DRIVER=postgres
export GOOSE_DBSTRING=${PG_URL}

.PHONY: compose-local
compose-local:
	docker-compose up --build -d postgres minio && docker-compose logs -f

.PHONY: compose-up
compose-up:
	docker-compose up --build -d && docker-compose logs -f

.PHONY: compose-down
compose-down:
	docker-compose down --remove-orphans

.PHONY: run
run:
	go mod tidy && go mod download && \
	go run ./cmd/app

.PHONY: run-migrate
run-migrate:
	go mod tidy && go mod download && \
	go run ./cmd/app -migrate

.PHONY: dry-run
dry-run: goose-reset run-migrate

.PHONY: lint
lint:
	golangci-lint run

.PHONY: test
test:
	go test -v -cover -race -count 1 ./internal/...

.PHONY: goose-new
goose-new:
	@read -p "Enter the name of the new migration: " name; \
	goose -s -dir migrations create $${name// /_} sql

.PHONY: goose-up
goose-up:
	@echo "Running all new database migrations..."
	goose -dir migrations validate
	goose -dir migrations up

.PHONY: goose-down
goose-down:
	@echo "Running all down database migrations..."
	goose -dir migrations down

.PHONY: goose-reset
goose-reset:
	@echo "Dropping everything in database..."
	goose -dir migrations reset

.PHONY: goose-status
goose-status:
	goose -dir migrations status

.PHONY:
generate:
	buf generate api/proto

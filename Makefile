include .env
export

MIGRATE := migrate -path migrations -database "$(PG_URL)?sslmode=disable"

.PHONY: compose-up
compose-up:
	docker-compose up --build -d postgres && docker compose logs -f

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

.PHONY: gen-oapi
gen-oapi:
	oapi-codegen \
	--package v1 \
	-generate chi-server,types \
	./openapi3.yaml > ./internal/gen/oapi/v1/v1.gen.go

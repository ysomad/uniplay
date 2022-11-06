PHONY: run
run:
	go mod tidy && go mod download && \
	go run ./cmd/app
# Step 1: Modules caching
FROM golang:1.20.4-alpine as modules
COPY go.mod go.sum /modules/
WORKDIR /modules
RUN go mod download

# Step 2: Builder
FROM golang:1.20.4-alpine as builder
COPY --from=modules /go/pkg /go/pkg
COPY . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -tags migrate -o /bin/app ./cmd/app

# Step 3: Final
FROM scratch
COPY --from=builder /app/configs /configs
COPY --from=builder /app/migrations /migrations
COPY --from=builder /bin/app /app
CMD ["/app", "-docker", "true"]
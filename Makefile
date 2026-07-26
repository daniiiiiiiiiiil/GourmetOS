.PHONY: run build test clean migrate

BINARY_NAME=gourmetos
PORT=8080

run:
	go run cmd/api/main.go

build:
	go build -o bin/$(BINARY_NAME) cmd/api/main.go

test:
	go test -v ./...

test-coverage:
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

clean:
	rm -rf bin/
	rm -f coverage.out

deps:
	go mod download
	go mod tidy

migrate:
	go run cmd/migrate/main.go

watch:
	air

migrate-up:
	migrate -path migrations -database ${DATABASE_URL} up

migrate-down:
	migrate -path migrations -database ${DATABASE_URL} down
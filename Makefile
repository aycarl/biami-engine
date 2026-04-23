.PHONY: run build test docker-build compose-up compose-down compose-test

run:
	go run ./cmd/api/main.go

build:
	mkdir -p bin
	go build -o bin/biami ./cmd/api/main.go

test:
	go test ./...

docker-build:
	docker build -t biami-api:latest .

compose-up:
	docker compose up --build api

compose-down:
	docker compose down

compose-test:
	docker compose --profile test run --rm test

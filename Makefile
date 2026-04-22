.PHONY: run build test docker-build

run:
	go run ./cmd/api/main.go

build:
	mkdir -p bin
	go build -o bin/biami ./cmd/api/main.go

test:
	go test ./...

docker-build:
	docker build -t biami-api:latest .


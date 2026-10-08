.PHONY: build run air test compose

build:
	go build -o app main.go

run:
	go run main.go

air:
	air

test:
	go test ./...

compose:
	docker compose up -d --build

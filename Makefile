.PHONY: build run air test

build:
	go build -o app main.go

run:
	go run main.go

air:
	air

test:
	go test ./...

# Бинарник по умолчанию
BINARY  := bin/game
MAIN    := ./cmd/game

.PHONY: build run test clean

build:
	@mkdir -p bin
	go build -o $(BINARY) $(MAIN)

run: build
	./$(BINARY)

test:
	go test ./...

clean:
	rm -rf bin

BIN := schema2json

.PHONY: all
all: build

.PHONY: build
build:
	go build -o $(BIN) ./cmd

.PHONY: test
test:
	go test ./...
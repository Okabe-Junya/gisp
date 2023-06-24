BINARY_NAME=gisp

.PHONY: fmt
fmt:
	@go fmt ./...

.PHONY: lint
lint:
	@golangci-lint run

.PHONY: test
test:
	@go test -v ./...

.PHONY: build
build:
	@go build -o bin/$(BINARY_NAME) cmd/$(BINARY_NAME)/main.go

.PHONY: run
run:
	@go run cmd/$(BINARY_NAME)/main.go

.PHONY: clean
clean:
	@rm -rf bin/$(BINARY_NAME)

.PHONY: install
install:
	@go install cmd/$(BINARY_NAME)/main.go

.PHONY: all
all: fmt lint test build

.PHONY: help
help:
	@echo "fmt: format code"
	@echo "lint: lint code"
	@echo "test: run tests"
	@echo "build: build binary"
	@echo "run: run binary"
	@echo "clean: remove binary"
	@echo "install: install binary"
	@echo "all: run all targets"
	@echo "help: show this help message"

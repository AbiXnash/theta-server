APP_NAME=theta-server
PKG=./...
PORT?=8080
THETA_DATA_DIR?=data

.PHONY: help fmt vet test build run clean check

help:
	@echo "Available targets:"
	@echo "  make fmt      - format Go files"
	@echo "  make vet      - run go vet"
	@echo "  make test     - run tests"
	@echo "  make build    - build binary"
	@echo "  make run      - run server with env vars"
	@echo "  make check    - fmt + vet + test + build"
	@echo "  make clean    - remove built binary"

fmt:
	gofmt -w $$(rg --files -g '*.go')

vet:
	go vet $(PKG)

test:
	go test $(PKG)

build:
	go build -o $(APP_NAME) .

run:
	PORT=$(PORT) THETA_API_KEY=$${THETA_API_KEY:?THETA_API_KEY is required} THETA_DATA_DIR=$(THETA_DATA_DIR) go run .

check: fmt vet test build

clean:
	rm -f $(APP_NAME)

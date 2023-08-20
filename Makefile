BINARY_NAME:=veego-cli
BUILD_DIR = $(PWD)/build

dev:
	air

build:
	go build -o ${BUILD_DIR}/$(BINARY_NAME) ./cmd/veego-cli

.PHONY: install
install:
	go install ./cmd/veego-cli
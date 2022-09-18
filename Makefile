.PHONY: all build lint fmt

default: all

all: build

build: lint
	go build

lint: 
	golangci-lint run

fmt:
	gofmt -w .

PACKAGE=github.com/dbrabera/slip
VERSION?=$(shell git describe --always --dirty)
LDFLAGS=-ldflags "-X $(PACKAGE)/internal.Version=$(VERSION)"

build:
	@go build $(LDFLAGS) -o bin/slip main.go

lint:
	@golangci-lint run ./...

test:
	@go test ./...

clean:
	@rm -rf bin

.PHONY: build lint test clean
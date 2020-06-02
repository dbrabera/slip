GIT_COMMIT=$(shell git rev-parse HEAD)
LDFLAGS=-ldflags "-X main.GitCommit=$(GIT_COMMIT)"

build:
	@go build $(LDFLAGS) -o bin/slip main.go

test:
	@go test ./...

clean:
	@rm -rf bin

.PHONY: build test clean
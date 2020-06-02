GIT_COMMIT=$(shell git rev-parse HEAD)
LDFLAGS=-ldflags "-X main.GitCommit=$(GIT_COMMIT)"

build:
	@go build $(LDFLAGS) -o bin/slip slip/main.go

clean:
	@rm -rf bin

deps:
	@godep save github.com/dbrabera/slip

test:
	@go test ./...

.PHONY: help build clean deps test
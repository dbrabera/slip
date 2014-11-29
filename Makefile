GIT_COMMIT=$(shell git rev-parse HEAD)
LDFLAGS=-ldflags "-X main.GitCommit $(GIT_COMMIT)"

default: test

deps:
	@godep save github.com/dbrabera/slip

build:
	@mkdir -p $(GOPATH)/src/github.com/dbrabera
	@ln -sf $(shell pwd) $(GOPATH)/src/github.com/dbrabera
	@go build $(LDFLAGS) -o bin/slip slip/main.go

test:
	@go test ./...

clean:
	@rm -f slip

.PHONY: default deps run test clean
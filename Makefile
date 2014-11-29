GIT_COMMIT=$(shell git rev-parse HEAD)
LDFLAGS=-ldflags "-X main.GitCommit $(GIT_COMMIT)"

help:
	@echo "Usage make [target]"
	@echo ""
	@echo "Targets:"
	@echo ""
	@echo "  build    Build the Slip interpreter"
	@echo "  clean    Remove built files"
	@echo "  deps     Save dependencies"
	@echo "  test     Run tests"
	@echo ""

build:
	@mkdir -p $(GOPATH)/src/github.com/dbrabera
	@ln -sf $(shell pwd) $(GOPATH)/src/github.com/dbrabera
	@go build $(LDFLAGS) -o bin/slip slip/main.go

clean:
	@rm -rf bin

deps:
	@godep save github.com/dbrabera/slip

test:
	@go test ./...

.PHONY: help build clean deps test
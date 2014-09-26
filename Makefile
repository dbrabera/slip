GIT_COMMIT=$(shell git rev-parse HEAD)

default: test

build: deps
	go build -ldflags "-X main.GitCommit $(GIT_COMMIT)" ./...

deps:
	go get -d -v ./...

run: build
	./slip

test: deps
	go test ./...

clean:
	rm -f slip

.PHONY: default build deps run test clean
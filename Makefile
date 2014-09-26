default: test

build: deps
	go build ./...

deps:
	go get -d -v ./...

run: build
	./slip

test: deps
	go test ./...

clean:
	rm -f slip

.PHONY: default build deps run test clean
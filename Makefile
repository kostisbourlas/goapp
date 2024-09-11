.DEFAULT_GOAL := goapp

.PHONY: all
all: clean goapp

.PHONY: goapp
goapp:
	mkdir -p bin
	go build -o bin ./...

.PHONY: clean
clean:
	go clean
	rm -f bin/*

run:	goapp
	bin/server

test:
	go test ./... -v

bench:
	go test ./... -bench=.

build-cli:
	mkdir -p bin
	go build -o bin/cli ./cmd/cli/

.PHONY: all clean test test-worker

all: worker coord client

worker:
	go build -o bin/worker ./cmd/worker

coord:
	go build -o bin/coord ./cmd/coord

client:
	go build -o bin/client ./cmd/client

clean:
	rm -f bin/*
	go clean -testcache

test:
	go test ./bagel -test.v
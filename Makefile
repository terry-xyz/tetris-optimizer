.PHONY: build run test fmt clean

BINARY := tetris-optimizer

build:
	go build -o $(BINARY) ./cmd

run: build
	./$(BINARY) $(ARGS)

test:
	go test ./...

test-v:
	go test -v ./...

fmt:
	gofmt -w .

clean:
	rm -f $(BINARY)

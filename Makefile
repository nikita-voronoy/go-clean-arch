.PHONY: build test clean run debug

GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test

BINARY_NAME=server
BINARY_PATH=cmd/server

all: test build

build:
	$(GOBUILD) -o $(BINARY_NAME) -v ./$(BINARY_PATH)

run: build
	./$(BINARY_NAME)

debug: build
	./$(BINARY_NAME) -debug

test:
	$(GOTEST) -v ./...

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
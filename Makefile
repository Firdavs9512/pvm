GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

all: build

build:
	$(GOBUILD) -o bin/core cmd/core/main.go

clean:
	$(GOCLEAN)
	rm -rf bin/core

run:
	@$(GOCMD) run cmd/core/main.go

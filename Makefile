BIN := goredirect
IMAGE := goredirect
TARGET := distr
VERSION := $(shell git describe --tags --always --dirty)
TAG := $(VERSION)
REGISTRY := arpabet
PWD := $(shell pwd)
NOW := $(shell date +"%m-%d-%Y")

all: build

version:
	@echo $(TAG)

build:
	rm -rf rsrc.syso
	go test -cover ./...
	go build -o $(BIN) -v -ldflags "-X main.Version=$(VERSION) -X main.Build=$(NOW)"

distr: build
	rm -rf $(TARGET)
	mkdir $(TARGET)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(TARGET)/$(BIN)_linux -v -ldflags "-X main.Version=$(VERSION) -X main.Build=$(NOW)"
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o $(TARGET)/$(BIN)_darwin -v -ldflags "-X main.Version=$(VERSION) -X main.Build=$(NOW)"
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o $(TARGET)/$(BIN).exe -v -ldflags "-X main.Version=$(VERSION) -X main.Build=$(NOW)"


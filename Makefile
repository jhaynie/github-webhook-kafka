.PHONY: all clean version

VERSION := 1.0.1
NAME := github-webhook-kafka
PKG := jhaynie/github-webhook-kafka

SHELL := /bin/bash
BASEDIR := $(shell echo $${PWD})
BUILD := $(shell git rev-parse HEAD | cut -c1-8)
SRC := $(shell find . -type f -name '*.go' -not -path './vendor/*' -not -path './.git/*')

L="-X=github.com/$(PKG)/main.Build=$(BUILD) -X=github.com/$(PKG)/main.Version=$(VERSION)"

all: version build

version:
	@echo "version: $(VERSION) (build: $(BUILD))"

clean:
	@rm -rf build

fmt:
	@gofmt -s -l -w $(SRC)

linux:
	@docker run --rm -v $(GOPATH):/go -w /go/src/github.com/$(PKG) golang:latest go build -ldflags $(L) -o build/linux/$(NAME)-linux-$(VERSION) main.go

alpine:
	@docker run --rm -v $(GOPATH):/go -w /go/src/github.com/$(PKG) jhaynie/golang-alpine go build -ldflags $(L) -o build/alpine/$(NAME)-alpine-$(VERSION) main.go

osx:
	@go build -ldflags $(L) -o build/osx/$(NAME)-osx-$(VERSION) main.go

build: version osx

docker: alpine
	@docker build -t $(PKG) .

docker-dev: linux
	@docker build -t $(PKG)-dev -f Dockerfile-dev .

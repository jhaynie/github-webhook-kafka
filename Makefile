.PHONY: all clean version

VERSION := 1.0.2
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

protoc:
	@docker run --rm -v $(BASEDIR):/app -w /app znly/protoc --go_out=. -I. *.proto

linux:
	@docker run --rm -v $(GOPATH):/go -w /go/src/github.com/$(PKG) golang:latest go build -ldflags $(L) -o build/linux/$(NAME)-linux-$(VERSION) $(SRC)

alpine:
	@docker run --rm -v $(GOPATH):/go -w /go/src/github.com/$(PKG) jhaynie/golang-alpine go build -ldflags $(L) -o build/alpine/$(NAME)-alpine-$(VERSION) $(SRC)

osx:
	@go build -ldflags $(L) -o build/osx/$(NAME)-osx-$(VERSION) $(SRC)

build: version osx

docker: alpine
	@docker build -t $(PKG) .

docker-dev: alpine
	@docker build -t $(PKG)-dev -f Dockerfile-dev .

dev: docker-dev
	@docker run --rm -it -p 8000:8000 -e GWK_GITHUB_SECRET=1234 $(PKG)-dev

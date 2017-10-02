
LDFLAGS = -X github.com/srossross/template/cmd.VERSION=$(shell echo $${CIRCLE_TAG:-?}) \
	-X github.com/srossross/template/cmd.BUILD_TIME=$(shell date -u +%Y-%m-%d)


build: ## build for any arch
	mkdir -p /tmp/commands
	$(eval FILE_PART := $(shell go env GOOS)-$(shell go env GOARCH))
	go build -ldflags "$(LDFLAGS)" -o ./template-$(FILE_PART) ./main.go
	tar -zcvf /tmp/commands/template-$(FILE_PART).tgz ./template-$(FILE_PART)


build-linux-amd64: export GOOS = linux
build-linux-amd64: export GOARCH = amd64
build-linux-amd64: build ## build for linux 64bit

build-linux-386: export GOOS = linux
build-linux-386: export GOARCH = 386
build-linux-386: build ## build for linux 32bit

build-darwin-amd64: export GOOS = darwin
build-darwin-amd64: export GOARCH = amd64
build-darwin-amd64: build ## build for darwin

all: build-darwin-amd64 build-linux-amd64

.PHONY: help

help: ## show this help and exit
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help

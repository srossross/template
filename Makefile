
LDFLAGS = -X github.com/srossross/template/cmd.VERSION=99 -X github.com/srossross/template/cmd.BUILD_TIME=$(shell date -u +%Y-%m-%d)

echoflags:
	echo $(LDFLAGS)
build: ## build for any arch
	mkdir -p /tmp/commands
	go build -ldflags "$(LDFLAGS)" -o ./template-${GOOS}-${GOARCH} ./main.go
	tar -zcvf /tmp/commands/template-${GOOS}-${GOARCH}.tgz ./template-${GOOS}-${GOARCH}


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

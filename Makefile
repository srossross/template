

build: ## build for any arch
	mkdir -p /tmp/commands
	go build -o ./template-${GOOS}-${GOARCH} ./main.go
	tar -zcvf /tmp/commands/template-${GOOS}-${GOARCH}.tgz ./template-${GOOS}-${GOARCH}


build-linux-amd64: export GOOS = linux
build-linux-amd64: export GOARCH = amd64
build-linux-amd64: build ## build for linux

build-darwin-amd64: export GOOS = darwin
build-darwin-amd64: export GOARCH = amd64
build-darwin-amd64: build ## build for darwin

all: build-darwin-amd64 build-linux-amd64


.PHONY: help

help: ## show this help and exit
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help

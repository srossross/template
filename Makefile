GOOS := $(shell go env GOHOSTOS)
GOARCH := $(shell go env GOHOSTARCH)

LDFLAGS := -X github.com/srossross/template/cmd.VERSION=$(shell echo $${CIRCLE_TAG:-?}) \
-X github.com/srossross/template/cmd.BUILD_TIME=$(shell date -u +%Y-%m-%d)

USERNAME := $(shell echo ${CIRCLE_PROJECT_USERNAME})
REPONAME := $(shell echo ${CIRCLE_PROJECT_REPONAME})
TAG := $(shell echo ${CIRCLE_TAG:-$(shell git describe --always)} )



build: ## build for any arch
	mkdir -p /tmp/commands

	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -ldflags "$(LDFLAGS)" -o ./template-$(GOOS)-$(GOARCH) ./main.go
	tar -zcvf /tmp/commands/template-$(GOOS)-$(GOARCH).tgz ./template-$(GOOS)-$(GOARCH)


all: build-darwin-amd64 build-linux-amd64

release: ## Create github release
	github-release release \
		--user $(USERNAME) \
		--repo $(REPONAME) \
		--tag $(TAG) \
		--name "Release $(TAG)" \
		--description "TODO: Description"

upload: ## Upload build artifacts to github

	github-release upload \
		--user $(USERNAME) \
		--repo $(REPONAME) \
		--tag $(TAG) \
		--name "template-linux-amd64.tgz" \
		--file /tmp/commands/template-linux-amd64.tgz

	github-release upload \
		--user $(USERNAME) \
		--repo $(REPONAME) \
		--tag $(TAG) \
		--name "template-linux-386.tgz" \
		--file /tmp/commands/template-linux-386.tgz

	github-release upload \
		--user $(USERNAME) \
		--repo $(REPONAME) \
		--tag $(TAG) \
		--name "template-darwin-amd64.tgz" \
		--file /tmp/commands/template-darwin-amd64.tgz

.PHONY: help

help: ## show this help and exit
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help

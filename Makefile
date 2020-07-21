WEBSITE ?= website

all: check-style test

## Runs golangci-lint and npm run lint.
.PHONY: check-style
check-style: golangci-lint
	@echo Checking for style guide compliance
	cd $(WEBSITE) && npm run lint

## Runs a local environment using docker-compose
.PHONY: start
start: check-style
	docker-compose up --build

## Run golangci-lint on codebase.
.PHONY: golangci-lint
golangci-lint:
	@if ! [ -x "$$(command -v golangci-lint)" ]; then \
		echo "golangci-lint is not installed. Please see https://golangci-lint.run/usage/install/#ci-installation for installation instructions."; \
		exit 1; \
	fi; \

	@echo Running golangci-lint
	golangci-lint run ./...

## Runs any lints and unit tests defined for the server and webapp, if they exist.
.PHONY: test
test:
	go test -v ./...
	cd $(WEBSITE) && npm test -- --watchAll=false;

# Help documentation à la https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
help:
	@cat Makefile | grep -v '\.PHONY' |  grep -v '\help:' | grep -B1 -E '^[a-zA-Z0-9_.-]+:.*' | sed -e "s/:.*//" | sed -e "s/^## //" |  grep -v '\-\-' | sed '1!G;h;$$!d' | awk 'NR%2{printf "\033[36m%-30s\033[0m",$$0;next;}1' | sort
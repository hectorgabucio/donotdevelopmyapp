WEBSITE ?= website

all: check-style test

## Runs golangci-lint and npm run lint.
.PHONY: check-style
check-style: golangci-lint
	@echo Checking for style guide compliance
	cd $(WEBSITE) && npm run lint


## Run golangci-lint on codebase.
.PHONY: golangci-lint
golangci-lint:
	@if ! [ -x "$$(command -v golangci-lint)" ]; then \
		echo "golangci-lint is not installed. Please see https://github.com/golangci/golangci-lint#install for installation instructions."; \
		exit 1; \
	fi; \

	@echo Running golangci-lint
	golangci-lint run ./...

## Runs any lints and unit tests defined for the server and webapp, if they exist.
.PHONY: test
test:
	go test -v ./...
	cd $(WEBSITE) && npm run fix && npm run test;
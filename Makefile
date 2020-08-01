WEBSITE ?= website

all: check-style test

## Create certificates to encrypt the gRPC connection.
.PHONY: cert
cert: 
	rm -rf tls/ca.cert tls/ca.key tls/ca.srl tls/service.csr tls/service.key tls/service.key tls/service.pem
	openssl genrsa -out ./tls/ca.key 4096
	openssl req -new -x509 -key ./tls/ca.key -sha256 -subj "/C=US/ST=NJ/O=CA, Inc." -days 365 -out ./tls/ca.cert
	openssl genrsa -out ./tls/service.key 4096
	openssl req -new -key ./tls/service.key -out ./tls/service.csr -config ./tls/certificate.conf
	openssl x509 -req -in ./tls/service.csr -CA ./tls/ca.cert -CAkey ./tls/ca.key -CAcreateserial \
		-out ./tls/service.pem -days 365 -sha256 -extfile ./tls/certificate.conf -extensions req_ext

## Runs golangci-lint and npm run lint.
.PHONY: check-style
check-style: golangci-lint
	@echo Checking for style guide compliance
	cd $(WEBSITE) && npm run lint && cd ..

## Runs a local environment using docker-compose
.PHONY: start
start: check-style
	docker-compose -f deployments/docker-compose.yml up --build

## Removes volumes and all containers
.PHONY: clean
clean: 
	docker-compose -f deployments/docker-compose.yml down --rmi local -v
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
test: mocks
	go test -v ./...
	cd $(WEBSITE) && npm test -- --watchAll=false;

## Runs tests and generates coverage files
.PHONY: cov
cov: check-style
	go test -cover -coverprofile=c.out -v ./...
	go tool cover -html=c.out -o coverage.html
	rm -rf c.out


## Autogenerates mocks
.PHONY: mocks
mocks: 
	mockery -all -output ./test/mocks

# Help documentation à la https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
help:
	@cat Makefile | grep -v '\.PHONY' |  grep -v '\help:' | grep -B1 -E '^[a-zA-Z0-9_.-]+:.*' | sed -e "s/:.*//" | sed -e "s/^## //" |  grep -v '\-\-' | sed '1!G;h;$$!d' | awk 'NR%2{printf "\033[36m%-30s\033[0m",$$0;next;}1' | sort
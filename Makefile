WEBSITE ?= website
CUR_DIR = $(CURDIR)

all: check-style test

## Create certificates to encrypt the gRPC connection.
.PHONY: cert
cert: 
	rm -rf tls/ca.cert tls/ca.key tls/ca.srl tls/service.csr tls/tls.key tls/tls.key tls/tls.crt
	openssl genrsa -out ./tls/ca.key 4096
	openssl req -new -x509 -key ./tls/ca.key -sha256 -subj "/C=US/ST=NJ/O=CA, Inc." -days 365 -out ./tls/ca.cert
	openssl genrsa -out ./tls/tls.key 4096
	openssl req -new -key ./tls/tls.key -out ./tls/service.csr -config ./tls/certificate.conf
	openssl x509 -req -in ./tls/service.csr -CA ./tls/ca.cert -CAkey ./tls/ca.key -CAcreateserial \
		-out ./tls/tls.crt -days 365 -sha256 -extfile ./tls/certificate.conf -extensions req_ext

## Runs golangci-lint and npm run lint.
.PHONY: check-style
check-style: golangci-lint
	@echo Checking for style guide compliance
	cd $(WEBSITE) && npm run lint && cd ..

## Runs a local environment using docker-compose
.PHONY: start
start: check-style
	docker-compose -f deployments/docker-compose.yml up

## Removes volumes and all containers
.PHONY: clean
clean: 
	docker-compose -f deployments/docker-compose.yml down --rmi local -v
## Run golangci-lint on codebase.
.PHONY: golangci-lint
golangci-lint:
	docker run --rm -v $(CUR_DIR):/app -w /app golangci/golangci-lint:latest golangci-lint run -v cmd/... internal/...

## Runs any lints and unit tests defined for the server, if they exist.
.PHONY: test-back
test-back:
	go test -race -v ./...

## Runs any lints and unit tests defined for the webapp if they exist
.PHONY: test-front
test-front:
	cd $(WEBSITE) && npm ci && npm test -- --watchAll=false;

## Runs test on backend and frontend
.PHONY: test
test: test-back test-front

## Runs tests and generates coverage files
.PHONY: cov
cov:
	go test -race -coverprofile=coverage.txt -covermode=atomic -v ./...

## Autogenerates mocks
.PHONY: mocks
mocks: 
	docker run --rm -v $(CUR_DIR):/src -w /src vektra/mockery:v2 --all --recursive --output ./test/mocks

# Help documentation à la https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
help:
	@cat Makefile | grep -v '\.PHONY' |  grep -v '\help:' | grep -B1 -E '^[a-zA-Z0-9_.-]+:.*' | sed -e "s/:.*//" | sed -e "s/^## //" |  grep -v '\-\-' | sed '1!G;h;$$!d' | awk 'NR%2{printf "\033[36m%-30s\033[0m",$$0;next;}1' | sort
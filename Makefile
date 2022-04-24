SHELL := /bin/bash
SOURCES := $(shell find . -name '*.go')
BINARY=service-hub
VERSION?=$(shell git describe --always --tags)
COMMIT=`git rev-parse HEAD`
ENV ?= dev

.PHONY: default
default: run

LDFLAGS=-ldflags "-X main.VERSION=${VERSION} -X main.COMMIT=${COMMIT} -X main.ENV=${ENV}"

.PHONY: root
root:
	@go run ${LDFLAGS} ./cmd/api/

.PHONY: migrate
migrate:
	@go run ${LDFLAGS} ./cmd/api/ migrate

.PHONY: run
run: check
	@go run ${LDFLAGS} ./cmd/api/ server

.PHONY: build
build:
	go build ${LDFLAGS} -o bin/${BINARY} ./cmd/api/

.PHONY: check
check: check-fmt
	@echo "[+] Done"

check-vet:
	@echo "[-] Checking go vet..."
	@go vet -v ./...

check-fmt:
	@echo "[-] Checking gofmt..."
	$(eval FMT_FILES := $(shell gofmt -l $(SOURCES)))
	@if [[ -n "$(FMT_FILES)" ]]; then \
		echo "[x] gofmt needs running on the following files:"; \
		echo "    - $(FMT_FILES)"; \
		echo "    [?] Use \`make fmt\` to reformat code."; \
		exit 1; \
	fi

.PHONY: fmt
fmt:
	@echo "[-] Formating gofmt..."
	@gofmt -w $(SOURCES)
	@echo "[+] Done"

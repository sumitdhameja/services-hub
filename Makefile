SHELL := /bin/bash

BINARY=service-hub
VERSION?=$(shell git describe --always --tags)
COMMIT=`git rev-parse HEAD`
ENV ?= dev

.PHONY: default
default: run

LDFLAGS=-ldflags "-X main.VERSION=${VERSION} -X main.COMMIT=${COMMIT} -X main.ENV=${ENV}"

.PHONY: run
run: check
	@go run ${LDFLAGS} ./cmd/api/ server

build:
	go build ${LDFLAGS} -o bin/${BINARY} ./cmd/api/

check: check-vet check-fmt
	@echo "[+] Done"

check-vet:
	@echo "[-] Checking go vet..."
	@go vet $$(go list ./... | grep -v vendor/) ; if [ $$? -eq 1 ]; then \
		echo "[x] Vet found suspicious constructs!"; \
		exit 1; \
	fi

check-fmt:
	@echo "[-] Checking gofmt..."
	$(eval FMT_FILES := $(shell gofmt -l $(SOURCES)))
	@if [[ -n "$(FMT_FILES)" ]]; then \
		echo "[x] gofmt needs running on the following files:"; \
		echo "    - $(FMT_FILES)"; \
		echo "    [?] Use \`make fmt\` to reformat code."; \
		exit 1; \
	fi

fmt:
	@echo "[-] Formating gofmt..."
	@gofmt -w $(SOURCES)
	@echo "[+] Done"

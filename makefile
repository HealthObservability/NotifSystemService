PROJECT_DIR = $(shell pwd)
PROJECT_BIN = $(PROJECT_DIR)/bin
$(shell [ -f bin ] || mkdir -p $(PROJECT_BIN))
PATH := $(PROJECT_BIN):$(PATH)

GOLANGCI_LINT = $(HOME)/bin/golangci-lint

.PHONY: lint
lint:
	### RUN GOLANGCI-LINT ###
	$(GOLANGCI_LINT) run ./... --config=golangci.yaml

.PHONY: lint-fast
lint-fast:
	$(GOLANGCI_LINT) run ./... --fast --config=golangci.yaml
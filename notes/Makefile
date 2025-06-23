SHELL := /bin/bash
SHELLFLAGS := -c
CURDIR := $(shell pwd)


.ONESHELL:
#======================================================================
# DEPS
#======================================================================
LOCAL_BIN := $(CURDIR)/bin
export GOBIN := $(LOCAL_BIN)

set-deps:
	@echo "Setting up project deps"
	@go install github.com/pressly/goose/v3/cmd/goose@latest
	@go install github.com/swaggo/swag/cmd/swag@latest

.PHONY: set-deps

#======================================================================
# BUILD
#======================================================================
build:
	@echo "Start project build"
	@docker-compose up -d

run:
	@echo "Start local app"
	@go run main.go

.PHONY: build run

#======================================================================
# LOAD TESTING
#======================================================================
testOneThread:
ifndef ENDPOINT
	@echo "Testing default endpoint: $(BASE_URL)/notes/?limit=2"
	@wrk -t 1 -c 1 -d 300s --latency "$(BASE_URL)/notes/?limit=2"
else
	@echo "Testing custom endpoint: $(BASE_URL)$(ENDPOINT)"
	@wrk -t 1 -c 1 -d 300s --latency "$(BASE_URL)$(ENDPOINT)"
endif

testMultiThread:
ifndef ENDPOINT
	@echo "Testing default endpoint: $(BASE_URL)/notes/?limit=2"
	@wrk -t 10 -c 400 -d 5m --latency "$(BASE_URL)/notes/?limit=2"
else
	@echo "Testing custom endpoint: $(BASE_URL)$(ENDPOINT)"
	@wrk -t 10 -c 400 -d 5m --latency "$(BASE_URL)$(ENDPOINT)"
endif


.PHONY: testMultiThread testOneThread 

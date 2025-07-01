SHELL := /bin/bash
SHELLFLAGS := -c
CURDIR := $(shell pwd)



# ENV
DEPLOYMENTS_PACKAGE := $(CURDIR)/deployments
ENV := $(DEPLOYMENTS_PACKAGE)/.env_local
TEST_ENV := $(DEPLOYMENTS_PACKAGE)/.test_env
.ONESHELL: 


define local-env-script
	set -a
	source $(ENV)
	set +a 
endef


set-local-env:
	@$(local-env-script)



# DEPS
LOCAL_BIN := $(CURDIR)/bin
export GOBIN := $(LOCAL_BIN)

set-deps:
	@echo "Setting up project deps"
	@go install github.com/swaggo/swag/cmd/swag@latest
	@go install github.com/rubenv/sql-migrate/...@latest
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest



.PHONY: set-deps

# LOCAL MIGRATIONS
MIGRATOR := $(LOCAL_BIN)/sql-migrate
DB_COMPOSE_FILE := ./deployments/docker_compose_db_only.yml
MIGRATIONS_DIR := ./internal/infra/postgres/schema


start-db-local: set-local-env
	@echo "Starting local database..."
	@$(local-env-script)
	@docker compose -f $(DB_COMPOSE_FILE) up -d

stop-db-local:
	@echo "Stopping local database..."
	@docker compose -f $(DB_COMPOSE_FILE) down
	@echo "Database stopped."

up: set-local-env
	@echo "Applying migrations..."
	@cd $(MIGRATIONS_DIR) && $(MIGRATOR) up -config=dbconfig.yaml -env=local

down:
	@echo "Rolling back migrations..."
	@cd $(MIGRATIONS_DIR) && $(MIGRATOR) down -config=dbconfig.yaml -env=local

status:
	@echo "Checking migration status..."
	@cd $(MIGRATIONS_DIR) && $(MIGRATOR) status -config=dbconfig.yaml -env=local

.PHONY: start-db-local stop-db-local up down status

# BUILD
build:
	@echo "Start project build"
	@docker-compose up -d

run:
	@echo "Start local app"
	@go run main.go

.PHONY: build run

# LOAD TESTING (NOT WORKING NOW)
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

#LINT
lint: 
	@echo "Start linting"
	@golangci-lint run

#AD-HOC
elastic-up:
	@docker compose -f $(DEPLOYMENTS_PACKAGE)/docker-compose-elastic.yaml up -d
elastic-down:
	@docker compose -f $(DEPLOYMENTS_PACKAGE)/docker-compose-elastic.yaml down -d
redis-up:
	@docker compose -f $(DEPLOYMENTS_PACKAGE)/docker-compose-redis.yaml up -d

redis-down:
	@docker compose -f $(DEPLOYMENTS_PACKAGE)/docker-compose-redis.yaml down -d
# From:	Paul Smith
# Subject:	Re: multiple .PHONY or single .PHONY
# Date:	Thu, 18 Feb 2010 10:23:19 -0500
# On Thu, 2010-02-18 at 09:37 -0500, Martin d'Anjou wrote:
# > Hello,
# > 
# > Is it okay to declare .PHONY multiple times, scattered throughout the 
# > makefile like this:
# > 
# > .PHONY: all
# > .PHONY: clean
# > .PHONY: other
# > 
# > Or must all .PHONY targets be consolidated into a single .PHONY 
# > declaration like this:
# > 
# > .PHONY:all clean other

# You can have as many .PHONY prerequisite definitions as you like;
# they're cumulative, just like any other target.

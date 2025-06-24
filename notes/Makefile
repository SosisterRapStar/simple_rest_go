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
	@go install github.com/swaggo/swag/cmd/swag@latest
	@go install github.com/rubenv/sql-migrate/...@latest


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
# LOAD TESTING (NOT WORKING NOW)
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

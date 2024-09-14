# make/rules.mk

.DEFAULT_GOAL := help


PROJECT_NAME := $(shell basename $(PWD))
CUR_FILE := "$(PWD)/$(RULES_FILE)"
BIN := $(realpath $(shell dirname "$(CUR_FILE)")/../bin)
HEY := $(PWD)
COVER_PROFILE=coverage.out


## help: this command
help:
	@$(BIN)/help.sh $(PROJECT_NAME) $(CUR_FILE)
.PHONY: help

## config: show docker compose yaml config
config:
	@# echo ${PROJECT_NAME}
	@echo ${CURDIR}
	@echo ""
	@echo ${CUR_FILE}
	@echo ""
	@ls ${BIN}
	@docker compose config
.PHONY: config


## run: go run app.go
run:
	@PORT=8080 go run ./cmd/server/main.go
.PHONY: local-up

## up: docker compose up
up:
	@docker compose up --build -d
.PHONY: up

## down: docker compose down
down:
	@docker compose down
.PHONY: down

## lint: linting by golangci
lint:
	@golangci-lint run
.PHONY: lint

## test: run unit tests
test:
	@$(BIN)/run-unit-tests.sh
.PHONY: test

## test-all: run unit tests and integration tests
test-all:
	@$(BIN)/run-all-tests.sh
.PHONY: test-all

$(COVER_PROFILE): $(shell find . -type f -print | grep -v vendor | grep "\.go")
	@go test -cover -coverprofile ./$(COVER_PROFILE).tmp ./...
	@cat ./$(COVER_PROFILE).tmp | grep -v '.pb.go' | grep -v 'mock_' > ./$(COVER_PROFILE)
	@rm ./$(COVER_PROFILE).tmp

## cover: show test coverage
cover: $(COVER_PROFILE)
	@echo ""
	@go tool cover -func ./$(COVER_PROFILE)
.PHONY: cover

## cover-html: html representation of coverage
cover-html: $(COVER_PROFILE)
	@go tool cover -html=./$(COVER_PROFILE)
.PHONY: cover-html

## clean: cleaning
clean:
	@rm ./$(COVER_PROFILE)
.PHONY: clean



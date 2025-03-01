SHELL := /bin/sh
.DEFAULT_GOAL := run

include .env
export

command?=run
logging_level?=info


.PHONY: dev
dev:
	@GOEXPERIMENT=noregabi air -- -logging_level=$(logging_level)

.PHONY: run
run:
	@GOEXPERIMENT=noregabi go run ./cmd/worker-receiver/main.go -logging_level=$(logging_level)

.PHONY: debug
debug:
	@$(shell which go) build -o ./build/remote_writer ./cmd/worker-receiver/main.go
	@$(shell go env GOPATH)/bin/dlv \
		--listen=:2345 \
		--api-version=2 \
		--headless=true \
		--log=true \
		--accept-multiclient \
			exec ./build/remote_writer -logging_level=$(logging_level)

.PHONY: test
test:
	@CGO_ENABLED=0
	@GOOS=linux
	@go test -v ./... -coverprofile ./cover.out
	@go tool cover -html=cover.out


.PHONY: up
up: down
	@export PORT=${PORT}; \
		envsubst < devops/docker/volumes/prometheus/prometheus.yaml > devops/docker/volumes/prometheus/prometheus.yaml.baked
	@export ENV_FILE=../../$(env); \
	 export MAKE_COMMAND=$(command); \
		docker compose --env-file $(env) --file devops/docker/docker-compose.yaml up \
			--build \
			--detach

.PHONY: down
down:
	@export ENV_FILE=../../$(env); \
		docker compose --env-file $(env) --file devops/docker/docker-compose.yaml down \
			--volumes

.PHONY: logs
logs:
	@export ENV_FILE=../../$(env); \
		docker compose --env-file $(env) --file devops/docker/docker-compose.yaml logs $(app) \
			--follow
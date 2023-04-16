COMPOSE_FILE := docker-compose.yml

.PHONY: build
build:
	docker-compose -f $(COMPOSE_FILE) build
.PHONY: up
# Start the Docker Compose services in the background
up:
	docker-compose -f $(COMPOSE_FILE) up -d
.PHONY: tools
tools:
	go install -C internal/tools \
		github.com/golangci/golangci-lint/cmd/golangci-lint
.PHONY: lint
lint:
	golangci-lint run --build-tags integration --disable-all -E revive -E staticcheck -E structcheck -E unused -E gocritic -E gocyclo -E gofmt -E misspell -E stylecheck -E unconvert -E unparam

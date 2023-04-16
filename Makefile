COMPOSE_FILE := docker-compose.yml

test: generate
	go test -v ./...
generate:
	go generate ./...
.PHONY: build
build:
	docker-compose -f $(COMPOSE_FILE) build
.PHONY: up
up:
	docker-compose -f $(COMPOSE_FILE) up -d
.PHONY: down
down:
	docker-compose down --volumes --remove-orphans
.PHONY: tools
tools:
	go install -C internal/tools \
		github.com/golangci/golangci-lint/cmd/golangci-lint \
		github.com/matryer/moq
.PHONY: lint
lint:
	golangci-lint run --build-tags integration --disable-all -E revive -E staticcheck -E structcheck -E unused -E gocritic -E gocyclo -E gofmt -E misspell -E stylecheck -E unconvert -E unparam

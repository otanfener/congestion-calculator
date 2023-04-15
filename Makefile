COMPOSE_FILE := docker-compose.yml

build:
	docker-compose -f $(COMPOSE_FILE) build

# Start the Docker Compose services in the background
up:
	docker-compose -f $(COMPOSE_FILE) up -d

.PHONY: import
import:
	sh data/import.sh "$(CURDIR)/data/import.json"

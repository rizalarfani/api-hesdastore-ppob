GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
CYAN   := $(shell tput -Txterm setaf 6)
RESET  := $(shell tput -Txterm sgr0)

# Live Reload
watch-prepare:
	curl sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh

# run dev
watch:
	bin/air

build:
	go build -o ppob-service


# Docker
docker-compose:
	docker-compose up -d --build --force-recreate

docker-build: ## Build the Docker image with a specified tag
	@echo "$(CYAN)Building Docker image...$(RESET)"
	@if [ -z "$(tag)" ]; then \
		echo "$(YELLOW)Error: Please specify the 'tag' parameter, e.g., make docker-build tag=1.0.0$(RESET)"; \
		exit 1; \
	fi
	docker build --platform linux/amd64 -t rijalarfani/ppob-service-hesda:$(tag) .
	@echo "$(GREEN)Docker image built with tag '$(tag)'$(RESET)"

docker-push:
	@echo "$(CYAN)Building Docker image...$(RESET)"
	@if [ -z "$(tag)" ]; then \
		echo "$(YELLOW)Error: Please specify the 'tag' parameter, e.g., make docker-push tag=1.0.0$(RESET)"; \
		exit 1; \
	fi
	docker push rijalarfani/ppob-service-hesda:$(tag)
	@echo "$(GREEN)Docker image built with tag '$(tag)'$(RESET)"
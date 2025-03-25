.DEFAULT_GOAL := help

GREEN=\033[0;92m
YELLOW=\033[0;93m
VIOLET=\033[0;96m
NC=\033[0m

ENV_FILE = .env
ENV_EXAMPLE = env.example
GO = go
SHELL := /bin/bash

DES = out/api
.PHONY: help
help: ## Print help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: build
build: ## Build the api
	@printf "➜ Build ${GREEN}application${NC}\n"
	@$(GO) build -o $(DES)
	@printf "➜ Fixing file permission for ${GREEN}$(PWD)/$(DES)${NC}\n"
	@chmod +x "$(PWD)/$(DES)";

.PHONY: build-env
build-env: ## Generate .env from env.example (if it doesn't exist)
	@if [ ! -f "$(ENV_FILE)" ]; then \
		@cp "$(ENV_EXAMPLE)" "$(ENV_FILE)"; \
		@printf "➜ ${GREEN}$(ENV_FILE)${NC} generated from ${GREEN}$(ENV_EXAMPLE)${NC}\n"; \
	else \
		echo "$(ENV_FILE) already exists. Run 'make force-build-env' to overwrite."; \
	fi

.PHONY: force-build-env
force-build-env: ## Force regenerate .env file
	@cp "$(ENV_EXAMPLE)" "$(ENV_FILE)"
	@printf "➜ $(ENV_FILE) regenerated ${GREEN}(overwritten if existing)${NC}\n"

.PHONY: serve
serve: ## Launch the api
	@printf "${GREEN}➜ Launching the api${NC}\n"
	$(PWD)/$(DES)

.PHONY: dev
dev: build-env build serve ## Build environment, compile and launch the api
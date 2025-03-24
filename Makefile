.DEFAULT_GOAL := help

ENV_FILE = .env
ENV_EXAMPLE = env.example
GO = go
BASH = bash

.PHONY: help
help: ## Print help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: build
build: ## Build the api
	$(GO) build -o out/api

.PHONY: build-env
build-env: ## Generate .env from env.example (if it doesn't exist)
	@if [ ! -f "$(ENV_FILE)" ]; then \
		cp "$(ENV_EXAMPLE)" "$(ENV_FILE)"; \
		echo "$(ENV_FILE) generated from $(ENV_EXAMPLE)"; \
	else \
		echo "$(ENV_FILE) already exists. Run 'make force-build-env' to overwrite."; \
	fi

.PHONY: force-build-env
force-build-env: ## Force regenerate .env file
	@cp "$(ENV_EXAMPLE)" "$(ENV_FILE)"
	@echo "$(ENV_FILE) regenerated (overwritten if existing)"

.PHONY: serve
serve: ## Launch the api
	$(BASH) ./out/api

.PHONY: dev
dev: build-env build serve ## Build environment, compile and launch the api
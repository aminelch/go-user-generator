.DEFAULT_GOAL := help

.PHONY: help
help: ## Print help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: build
build: ## Build the api
	$(go) go build -o out/api

.PHONY: serve
serve: ## Launch the api
	$(bash) ./out/api
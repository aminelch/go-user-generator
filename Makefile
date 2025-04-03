.DEFAULT_GOAL := help

GREEN=\033[0;92m
NC=\033[0m

ENV_FILE = .env
ENV_DIST = ./etc/env.dist
GO = go
SHELL := /bin/bash
DES = out/api

.PHONY: help
help: ## Print help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

version: VERSION ## Show project version
	@cat VERSION
	@printf "\n"

.PHONY: build
build: ## Build the api
	@printf "➜ Build ${GREEN}application${NC}\n"
	@$(GO) build -o $(DES)
	@printf "➜ Fixing file permission for ${GREEN}$(PWD)/$(DES)${NC}\n"
	@chmod +x "$(PWD)/$(DES)";

.PHONY: build-env
build-env: ## Generate .env from env.dist (if it doesn't exist)
	@if [ ! -f "$(ENV_FILE)" ]; then \
		@cp "$(ENV_DIST)" "$(ENV_FILE)"; \
		@printf "➜ ${GREEN}$(ENV_FILE)${NC} generated from ${GREEN}$(ENV_DIST)${NC}\n"; \
	else \
		printf "${GREEN}$(ENV_FILE)${NC} already exists. Run 'make ${GREEN} force-build-env'${NC} to overwrite.\n"; \
	fi

.PHONY: force-build-env
force-build-env: ## Force regenerate .env file
	@cp "$(ENV_DIST)" "$(ENV_FILE)"
	@printf "➜ $(ENV_FILE) regenerated ${GREEN}(overwritten if existing)${NC}\n"

.PHONY: serve
serve: ## Launch the api
	@printf "${GREEN}➜ Launching the api${NC}\n"
	$(PWD)/$(DES)

.PHONY: dev
dev: build-env build serve ## Build environment, compile and launch the api

.PHONY: tag
tag: ## Make a new tag
	$(if $(TAG),,$(error TAG is not defined. Pass via "make tag TAG=4.2.1"))
	@printf "➜ Tagging ${GREEN}$(TAG)${NC}\n"
	@$(SHELL) ./etc/tag.sh $(TAG)
	@printf "➜ Update project version to ${GREEN}$(TAG)${NC}\n"

.PHONY: tag-and-push
tag-and-push: tag ## Make a new tag and push it
	@printf "➜ Create a specific commit to mentioned tag\n"
	git add -A
	git commit -m '$(TAG) release'
	@printf "➜ Signing tag\n"
	git tag --sign 'v$(TAG)' -m'Version $(TAG)'
	@printf "➜ Push tag to remote \n"
	push git push origin 'v$(TAG)'
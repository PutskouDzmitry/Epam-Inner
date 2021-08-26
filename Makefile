REPO                  := github.com/PutskouDzmitry/Epam-Inner

PHONY: help
help: ## makefile targets description
	@echo "Usage:"
	@egrep '^[a-zA-Z_-]+:.*##' $(MAKEFILE_LIST) | sed -e 's/:.*##/#-/' | column -t -s "#"

.PHONY: fmt
fmt: ## automatically formats Go source code
	@echo "Running 'go fmt ...'"
	@go fmt -x "$(REPO)/..."

.PHONY: init
init: fmt
	@go mod init
	@go mod tidy
	@go mod vendor

.PHONY: image
image:  ## build image from Dockerfile ./docker/server/Dockerfile
	@docker build -t inner-app:latest .

.PHONY: initDb
initDb: image
	@docker build -f ./docker/db/Dockerfile -t db .

.PHONE: initDbUp
initDbUp: initDb
	@docker build -f ./docker/dbUp/Dockerfile -t dbup .

.PHONY: up
up : initDbUp ## up docker compose
	@docker-compose up



startK6:
	k6 run k6/script.js

.PHONY: down
down:
	@docker-compose down




.PHONY: build-api
client: ## Build and run api.
	go build -race -ldflags "-s -w" -o bin/api src/main.go
	bin/api

.PHONY: run-dev
run-dev:
	bash -c "export ENV=local && nodemon --exec go run src/main.go --signal SIGTERM"

.PHONY: push-tag
push-tag:
	bash -c "git tag -fa dev-v1.0.1 && git push --force origin dev-v1.0.1"
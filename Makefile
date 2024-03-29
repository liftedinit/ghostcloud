#### HELP ####

help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
.PHONY: help

#### LINT ####

golangci_version=v1.55.1

lint-install:
	@echo "--> Installing golangci-lint $(golangci_version)"
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@$(golangci_version)
	@echo "--> Installing golangci-lint $(golangci_version) complete"

lint: ## Run linter (golangci-lint)
	@echo "--> Running linter"
	$(MAKE) lint-install
	@golangci-lint run ./x/...

lint-fix:
	@echo "--> Running linter"
	$(MAKE) lint-install
	@golangci-lint run ./x/... --fix

.PHONY: lint lint-fix

#### FORMAT ####

goimports_version=latest

format-install:
	@echo "--> Installing goimports $(goimports_version)"
	@go install golang.org/x/tools/cmd/goimports@$(goimports_version)
	@echo "--> Installing goimports $(goimports_version) complete"

format: ## Run formatter (goimports)
	@echo "--> Running goimports"
	$(MAKE) format-install
	@find . -name '*.go' -exec goimports -w -local github.com/cosmos/cosmos-sdk,cosmossdk.io,github.com/cometbft,github.com/cosmos.ibc-go,github.com/liftedinit/ghostcloud  {} \;

#### COVERAGE ####

coverage: ## Run coverage report
	@echo "--> Running coverage"
	@go test -race -cpu=$$(nproc) -covermode=atomic -coverprofile=coverage.out $$(go list ./x/...) > /dev/null 2>&1
	@echo "--> Running coverage filter"
	@./scripts/filter-coverage.sh
	@echo "--> Running coverage report"
	@go tool cover -func=coverage-filtered.out
	@echo "--> Running coverage html"
	@go tool cover -html=coverage-filtered.out -o coverage.html
	@echo "--> Coverage report available at coverage.html"
	@echo "--> Cleaning up coverage files"
	@rm coverage.out
	@echo "--> Running coverage complete"

.PHONY: coverage

#### TEST ####

test: ## Run tests
	@echo "--> Running tests"
	@go test -race -cpu=$$(nproc) $$(go list ./x/...)

.PHONY: test

#### VET ####

vet: ## Run go vet
	@echo "--> Running go vet"
	@go vet ./...

.PHONY: vet

#### NANCY ###

nancy: ## Run nancy
	@echo "--> Running nancy (requires docker)"
	@command -v docker >/dev/null 2>&1 || { echo >&2 "nancy requires docker but it's not installed.  Aborting."; exit 1; }
	go list -json -deps ./... | docker run --rm -i sonatypecommunity/nancy:latest sleuth

.PHONY: nancy


#### GOVULNCHECK ####
govulncheck_version=latest

govulncheck-install:
	@echo "--> Installing govulncheck $(govulncheck_version)"
	@go install golang.org/x/vuln/cmd/govulncheck@$(govulncheck_version)
	@echo "--> Installing govulncheck $(govulncheck_version) complete"

govulncheck: ## Run govulncheck
	@echo "--> Running govulncheck"
	$(MAKE) govulncheck-install
	@govulncheck ./...

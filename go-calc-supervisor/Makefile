# Go Calculator Supervisor Implementation

.PHONY: all test build run lint coverage clean help

all: test lint ## Run tests and linting

test: ## Run TDD tests
	@echo "🧪 Running Go tests..."
	go test -v

build: ## Build the demo application
	@echo "🔨 Building Go calc supervisor..."
	mkdir -p bin
	go build -o bin/calc-supervisor .

run: ## Run the demo application
	@echo "🚀 Running Go calc supervisor..."
	go run .

lint: ## Run Go linting and formatting
	@echo "🔍 Linting Go code..."
	go fmt ./...
	go vet ./...
	@if ! command -v golangci-lint >/dev/null 2>&1; then \
		echo "Installing golangci-lint..."; \
		go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest; \
	fi
	~/go/bin/golangci-lint run

coverage: ## Generate test coverage report
	@echo "📊 Generating Go coverage report..."
	mkdir -p coverage
	go test -coverprofile=coverage/coverage.out
	go tool cover -html=coverage/coverage.out -o coverage/coverage.html
	@echo "Coverage report generated: coverage/coverage.html"

clean: ## Clean build artifacts
	@echo "🧹 Cleaning Go artifacts..."
	go clean
	rm -rf bin/ coverage/
	rm -f go.sum

help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "%-15s %s\n", $$1, $$2}'
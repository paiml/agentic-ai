# Go Simple Actor Implementation

.PHONY: test build run lint coverage clean help

test: ## Run TDD tests
	@echo "🧪 Running Go tests..."
	go test -v -run="TestThreeRoundPingPong|TestMessageOrdering|TestDeterministicBehavior|TestPerformance"

build: ## Build the demo application
	@echo "🔨 Building Go demo..."
	mkdir -p bin
	go build -o bin/ping-pong .

run: ## Run the demo application
	@echo "🚀 Running Go demo..."
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
	go test -coverprofile=coverage/coverage.out -run="TestThreeRoundPingPong|TestMessageOrdering|TestDeterministicBehavior|TestPerformance"
	go tool cover -html=coverage/coverage.out -o coverage/coverage.html
	@echo "Coverage report generated: coverage/coverage.html"

clean: ## Clean build artifacts
	@echo "🧹 Cleaning Go artifacts..."
	go clean
	rm -rf bin/

help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "%-15s %s\n", $$1, $$2}'
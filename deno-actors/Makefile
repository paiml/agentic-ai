# Deno Simple Actor Implementation

.PHONY: test build run lint fmt coverage clean help

test: ## Run TDD tests
	@echo "🧪 Running Deno tests..."
	deno test simple_test.ts

build: ## Build the demo application
	@echo "🔨 Building Deno demo..."
	mkdir -p bin
	deno compile --output=bin/ping-pong main.ts

run: ## Run the demo application
	@echo "🚀 Running Deno demo..."
	deno run main.ts

lint: ## Run Deno linting and formatting
	@echo "🔍 Linting Deno code..."
	deno fmt --check
	deno lint

fmt: ## Format Deno code
	@echo "🎨 Formatting Deno code..."
	deno fmt .

coverage: ## Generate test coverage report
	@echo "📊 Generating Deno coverage report..."
	mkdir -p coverage
	deno test --coverage=coverage simple_test.ts
	deno coverage coverage --html
	@echo "Coverage report generated: coverage/html/index.html"

clean: ## Clean build artifacts
	@echo "🧹 Cleaning Deno artifacts..."
	rm -rf bin/ coverage/

help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "%-15s %s\n", $$1, $$2}'
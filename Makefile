# Simple Actor Implementations

.PHONY: all test build run lint coverage clean help

all: test lint ## Run all tests and linting for all implementations
	@echo "✅ All tests and linting completed successfully!"

test: ## Run all implementation tests
	@echo "🧪 Testing all implementations..."
	@$(MAKE) -C go-actors test
	@$(MAKE) -C go-calc-supervisor test
	@$(MAKE) -C rust-actors test
	@$(MAKE) -C deno-actors test
	@$(MAKE) -C ruchy-actors test
	@echo "✅ All available tests completed!"

build: ## Build all implementations
	@echo "🔨 Building all implementations..."
	@$(MAKE) -C go-actors build
	@$(MAKE) -C go-calc-supervisor build
	@$(MAKE) -C rust-actors build
	@$(MAKE) -C deno-actors build
	@$(MAKE) -C ruchy-actors build
	@echo "✅ All available builds completed!"

run: ## Run all implementation demos
	@echo "🚀 Running all demos..."
	@echo "--- Go Demo ---"
	@$(MAKE) -C go-actors run
	@echo "--- Rust Demo ---"
	@$(MAKE) -C rust-actors run
	@echo "--- Deno Demo ---"
	@$(MAKE) -C deno-actors run
	@echo "--- Ruchy Demo ---"
	@$(MAKE) -C ruchy-actors run

lint: ## Run linting and formatting for all implementations
	@echo "🔍 Linting all implementations..."
	@$(MAKE) -C go-actors lint
	@$(MAKE) -C go-calc-supervisor lint
	@$(MAKE) -C rust-actors lint
	@$(MAKE) -C deno-actors lint
	@$(MAKE) -C ruchy-actors lint
	@echo "✅ All available linting completed!"

coverage: ## Generate test coverage reports for all implementations
	@echo "📊 Generating coverage reports for all implementations..."
	@$(MAKE) -C go-actors coverage
	@$(MAKE) -C go-calc-supervisor coverage
	@$(MAKE) -C rust-actors coverage
	@$(MAKE) -C deno-actors coverage
	@$(MAKE) -C ruchy-actors coverage
	@echo "✅ All available coverage reports generated!"

clean: ## Clean all build artifacts
	@echo "🧹 Cleaning all implementations..."
	@$(MAKE) -C go-actors clean
	@$(MAKE) -C go-calc-supervisor clean
	@$(MAKE) -C rust-actors clean
	@$(MAKE) -C deno-actors clean
	@$(MAKE) -C ruchy-actors clean
	@echo "✅ All clean!"

help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "%-15s %s\n", $$1, $$2}'
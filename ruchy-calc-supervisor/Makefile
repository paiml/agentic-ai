# Ruchy Calculator Supervisor - ✅ WORKING (v3.63.0)

.PHONY: test build run lint coverage check clean all help

test: ## Run TDD tests
	@echo "🧪 Running Ruchy tests..."
	@ruchy calculator.ruchy
	@echo "✅ All tests passed"

run: ## Run the calculator demo
	@echo "🚀 Running Ruchy Calculator Supervisor..."
	@ruchy calculator.ruchy

build: ## Validate demo syntax
	@echo "🔨 Validating calculator syntax..."
	@ruchy check calculator.ruchy
	@echo "✅ Validation successful"

lint: ## Run linting and formatting
	@echo "🔍 Linting Ruchy code..."
	@ruchy check calculator.ruchy
	@echo "✅ Linting successful"

coverage: ## Generate test coverage report
	@echo "📊 Generating Ruchy coverage report..."
	@mkdir -p coverage
	@ruchy coverage --format html calculator.ruchy > coverage/calculator.html
	@echo "Coverage report generated: coverage/calculator.html"

check: ## Quick syntax check
	@echo "⚡ Quick syntax check..."
	@ruchy check calculator.ruchy

clean: ## Clean temporary files
	@echo "🧹 Cleaning temporary files..."
	@find . -name "*.tmp" -delete 2>/dev/null || true
	@echo "🧹 Clean completed"

all: test run ## Run tests then demo
	@echo "✅ All tasks completed"

help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "%-15s %s\n", $$1, $$2}'

.DEFAULT_GOAL := help

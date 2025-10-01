# Ruchy Calculator Supervisor - ✅ WORKING (v3.62.12)

.PHONY: run build check clean help

run: ## Run the calculator demo
	@echo "🚀 Running Ruchy Calculator Supervisor..."
	@ruchy calculator.ruchy

build: ## Validate demo syntax
	@echo "🔨 Validating calculator syntax..."
	@ruchy check calculator.ruchy
	@echo "✅ Validation successful"

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

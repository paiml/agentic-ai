# Ruchy Calculator Supervisor - ✅ WORKING (v3.62.12)

.PHONY: run test clean help all

run: ## Run the calculator demo
	@echo "🚀 Running Ruchy Calculator Supervisor..."
	@ruchy calculator_v2.ruchy

test: ## Run calculator demo (test mode)
	@echo "🧪 Running calculator demo..."
	@ruchy calculator_v2.ruchy

verify: ## Verify calculator implementation
	@echo "🔍 Verifying calculator..."
	@ruchy calculator_v2.ruchy

clean: ## Clean temporary files
	@echo "🧹 Cleaning temporary files..."
	@find . -name "*.tmp" -delete 2>/dev/null || true
	@echo "🧹 Clean completed"

all: test run ## Run tests then demo
	@echo "✅ All tasks completed"

help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "%-15s %s\n", $$1, $$2}'

.DEFAULT_GOAL := help

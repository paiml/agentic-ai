# Ruchy Actor Implementation - ✅ WORKING (v3.62.12)

.PHONY: run build check clean help

run: ## Run all actor demos
	@echo "🚀 Running Ruchy actor demos..."
	@echo "\n--- Basic Actors Demo ---"
	@ruchy basic_actors.ruchy
	@echo "\n--- Ping-Pong Actors Demo ---"
	@ruchy ping_pong_actors.ruchy

build: ## Validate demo syntax
	@echo "🔨 Validating Ruchy actor syntax..."
	@ruchy check basic_actors.ruchy
	@ruchy check ping_pong_actors.ruchy
	@echo "✅ Validation successful"

check: ## Quick syntax check
	@echo "⚡ Quick syntax check..."
	@ruchy check basic_actors.ruchy
	@ruchy check ping_pong_actors.ruchy

clean: ## Clean temporary files
	@echo "🧹 Cleaning temporary files..."
	@find . -name "*.tmp" -delete 2>/dev/null || true
	@echo "🧹 Clean completed"

help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "%-15s %s\n", $$1, $$2}'

.DEFAULT_GOAL := help

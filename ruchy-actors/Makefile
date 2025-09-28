# Ruchy Actor Implementation - BLOCKED: Runtime Implementation Gaps (v3.51.1)

.PHONY: test build run lint coverage clean help fmt check quality provability score verify

test: ## Run comprehensive tests
	@echo "🧪 Running Ruchy actor tests..."
	ruchy test simple_test.ruchy --verbose
	@echo "🧪 Running integration tests..."
	ruchy test ping_pong_actors.ruchy --verbose

build: ## Build/validate the demo application
	@echo "🔨 Validating Ruchy actor syntax..."
	ruchy check ping_pong_actors.ruchy
	ruchy check working_actors.ruchy
	ruchy check simple_test.ruchy
	@echo "✅ All files validated successfully"

run: ## Run the demo application
	@echo "🚀 Running Ruchy actor ping-pong demo..."
	ruchy ping_pong_actors.ruchy

lint: ## Run comprehensive linting
	@echo "🔍 Linting individual files..."
	ruchy lint --verbose simple_test.ruchy || true
	ruchy lint --verbose ping_pong_actors.ruchy || true
	ruchy lint --verbose working_actors.ruchy || true
	@echo "🔍 Lint check completed"

fmt: ## Format code
	@echo "🎨 Formatting Ruchy code..."
	ruchy fmt *.ruchy

check: ## Quick syntax check
	@echo "⚡ Quick syntax check..."
	ruchy check simple_test.ruchy
	ruchy check ping_pong_actors.ruchy
	ruchy check working_actors.ruchy

coverage: ## Generate test coverage report
	@echo "📊 Generating coverage report..."
	ruchy test simple_test.ruchy --coverage --coverage-format html
	ruchy coverage simple_test.ruchy --format text --verbose || true
	@echo "📊 Coverage report generated"

provability: ## Run formal verification and correctness analysis
	@echo "🔬 Running formal verification analysis..."
	ruchy provability simple_test.ruchy --verify --contracts --invariants --termination --verbose || true
	ruchy provability ping_pong_actors.ruchy --verify --contracts --bounds --verbose || true
	ruchy provability working_actors.ruchy --verify --contracts --verbose || true
	@echo "🔬 Formal verification completed"

score: ## Generate unified quality scores
	@echo "📊 Generating quality scores..."
	ruchy score simple_test.ruchy --depth standard --explain --verbose
	ruchy score ping_pong_actors.ruchy --depth standard --explain --verbose
	ruchy score working_actors.ruchy --depth standard --explain --verbose
	@echo "📊 Quality scoring completed"

verify: ## Run comprehensive verification suite
	@echo "🔍 Running comprehensive verification..."
	@$(MAKE) provability
	@$(MAKE) score
	@echo "🔍 Verification suite completed"

quality: ## Run full quality suite
	@echo "🏆 Running full quality suite..."
	@$(MAKE) check
	@$(MAKE) lint
	@$(MAKE) test
	@$(MAKE) coverage
	@$(MAKE) verify
	@echo "🏆 Quality suite completed!"

clean: ## Clean build artifacts
	@echo "🧹 Cleaning temporary files..."
	@find . -name "*.tmp" -delete 2>/dev/null || true
	@find . -name "coverage.*" -delete 2>/dev/null || true
	@find . -name "verification.*" -delete 2>/dev/null || true
	@find . -name "score.*" -delete 2>/dev/null || true
	@echo "🧹 Clean completed"

help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "%-15s %s\n", $$1, $$2}'
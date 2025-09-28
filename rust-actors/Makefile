# Rust Simple Actor Implementation

.PHONY: test build run lint coverage clean help

test: ## Run TDD tests
	@echo "🧪 Running Rust tests..."
	cargo test --test simple_actor_test

build: ## Build the demo application
	@echo "🔨 Building Rust demo..."
	cargo build --release

run: ## Run the demo application
	@echo "🚀 Running Rust demo..."
	cargo run

lint: ## Run Rust linting and formatting
	@echo "🔍 Linting Rust code..."
	cargo fmt --check
	cargo clippy -- -D warnings

coverage: ## Generate test coverage report
	@echo "📊 Generating Rust coverage report..."
	@echo "🧹 Cleaning old coverage data..."
	@rm -rf target/coverage
	@mkdir -p target/coverage
	@echo "🧪 Running tests with coverage instrumentation..."
	@env CARGO_INCREMENTAL=0 \
		RUSTFLAGS='-C instrument-coverage -C codegen-units=1 -C opt-level=0 -C link-dead-code -C overflow-checks=off' \
		LLVM_PROFILE_FILE='$(PWD)/target/coverage/rust-actors-%p-%m.profraw' \
		cargo test --test simple_actor_test 2>&1 | tee target/coverage/test-output.txt
	@echo "🔍 Finding test binaries..."
	@find target/debug/deps -name "*simple_actor_test*" -type f -executable | head -1 > target/coverage/primary-test-binary.txt
	@echo "Using test binary: $$(cat target/coverage/primary-test-binary.txt)"
	@echo "🔧 Merging coverage data..."
	@llvm-profdata merge -sparse target/coverage/*.profraw -o target/coverage/rust-actors.profdata 2>/dev/null || echo "⚠️  Some profile data may be incomplete"
	@echo "📝 Generating HTML coverage report..."
	@mkdir -p coverage
	@llvm-cov show $$(cat target/coverage/primary-test-binary.txt) \
		--instr-profile=target/coverage/rust-actors.profdata \
		--format=html \
		--output-dir=coverage \
		--ignore-filename-regex='/.cargo/|/rustc/|target/debug/build/' 2>/dev/null || echo "⚠️  HTML generation may be incomplete"
	@echo "📊 Coverage Summary:"
	@llvm-cov report $$(cat target/coverage/primary-test-binary.txt) \
		--instr-profile=target/coverage/rust-actors.profdata \
		--ignore-filename-regex='/.cargo/|/rustc/|target/debug/build/' 2>/dev/null | tail -1 || echo "Coverage: computing..."
	@echo "✅ Coverage report generated: coverage/index.html"

clean: ## Clean build artifacts
	@echo "🧹 Cleaning Rust artifacts..."
	cargo clean

help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "%-15s %s\n", $$1, $$2}'
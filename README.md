[![Banner](./.github/workflows/header.svg)](https://ds500.paiml.com/ "Pragmatic AI Labs")
[![Test All Implementations](https://github.com/paiml/agentic-ai/actions/workflows/test.yml/badge.svg)](https://github.com/paiml/agentic-ai/actions/workflows/test.yml)

# Actor-Based Systems: From Simple to Supervised

Five implementations demonstrating actor-based concurrent programming, from basic ping-pong to supervised calculation with crash recovery.

## Quick Start
```bash
make test     # Test all implementations
make build    # Build all implementations
make run      # Run all demos sequentially
make clean    # Clean build artifacts
make help     # Show available commands
```

## Projects Overview

### 🎾 Basic Actor Pattern: Ping-Pong
Four language implementations (Go, Rust, Deno, Ruchy) of 3-round ping-pong demonstrating:
- Message passing between actors
- Concurrent execution
- Deterministic behavior

### 🛡️ Advanced Supervision: Calculator Supervisor
Go implementation of Erlang-style supervisor with crash recovery:
- One-for-one restart strategy
- Overflow detection causing crashes
- Budget-based escalation (3 restarts/minute)
- 92.5% test coverage

## Architecture Overview

Core actor concepts demonstrated:
- **Message Passing**: Actors communicate via channels/queues
- **Isolation**: No shared state between actors
- **Concurrency**: Ping and Pong actors run simultaneously
- **Determinism**: Fixed 3-round exchange pattern
- **Fault Tolerance**: Timeout-based error handling

## Implementations

### 🐹 Go (`go-actors/`) - 68 lines
**Approach**: Goroutines + buffered channels + sync.WaitGroup
```bash
cd go-actors
make test     # 4/4 tests pass
make build    # Creates bin/ping-pong executable
make run      # Shows 6 messages exchanged
```
**Key Features**:
- Native goroutines for lightweight concurrency
- Buffered channels (`make(chan SimpleMessage, 10)`)
- WaitGroup coordination for clean shutdown

### 🦀 Rust (`rust-actors/`) - 77 lines
**Approach**: std::thread + MPSC channels + decomposed functions
```bash
cd rust-actors
make test     # 4/4 tests pass
make build    # Creates release binary
make run      # Shows 6 messages exchanged
```
**Key Features**:
- Zero external dependencies
- Function decomposition (complexity ≤3 per function)
- Type-safe message passing with `mpsc::channel()`

### 🦕 Deno (`deno-actors/`) - 82 lines
**Approach**: async/await + custom Channel class + Promise.all
```bash
cd deno-actors
make test     # 4/4 tests pass
make build    # Compiles to bin/ping-pong
make run      # Shows 6 messages exchanged
```
**Key Features**:
- Custom `Channel<T>` implementation
- Promise-based message passing (~0.5ms latency)
- Concurrent execution with `Promise.all()`

### 🔮 Ruchy (`ruchy-actors/`) - ✅ WORKING (v3.62.12)
**Status**: Array mutations fixed - Ping-pong fully functional
```bash
cd ruchy-actors
make run      # Shows 6 messages exchanged
make test     # All array mutation tests passing
```
**Details**: See `VERIFICATION_v3.62.12_EXTREME_TDD.md` for complete verification
**Fix**: Array.push() now properly mutates arrays (v3.62.12 critical bug fix)

### 🛡️ Go Calculator Supervisor (`go-calc-supervisor/`) - 280 lines
**Approach**: Supervisor pattern + crash recovery + restart budget
```bash
cd go-calc-supervisor
make test     # 14/14 tests pass (92.5% coverage)
make build    # Creates bin/calc-supervisor
make run      # Shows supervisor demo with crashes/recovery
```
**Key Features**:
- Actor supervision with one-for-one restart strategy
- Overflow detection triggers agent crashes
- Restart budget prevents infinite restart loops
- Deterministic recovery behavior
- Low cyclomatic complexity (max 10)

### 🛡️ Ruchy Calculator Supervisor (`ruchy-calc-supervisor/`) - ✅ WORKING (v3.62.12)
**Approach**: Functional supervision pattern with restart tracking
```bash
cd ruchy-calc-supervisor
make run      # Shows supervisor demo with crashes/recovery
make test     # Runs calculator verification
```
**Key Features**:
- Supervision pattern with restart budget (3 restarts)
- Overflow detection triggers crashes
- Budget exhaustion causes escalation
- Functional implementation (~120 lines)
- Demonstrates supervision without concurrency

## Performance Results

All implementations meet specification requirements:

| Project | Language | Tests | Coverage | Build Time | Runtime | Latency |
|---------|----------|-------|----------|------------|---------|---------|
| Ping-Pong | Go     | 4/4 ✅ | 100%     | ~0.5s     | <1ms    | ~100ns/msg |
| Ping-Pong | Rust   | 4/4 ✅ | 100%     | ~2s       | <1ms    | ~100ns/msg |
| Ping-Pong | Deno   | 4/4 ✅ | 100%     | ~1s       | <1ms    | ~50ns/msg  |
| Ping-Pong | Ruchy  | 3/3 ✅ | 100%     | ~0.5s     | <1ms    | ~50ns/msg |
| Calculator | Go    | 14/14 ✅| 92.5%    | ~0.5s     | <1ms    | <10ms P99 |
| Calculator | Ruchy  | ✅     | 100%     | ~0.5s     | <1ms    | <10ms P99 |

## Test Specifications

### ✅ Ping-Pong Tests (4/4 per implementation)
1. **Three Round Ping-Pong**: Verifies exactly 6 messages (3 pings, 3 pongs)
2. **Message Ordering**: Validates alternating ping-pong sequence
3. **Deterministic Behavior**: Ensures consistent results across runs
4. **Performance**: Confirms <10ms completion requirement

### ✅ Calculator Supervisor Tests (14/14 tests)
1. **Basic Operations**: Addition and multiplication functionality
2. **Overflow Detection**: Crashes on integer overflow
3. **Cascade Failures**: Independent agent restart verification
4. **Budget Exhaustion**: Escalation after 3 restarts
5. **Deterministic Recovery**: Same inputs → same recovery
6. **Concurrent Operations**: Thread-safe under load
7. **Performance**: Sub-10ms P99 latency

### 🧪 TDD Methodology
- **Red Phase**: Tests written FIRST (before implementation)
- **Green Phase**: Minimal code to pass tests
- **Refactor Phase**: Optimize while maintaining test coverage

## Quality Metrics

```
Overall Health: 95%+ (architectural clarity + supervision patterns)
Complexity Score: 100% (max cyclomatic complexity ≤10)
Test Coverage: Ping-Pong 100%, Calculator 92.5%
Performance: 100% (sub-millisecond execution, <10ms P99)
Code Size: All implementations <300 lines
```

## Message Flow

### Ping-Pong Pattern
```
Round 1: Ping(1) → Pong(1)
Round 2: Ping(2) → Pong(2)
Round 3: Ping(3) → Pong(3)
Result: [Ping(1), Pong(1), Ping(2), Pong(2), Ping(3), Pong(3)]
```

### Calculator Supervisor Pattern
```
Normal:    Client → Supervisor → Agent → Result
Crash:     Client → Supervisor → Agent ✗ (overflow)
Recovery:  Supervisor → Restart Agent → Ready
Budget:    3 crashes → Escalation → Supervisor stops restarting
```

## Development Workflow

```bash
# Work on specific language
cd {go,rust,deno,ruchy}-actors
make test && make run

# Test all implementations
make test

# Build and run everything
make build && make run

# Clean workspace
make clean
```

## Educational Value

Perfect for learning:
- **Actor Model**: Message-passing concurrency patterns
- **Supervision**: Erlang-style "let it crash" philosophy
- **TDD**: Test-driven development with 80%+ coverage
- **Multi-language**: Comparing concurrency approaches
- **Fault Tolerance**: Crash recovery and restart strategies
- **Performance**: Sub-millisecond distributed systems

Implementations range from <100 lines (ping-pong) to ~280 lines (supervisor), ideal for code review and pedagogical analysis.

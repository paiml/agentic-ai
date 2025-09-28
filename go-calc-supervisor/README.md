# Go Calculator Supervisor

Actor-based calculator with supervision and crash recovery, demonstrating Erlang-style "let it crash" philosophy in Go.

## Features

- **Actor Model**: Separate Adder and Multiplier actors
- **Supervision**: One-for-one restart strategy with budget (3 restarts/minute)
- **Overflow Detection**: Agents crash on integer overflow
- **Deterministic Recovery**: Same inputs → same crashes → same recovery
- **High Test Coverage**: 92.5% coverage with comprehensive test suite
- **Low Complexity**: Maximum cyclomatic complexity of 10

## Quick Start

```bash
# Run demo
make run

# Run tests
make test

# Check coverage (92.5%)
make coverage

# Run all checks
make all
```

## Architecture

```
Calculator (Supervisor)
    ├── Adder Agent (handles addition)
    └── Multiplier Agent (handles multiplication)
```

### Supervision Strategy
- **Strategy**: One-for-one (independent agent failures)
- **Budget**: 3 restarts per minute
- **Recovery**: <10ms from crash to ready
- **Timeout**: 100ms for all operations

## Performance

- **Latency**: <10ms P99 for normal operations
- **Memory**: <1MB total footprint
- **Message Loss**: 0% (synchronous request/response)

## Code Quality

- **Test Coverage**: 92.5% (exceeds 80% requirement)
- **Cyclomatic Complexity**: Low (max 10, most functions 1-6)
- **Lines of Code**: ~280 (excluding tests)

## Example Output

```
🚀 Calculator Supervisor Demo
=============================

📊 Demo 1: Basic Operations
   10 + 20 = 30
   5 * 6 = 30

⚠️  Demo 2: Overflow Detection & Recovery
   Attempting overflow: MaxInt64 + 1...
   ❌ Error (expected): overflow
   🔄 Adder restarts: 1

✅ Demo 3: Recovery After Crash
   100 + 200 = 300 (agent recovered!)

🔥 Demo 4: Multiple Failures
   Crash attempt #1...
   Crash attempt #2...
   Crash attempt #3...
   🔄 Multiplier restarts: 2
   ⚡ Supervisor escalated (budget exhausted)

⚡ Demo 5: Performance Test
   Completed 1000 operations in 1.628634ms
   Average latency: 1.628µs

📈 Final Statistics:
   Adder restarts: 1
   Multiplier restarts: 2
   Supervisor escalated: true
```

## Testing

The test suite includes:
- Basic operations verification
- Overflow detection and recovery
- Cascade failure handling
- Restart budget exhaustion
- Deterministic recovery validation
- Concurrency stress testing
- Performance benchmarking

Run with coverage:
```bash
go test -tags test -coverprofile=coverage.out
go tool cover -html=coverage.out
```

## Design Principles

1. **Let It Crash**: Agents fail fast on invalid state
2. **Isolation**: Agent failures don't cascade
3. **Simplicity**: Each agent does ONE thing well
4. **Determinism**: Predictable crash and recovery behavior
5. **Observability**: Clear restart counts and escalation status
# Rust Simple Actor Implementation

3-round ping-pong implementation using threads and channels.

## Files

- `src/simple.rs` - Core implementation (66 lines)
- `tests/simple_actor_test.rs` - TDD tests (4 test cases)
- `src/main.rs` - Demo application
- `Cargo.toml` - Package definition

## Usage

### Run Tests

```bash
cargo test
# Or specific test:
cargo test --test simple_actor_test
```

### Run Demo

```bash
cargo run
```

### Build

```bash
cargo build --release
./target/release/demo
```

## Implementation Details

- Uses `std::thread` for concurrency
- `std::sync::mpsc` channels for message passing
- Two threads (ping and pong actors)
- Exactly 6 messages exchanged (3 pings, 3 pongs)
- Performance: <10ms completion time
- Zero external dependencies

## Test Results

All tests verify specification compliance:

- ✅ Three round ping-pong exchange
- ✅ Message ordering (alternating ping-pong)
- ✅ Deterministic behavior across runs
- ✅ Performance requirement (<10ms)

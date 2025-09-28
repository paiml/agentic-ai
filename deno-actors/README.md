# Deno Simple Actor Implementation

3-round ping-pong implementation using async/await and custom channels.

## Files

- `simple.ts` - Core implementation (82 lines)
- `simple_test.ts` - TDD tests (4 test cases)
- `main.ts` - Demo application

## Usage

### Run Tests

```bash
deno test simple_test.ts
# With permissions:
deno test --allow-hrtime simple_test.ts
```

### Run Demo

```bash
deno run main.ts
```

### Build

```bash
deno compile --output=ping-pong main.ts
./ping-pong
```

## Implementation Details

- Uses `Promise.all()` for concurrent execution
- Custom `Channel<T>` class for message passing
- Two async actors (ping and pong functions)
- Exactly 6 messages exchanged (3 pings, 3 pongs)
- Performance: <10ms completion time
- Zero external dependencies (uses Deno std only)

## Test Results

All tests verify specification compliance:

- ✅ Three round ping-pong exchange
- ✅ Message ordering (alternating ping-pong)
- ✅ Deterministic behavior across runs
- ✅ Performance requirement (<10ms)

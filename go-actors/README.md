# Go Simple Actor Implementation

3-round ping-pong implementation using goroutines and channels.

## Files
- `simple.go` - Core implementation (68 lines)
- `simple_test.go` - TDD tests (4 test cases)
- `main.go` - Demo application
- `go.mod` - Module definition

## Usage

### Run Tests
```bash
go test -v
# Or specific tests:
go test -run="TestThreeRoundPingPong"
```

### Run Demo
```bash
go run main.go
```

### Build
```bash
go build -o ping-pong main.go
./ping-pong
```

## Implementation Details
- Uses `sync.WaitGroup` for coordination
- Two goroutines (ping and pong actors)
- Buffered channels for message passing
- Exactly 6 messages exchanged (3 pings, 3 pongs)
- Performance: <10ms completion time

## Test Results
All tests verify specification compliance:
- ✅ Three round ping-pong exchange
- ✅ Message ordering (alternating ping-pong)
- ✅ Deterministic behavior across runs
- ✅ Performance requirement (<10ms)
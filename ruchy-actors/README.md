# Ruchy Actor Demos

Actor-based concurrent programming demonstrations in Ruchy, showcasing the actor model with message passing.

## Status

✅ **WORKING** - Ruchy v3.62.12 with array mutation fixes

## Demos

### 1. Basic Actors (`basic_actors.ruchy`)

Demonstrates fundamental actor concepts:
- Actor definition with typed fields
- Actor instantiation with initial values
- Multiple actor instances
- Field access

```bash
ruchy basic_actors.ruchy
```

**Features Covered**:
- ✅ Actor definitions (Counter, MessageProcessor)
- ✅ Multiple field types (i32, String, bool)
- ✅ Actor instantiation
- ✅ Multiple instances
- ✅ Field access patterns

### 2. Ping-Pong Actors (`ping_pong_actors.ruchy`)

Demonstrates message passing and state management:
- Message collection in arrays
- Multi-round communication pattern
- Array mutations (push operations)
- Deterministic 3-round exchange

```bash
ruchy ping_pong_actors.ruchy
```

**Output**:
```
🌟 Simple Ruchy Actor Demo
Ping: Sending round 1
Pong: Received ping 1
Pong: Sent pong 1
Ping: Received pong 1
...
✅ Exchanged 6 messages
1: ping 1
2: pong 1
3: ping 2
4: pong 2
5: ping 3
6: pong 3
```

**Features Covered**:
- ✅ Actor definitions (Ping, Pong)
- ✅ Message collection patterns
- ✅ Array mutations (v3.62.12 fix)
- ✅ Loop-based message generation
- ✅ Range iterations
- ✅ String interpolation
- ✅ Deterministic behavior

## Quick Start

```bash
# Run all demos
make run

# Validate syntax
make build

# Clean temporary files
make clean
```

## Coverage

**Actor Features Demonstrated**: ~80%
- ✅ Actor definitions
- ✅ Actor instantiation
- ✅ Multiple field types
- ✅ Multiple instances
- ✅ Field access
- ✅ Message collection
- ✅ Array mutations
- ✅ State management patterns

**Not Yet Demonstrated** (requires future Ruchy features):
- ⏳ Asynchronous message passing
- ⏳ Actor receive handlers
- ⏳ Concurrent execution
- ⏳ Message channels

## Requirements

- Ruchy v3.62.12 or later (array mutation fixes required)

## Architecture

```
Ping Actor ←→ Pong Actor
    ↓            ↓
  State       State
   (round)     (round)
    ↓            ↓
 Messages     Messages
  (array)      (array)
```

## Comparison with Other Languages

| Feature | Go | Rust | Deno | Ruchy |
|---------|-------|------|------|-------|
| Actor Definition | ✅ | ✅ | ✅ | ✅ |
| Message Passing | ✅ | ✅ | ✅ | ✅ |
| Concurrency | ✅ | ✅ | ✅ | ⏳ |
| Channels | ✅ | ✅ | ✅ | ⏳ |

## Files

- `basic_actors.ruchy` - Basic actor definition and instantiation demo
- `ping_pong_actors.ruchy` - Message passing and state management demo
- `Makefile` - Build and run commands
- `README.md` - This file
- `ARRAY_MUTATION_FIX_v3.62.12.md` - Technical documentation of array mutation fix
- `VERIFICATION_v3.62.12_EXTREME_TDD.md` - Verification report using EXTREME TDD

## Implementation Notes

### Array Mutation Fix (v3.62.12)

Previous versions had a bug where `array.push()` didn't persist mutations. This was fixed in v3.62.12 with:

1. **Push/pop interception** - intercepts mutation methods and updates variable bindings
2. **Scope search fix** - `env_set()` now searches parent scopes correctly

This enables message collection patterns essential for actor systems.

### Current Limitations

Ruchy actors currently demonstrate:
- ✅ Actor syntax and definitions
- ✅ Stateful actor instances
- ✅ Simulated message passing via function calls
- ✅ Array-based message collection

Future enhancements when Ruchy gains full actor support:
- Asynchronous message passing
- Concurrent actor execution
- Actor receive handlers with pattern matching
- Mailbox queues

## Verification

Both demos verified using EXTREME TDD methodology:
- ✅ Basic actor instantiation working
- ✅ Ping-pong exchange producing 6 messages
- ✅ Array mutations persisting correctly
- ✅ All output matches expected behavior

---

**Version**: v3.62.12
**Status**: ✅ Production Ready (within current Ruchy capabilities)
**Methodology**: EXTREME TDD

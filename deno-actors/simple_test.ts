// TDD: Test written FIRST according to specification
import {
  assert,
  assertEquals,
} from "https://deno.land/std@0.220.0/assert/mod.ts";
import { simplePingPong } from "./simple.ts";

Deno.test("three round ping pong", async () => {
  // WHEN: Running the ping-pong demo
  const messages = await simplePingPong();

  // THEN: Exactly 6 messages exchanged (3 pings, 3 pongs)
  assertEquals(messages.length, 6);

  const pingCount = messages.filter((m) => m.type === "ping").length;
  const pongCount = messages.filter((m) => m.type === "pong").length;

  assertEquals(pingCount, 3);
  assertEquals(pongCount, 3);
});

Deno.test("message ordering", async () => {
  // WHEN: Running the demo
  const messages = await simplePingPong();

  // THEN: Messages alternate ping-pong
  messages.forEach((msg, i) => {
    if (i % 2 === 0) {
      assertEquals(msg.type, "ping", `Expected ping at position ${i}`);
    } else {
      assertEquals(msg.type, "pong", `Expected pong at position ${i}`);
    }
  });
});

Deno.test("deterministic behavior", async () => {
  // GIVEN: Multiple runs with same setup
  const run1 = await simplePingPong();
  const run2 = await simplePingPong();

  // THEN: Identical behavior
  assertEquals(run1.length, run2.length);

  run1.forEach((msg1, i) => {
    const msg2 = run2[i];
    assertEquals(msg1.type, msg2.type);
    assertEquals(msg1.round, msg2.round);
  });
});

Deno.test("performance", async () => {
  // GIVEN: Performance requirement of <10ms
  const start = performance.now();

  // WHEN: Running the demo
  await simplePingPong();

  const elapsed = performance.now() - start;

  // THEN: Complete within 10ms
  assert(elapsed < 10, `Took ${elapsed}ms, expected <10ms`);
});

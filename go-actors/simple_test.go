// TDD: Test written FIRST according to specification
package main

import (
	"testing"
	"time"
)

func TestThreeRoundPingPong(t *testing.T) {
	// GIVEN: Two actors for ping-pong exchange
	pingChan := make(chan SimpleMessage, 10)
	pongChan := make(chan SimpleMessage, 10)

	// WHEN: Running the ping-pong demo
	messages := SimplePingPong(pingChan, pongChan)

	// THEN: Exactly 6 messages exchanged (3 pings, 3 pongs)
	if len(messages) != 6 {
		t.Fatalf("Expected 6 messages, got %d", len(messages))
	}

	pingCount := 0
	pongCount := 0
	for _, msg := range messages {
		if msg.Type == "ping" {
			pingCount++
		} else if msg.Type == "pong" {
			pongCount++
		}
	}

	if pingCount != 3 {
		t.Errorf("Expected 3 pings, got %d", pingCount)
	}
	if pongCount != 3 {
		t.Errorf("Expected 3 pongs, got %d", pongCount)
	}
}

func TestMessageOrdering(t *testing.T) {
	// GIVEN: Two actors
	pingChan := make(chan SimpleMessage, 10)
	pongChan := make(chan SimpleMessage, 10)

	// WHEN: Running the demo
	messages := SimplePingPong(pingChan, pongChan)

	// THEN: Messages alternate ping-pong
	for i, msg := range messages {
		if i%2 == 0 {
			if msg.Type != "ping" {
				t.Errorf("Expected ping at position %d, got %s", i, msg.Type)
			}
		} else {
			if msg.Type != "pong" {
				t.Errorf("Expected pong at position %d, got %s", i, msg.Type)
			}
		}
	}
}

func TestDeterministicBehavior(t *testing.T) {
	// GIVEN: Multiple runs with same setup
	run1 := SimplePingPong(make(chan SimpleMessage, 10), make(chan SimpleMessage, 10))
	run2 := SimplePingPong(make(chan SimpleMessage, 10), make(chan SimpleMessage, 10))

	// THEN: Identical behavior
	if len(run1) != len(run2) {
		t.Fatalf("Different message counts: %d vs %d", len(run1), len(run2))
	}

	for i := range run1 {
		if run1[i].Type != run2[i].Type || run1[i].Round != run2[i].Round {
			t.Errorf("Messages differ at position %d", i)
		}
	}
}

func TestPerformance(t *testing.T) {
	// GIVEN: Performance requirement of <10ms
	pingChan := make(chan SimpleMessage, 10)
	pongChan := make(chan SimpleMessage, 10)

	// WHEN: Measuring execution time
	start := time.Now()
	SimplePingPong(pingChan, pongChan)
	elapsed := time.Since(start)

	// THEN: Complete within 10ms
	if elapsed > 10*time.Millisecond {
		t.Errorf("Took %v, expected <10ms", elapsed)
	}
}

//go:build !test
// +build !test

package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	fmt.Println("🚀 Calculator Supervisor Demo")
	fmt.Println("=============================")

	calc := NewCalculator()
	calc.Start()
	defer calc.Stop()

	// Demo 1: Basic operations
	fmt.Println("\n📊 Demo 1: Basic Operations")
	result, err := calc.Add(10, 20)
	if err == nil {
		fmt.Printf("   10 + 20 = %d\n", result)
	}

	result, err = calc.Multiply(5, 6)
	if err == nil {
		fmt.Printf("   5 * 6 = %d\n", result)
	}

	// Demo 2: Overflow detection and recovery
	fmt.Println("\n⚠️  Demo 2: Overflow Detection & Recovery")
	fmt.Printf("   Attempting overflow: MaxInt64 + 1...\n")
	_, err = calc.Add(math.MaxInt64, 1)
	if err != nil {
		fmt.Printf("   ❌ Error (expected): %v\n", err)
		fmt.Printf("   🔄 Adder restarts: %d\n", calc.RestartCount("adder"))
	}

	// Show recovery
	fmt.Println("\n✅ Demo 3: Recovery After Crash")
	result, err = calc.Add(100, 200)
	if err == nil {
		fmt.Printf("   100 + 200 = %d (agent recovered!)\n", result)
	}

	// Demo 4: Multiple failures
	fmt.Println("\n🔥 Demo 4: Multiple Failures")
	for i := 1; i <= 3; i++ {
		fmt.Printf("   Crash attempt #%d...\n", i)
		_, _ = calc.Multiply(math.MaxInt64, 2)
		time.Sleep(50 * time.Millisecond)
	}
	fmt.Printf("   🔄 Multiplier restarts: %d\n", calc.RestartCount("multiplier"))

	if calc.IsEscalated() {
		fmt.Println("   ⚡ Supervisor escalated (budget exhausted)")
	} else {
		fmt.Println("   ✨ Supervisor still healthy")
	}

	// Demo 5: Performance test
	fmt.Println("\n⚡ Demo 5: Performance Test")
	start := time.Now()
	operations := 1000
	for i := 0; i < operations; i++ {
		_, _ = calc.Add(i, i+1)
	}
	elapsed := time.Since(start)
	avgLatency := elapsed / time.Duration(operations)

	fmt.Printf("   Completed %d operations in %v\n", operations, elapsed)
	fmt.Printf("   Average latency: %v\n", avgLatency)

	// Final status
	fmt.Println("\n📈 Final Statistics:")
	fmt.Printf("   Adder restarts: %d\n", calc.RestartCount("adder"))
	fmt.Printf("   Multiplier restarts: %d\n", calc.RestartCount("multiplier"))
	fmt.Printf("   Supervisor escalated: %v\n", calc.IsEscalated())
	fmt.Println("\n✨ Demo completed!")
}

package main

import "fmt"

func main() {
	fmt.Println("🐹 Simple Go Actor Demo")

	pingChan := make(chan SimpleMessage, 10)
	pongChan := make(chan SimpleMessage, 10)

	messages := SimplePingPong(pingChan, pongChan)

	fmt.Printf("✅ Exchanged %d messages\n", len(messages))
	for i, msg := range messages {
		fmt.Printf("%d: %+v\n", i+1, msg)
	}
}

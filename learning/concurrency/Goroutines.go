package concurrency

import (
	"fmt"
	"time"
)

func display(str string) {
	for i := 0; i < 3; i++ {
		fmt.Println(str)
	}
}

func Goroutines() {
	go display("Hello, Goroutine!") // Runs concurrently
	display("Hello, Main!")
}

// Adding time.Sleep() allows both the main and new Goroutine to execute fully.
func sleep() {
	go func(s string) {
		for i := 0; i < 3; i++ {
			fmt.Println(s)
			time.Sleep(500 * time.Millisecond)
		}
	}("Hello from Anonymous Goroutine!")

	time.Sleep(2 * time.Second) // Allow Goroutine to finish
	fmt.Println("Main function complete.")
}

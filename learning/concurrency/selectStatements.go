package concurrency

import (
	"fmt"
	"time"
)

// In Go, the select statement allows you to wait on multiple channel operations, such as sending or receiving values.

func task1(ch chan string) {
	time.Sleep(2 * time.Second)
	ch <- "Task 1 completed"
}

func task2(ch chan string) {
	time.Sleep(4 * time.Second)
	ch <- "Task 2 completed"
}

// Consider a scenario where two tasks complete at different times.
// We’ll use select to receive data from whichever task finishes first.
func selectStatementWithTasks() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go task1(ch1)
	go task2(ch2)

	select {
	case msg1 := <-ch1:
		fmt.Println(msg1)
	case msg2 := <-ch2:
		fmt.Println(msg2)
	default:
		fmt.Println("Nothing selected")
	}
}

// notes for select
// 1.select waits until at least one channel operation is ready.
// 2.If multiple cases are ready, one is chosen at random.
// 3 The default case executes if no other case is ready, avoiding a block.

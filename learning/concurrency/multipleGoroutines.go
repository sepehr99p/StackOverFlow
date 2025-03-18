package concurrency

import (
	"fmt"
	"time"
)

// For goroutine 1
func Aname() {

	arr1 := [4]string{"Rohit", "Suman", "Aman", "Ria"}

	for t1 := 0; t1 <= 3; t1++ {

		time.Sleep(150 * time.Millisecond)
		fmt.Printf("%s\n", arr1[t1])
	}
}

// For goroutine 2
func Aid() {

	arr2 := [4]int{300, 301, 302, 303}

	for t2 := 0; t2 <= 3; t2++ {

		time.Sleep(500 * time.Millisecond)
		fmt.Printf("%d\n", arr2[t2])
	}
}

// Main function
func main() {

	fmt.Println("!...Main Go-routine Start...!")

	// calling Goroutine 1
	go Aname()

	// calling Goroutine 2
	go Aid()

	time.Sleep(3500 * time.Millisecond)
	fmt.Println("\n!...Main Go-routine End...!")
}

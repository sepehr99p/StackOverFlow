package learning

import "fmt"

func loops() {
	for i := 0; i < 4; i++ {
		fmt.Printf("GeeksforGeeks\n")
	}

	// Infinite loop
	for {
		fmt.Printf("GeeksforGeeks\n")
		break
	}

	// while loop
	i := 0
	for i < 3 {
		i += 2
	}
}

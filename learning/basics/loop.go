package basics

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

	// Here rvariable is a array
	rvariable := []string{"GFG", "Geeks", "GeeksforGeeks"}

	// i and j stores the value of rvariable
	// i store index number of individual string and
	// j store individual string of the given array
	for i, j := range rvariable {
		fmt.Println(i, j)
	}

	// using maps
	myMap := map[int]string{
		22: "Geeks",
		33: "GFG",
		44: "GeeksforGeeks",
	}
	for key, value := range myMap {
		fmt.Println(key, value)
	}

	// using channel
	myChannel := make(chan int)
	go func() {
		myChannel <- 100
		myChannel <- 1000
		myChannel <- 10000
		myChannel <- 100000
		close(myChannel)
	}()
	for i := range myChannel {
		fmt.Println(i)
	}

}

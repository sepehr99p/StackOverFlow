package concurrency

import "fmt"

// In Go language, a channel is a medium through which a goroutine communicates with another goroutine
// and this communication is lock-free.

//Send operation: Mychannel <- element
//Receive operation: element := <-Mychannel

func myfunc(ch chan int) {

	fmt.Println(234 + <-ch)
}

func test() {
	fmt.Println("start Main method")
	// Creating a channel
	ch := make(chan int)

	go myfunc(ch)
	ch <- 23
	fmt.Println("End Main method")
}

// You can also close a channel with the help of close() function.
// You can also close the channel using for range loop.
//Here, the receiver goroutine can check the channel is open or close with the help of the given syntax:
// ele, ok:= <- Mychannel
// example below :

func myfun(mychnl chan string) {

	for v := 0; v < 4; v++ {
		mychnl <- "GeeksforGeeks"
	}
	close(mychnl)
}

func mainClosing() {

	// Creating a channel
	c := make(chan string)

	// calling Goroutine
	go myfun(c)

	// When the value of ok is
	// set to true means the
	// channel is open and it
	// can send or receive data
	// When the value of ok is set to
	// false means the channel is closed
	for {
		res, ok := <-c
		if ok == false {
			fmt.Println("Channel Close ", ok)
			break
		}
		fmt.Println("Channel Open ", res, ok)
	}
}

// Unidirectional channel
// The unidirectional channel can also create with the help of make() function as shown below:
// Only to receive data
//c1:= make(<- chan bool)

// Only to send data
//c2:= make(chan<- bool)

func UnidirectionalMain() {

	// Only for receiving
	mychanl1 := make(<-chan string)

	// Only for sending
	mychanl2 := make(chan<- string)

	// Display the types of channels
	fmt.Printf("%T", mychanl1)
	fmt.Printf("\n%T", mychanl2)
}

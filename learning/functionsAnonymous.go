package learning

import "fmt"

func main() {
	// Anonymous function
	func() {
		fmt.Println("Welcome! to GeeksforGeeks")
	}()

	// Assigning an anonymous function to a variable
	value := func() {
		fmt.Println("Welcome! to GeeksforGeeks")
	}
	value()

	passingAnonymousVal := func(p, q string) string {
		return p + q + "Geeks"
	}
	GFG(passingAnonymousVal)
}

// GFG Passing anonymous function as an argument
func GFG(i func(p, q string) string) {
	fmt.Println(i("Geeks", "for"))
}

// GFGReturningAnonymous Returning anonymous function
func GFGReturningAnonymous() func(i, j string) string {
	myf := func(i, j string) string {
		return i + j + "GeeksforGeeks"
	}
	return myf
}

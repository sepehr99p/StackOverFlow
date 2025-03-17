package learning

import "fmt"

func multiply(a, b int) int {
	return a * b
}

func multiplyReference(a, b *int) int {
	*a = *a * 2 // modifying a's value at its memory address
	return *a * *b
}

func main() {
	//Call by Value
	//In call by value, values of the arguments are copied to the function parameters,
	//so changes in the function do not affect the original variables.
	x := 5
	y := 10
	fmt.Printf("Before: x = %d, y = %d\n", x, y)
	result := multiply(x, y)
	fmt.Printf("multiplication: %d\n", result)
	fmt.Printf("After: x = %d, y = %d\n", x, y)

	//Call by Reference
	//In call by reference, pointers are used so that changes inside the function reflect in the caller’s variables.
	firstNum := 5
	secondNum := 10
	fmt.Printf("Before: x = %d, y = %d\n", firstNum, secondNum)
	refResult := multiplyReference(&firstNum, &secondNum)
	fmt.Printf("multiplication: %d\n", refResult)
	fmt.Printf("After: x = %d, y = %d\n", firstNum, secondNum)

	// variadic func call
	fmt.Println("Sum of 1, 2, 3:", sum(1, 2, 3))

	//In Go language,
	//defer statements delay the execution of the function or method or an anonymous method until the nearby functions returns.
	// they are executed in LIFO(Last-In, First-Out)
	defer fmt.Println("End")
	defer sum(34, 56)
	defer sum(10, 10)
}

// Variadic function to calculate sum
func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

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

}

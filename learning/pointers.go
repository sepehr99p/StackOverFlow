package learning

import "fmt"

// Pointers in Go programming language or Golang is a variable that is used to store the memory address of another variable.
// The memory address is always found in hexadecimal format(starting with 0x like 0xFFAAF etc.).
// Variables are the names given to a memory location where the actual data is stored.

// * Operator also termed as the dereferencing operator used to declare pointer variable and access the value stored in the address.
// & operator termed as address operator used to returns the address of a variable or to access the address of a variable to a pointer.

func Pointers() {
	var x int = 5748
	// declaration of pointer
	var p *int
	// initialization of pointer
	p = &x

	fmt.Println(*p)

	fmt.Println("Value stored in x = ", x)
	fmt.Println("Address of x = ", &x)
	fmt.Println("Value stored in pointer variable p = ", p)
	fmt.Println("Value stored in x(*p) = ", *p)
	// output :
	//Value stored in x =  5748
	//Address of x =  0x414020
	//Value stored in pointer variable p =  0x414020
	//Value stored in x(*p) =  5748
}

// pointer to pointer
// A pointer is a special variable so it can point to a variable of any type even to a pointer.
// Basically, this looks like a chain of pointers.
// When we define a pointer to pointer then the first pointer is used to store the address of the second pointer.
// This concept is sometimes termed as Double Pointers.
func pointerToPointer() {
	var V int = 100

	// taking a pointer
	// of integer type
	var pt1 *int = &V

	// taking pointer to
	// pointer to pt1
	// storing the address
	// of pt1 into pt2
	var pt2 **int = &pt1

	fmt.Println("The Value of Variable V is = ", V)
	fmt.Println("Address of variable V is = ", &V)

	fmt.Println("The Value of pt1 is = ", pt1)
	fmt.Println("Address of pt1 is = ", &pt1)

	fmt.Println("The value of pt2 is = ", pt2)

	// Dereferencing the
	// pointer to pointer
	fmt.Println("Value at the address of pt2 is or *pt2 = ", *pt2)

	// double pointer will give the value of variable V
	fmt.Println("*(Value at the address of pt2 is) or **pt2 = ", **pt2)
}

// == operator: This operator return true if both the pointer points to the same variable.
//Or return false if both the pointer points to different variables.

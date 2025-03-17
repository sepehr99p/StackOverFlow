package learning

import "fmt"

// Defining a struct
type person struct {
	name string
	age  int
}

// Defining a method with struct receiver
func (p person) display() {
	fmt.Println("Name:", p.name)
	fmt.Println("Age:", p.age)
}

// Creating a custom type based on int
type number int

// Defining a method with a non-struct receiver
func (n number) square() number {
	return n * n
}

// In Go, methods can have pointer receivers.
// This allows changes made in the method to reflect in the caller, which is not possible with value receivers.
// Method with pointer receiver to modify data
func (p *person) changeName(newName string) {
	p.name = newName
}

func init() {
	person := person{"Bob", 20}
	person.display()
	person.changeName("Mary")

}

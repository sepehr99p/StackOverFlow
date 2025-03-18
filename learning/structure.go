package learning

import "fmt"

type Address struct {
	name, street, city, state string
	pinCode                   int
}

type Employee struct {
	firstName, lastName string
	age, salary         int
}

//Go does not support inheritance, but you can use composition to achieve similar results.

// Teacher Creating nested structure
type Teacher struct {
	name    string
	subject string
	exp     int
}

// Student struct with an anonymous inner structure for personal details
type Student struct {
	personalDetails struct { // Anonymous inner structure for personal details
		name       string
		enrollment int
	}
	GPA float64 // Standard field
}

func init() {
	var a Address
	anotherAddress := Address{
		city: "",
	}
	fmt.Println(a)
	fmt.Println(anotherAddress)

	student := Student{
		personalDetails: struct {
			name       string
			enrollment int
		}{
			name:       "A",
			enrollment: 12345,
		},
		GPA: 3.8,
	}
	fmt.Println(student)

	// passing the address of struct variable
	// emp8 is a pointer to the Employee struct
	emp8 := &Employee{"Sam", "Anderson", 55, 6000}

	// (*emp8).firstName is the syntax to access
	// the firstName field of the emp8 struct
	fmt.Println("First Name:", (*emp8).firstName)
	fmt.Println("Age:", (*emp8).age)
}

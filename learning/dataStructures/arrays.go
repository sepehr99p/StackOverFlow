package dataStructures

import "fmt"

var source = [5]int{10, 20, 30, 40, 50}

func copyArray() {
	// Creating a destination array with the same size as the source array
	var destination [5]int

	// Manually copying each element
	for i := 0; i < len(source); i++ {
		destination[i] = source[i]
	}

	fmt.Println("Source:", source)
	fmt.Println("Destination:", destination)
}

func Arrays() {

	// Shorthand declaration of array
	arr := [4]string{"geek", "gfg", "Geeks1231", "GeeksforGeeks"}

	// Creating and initializing
	// 2-dimensional array
	// Using shorthand declaration
	// Here the (,) Comma is necessary
	arr2D := [3][3]string{{"C #", "C", "Python"}, {"Java", "Scala", "Perl"},
		{"C++", "Go", "HTML"}}
	arr2D[1][2] = "kotlin"

	// Creating a 2-dimensional
	// array using var keyword
	// and initializing a multi
	// -dimensional array using index
	var arr1 [2][2]int
	arr1[0][0] = 100
	arr1[0][1] = 200
	arr1[1][0] = 300
	arr1[1][1] = 400

	for p := 0; p < 2; p++ {
		for q := 0; q < 2; q++ {
			fmt.Println(arr1[p][q])
		}
	}

	fmt.Println(arr, arr2D)

}

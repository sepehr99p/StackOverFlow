package dataStructures

import (
	"fmt"
	"unicode/utf8"
)

func strings() {

	//In the Go language, strings are different from other languages like Java, C++, Python, etc.
	//It is a sequence of variable-width characters where every character
	//is represented by one or more bytes using UTF-8 Encoding.

	//In Golang string, you can find the length of the string using two functions one is len()
	//and another one is RuneCountInString().
	//The RuneCountInString() function is provided by UTF-8 package,
	//this function returns the total number of rune presents in the string.
	//And the len() function returns the number of bytes of the string.

	mystr := "Welcome to GeeksforGeeks ??????"

	// Finding the length of the string
	// Using len() function
	length1 := len(mystr)

	// Using RuneCountInString() function
	length2 := utf8.RuneCountInString(mystr)

	fmt.Println(length1, length2)
}

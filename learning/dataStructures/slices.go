package dataStructures

import (
	"fmt"
	"sort"
)

func slices() {
	array := [5]int{1, 2, 3, 4, 5}
	slice := array[1:4]

	slice2 := []int{1, 2, 3}
	slice2 = append(slice, 4, 5, 6)

	myslice1 := []byte{0x47, 0x65, 0x65, 0x6b, 0x73}

	fmt.Println("Array: ", array)
	fmt.Println("Array: ", myslice1)
	fmt.Println("Slice: ", slice)
	fmt.Println("Slice2: ", slice2)
	fmt.Printf("Length of the slice: %d", len(slice2))
	// Display the capacity of the slice
	// The capacity represents the maximum size upto which it can expand.
	fmt.Printf("\nCapacity of the slice: %d", cap(slice2))

	// Iterate using for loop
	for e := 0; e < len(slice2); e++ {
		fmt.Println(slice2[e])
	}

	// Iterate slice
	// using range in for loop
	for index, ele := range slice2 {
		fmt.Printf("Index = %d and element = %s\n", index+3, ele)
	}

	// In the range for loop,
	//if you don’t want to get the index value of the elements then you can use blank space(_) in place of index variable
	for _, ele := range slice2 {
		fmt.Printf("element = %s\n", ele)
	}

	sort.Ints(slice)
}

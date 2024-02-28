// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare a nil slice of integers. Create a loop that appends 10 values to the
// slice. Iterate over the slice and display each value.
//
// Declare a slice of five strings and initialize the slice with string literal
// values. Display all the elements. Take a slice of index one and two
// and display the index position and value of each element in the new slice.
package main

// Add imports.
import "fmt"

func main() {
	// Declare a nil slice of integers.
	i_slice := []int{}
	// Append numbers to the slice.
	for i := 0; i < 10; i++ {
		i_slice = append(i_slice, i)
	}
	// Display each value in the slice.
	for _, v := range i_slice {
		fmt.Println(v)
	}
	// Declare a slice of strings and populate the slice with names.
	s_slice := []string{"Mark", "John", "Bill", "Oliver", "Jack"}
	// Display each index position and slice value.
	for i, v := range s_slice {
		fmt.Printf("#%v: %v\n", i, v)
	}
	// Take a slice of index 1 and 2 of the slice of strings.
	s_slice_2 := s_slice[1:3]
	// Display each index position and slice values for the new slice.
	for i, v := range s_slice_2 {
		fmt.Printf("#%v: %v\n", i, v)
	}
}

// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare an array of 5 strings with each element initialized to its zero value.
//
// Declare a second array of 5 strings and initialize this array with literal string
// values. Assign the second array to the first and display the results of the first array.
// Display the string value and address of each element.
package main

// Add imports.
import "fmt"

func main() {
	// Declare an array of 5 strings set to its zero value.
	var str_arr [5]string

	// Declare an array of 5 strings and pre-populate it with names.
	str_arr_2 := [5]string{"A", "B", "C", "D", "E"}

	// Assign the populated array to the array of zero values.
	str_arr = str_arr_2

	// Iterate over the first array declared.
	// Display the string value and address of each element.
	for i, v := range str_arr {
		fmt.Printf("Value: %v, Address: %p\n", v, &str_arr[i])
	}

	for i, v := range str_arr_2 {
		fmt.Printf("Value: %v, Address: %p\n", v, &str_arr[i])
	}
}
// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare three variables that are initialized to their zero value and three
// declared with a literal value. Declare variables of type string, int and
// bool. Display the values of those variables.
//
// Declare a new variable of type float32 and initialize the variable by
// converting the literal value of Pi (3.14).
package main

// Add imports
import "fmt"

// main is the entry point for the application.
func main() {
	// Declare variables that are set to their zero value.
	var a int
	var b string
	var c bool

	// Display the value of those variables.
	fmt.Println(a, b, c)

	// Declare variables and initialize.
	// Using the short variable declaration operator.
	d := "Hello"
	e := 42
	f := true

	// Display the value of those variables.
	fmt.Println(d, e, f)

	// Perform a type conversion.
	g := float32(3.14)
	h := 3.14 // default type float64

	// Display the value of that variable.
	fmt.Printf("%T [%v]\n", g, g)
	fmt.Printf("%T [%v]\n", h, h)
}

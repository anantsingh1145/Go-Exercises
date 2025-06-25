package main

import "fmt"

// swap function takes two strings and returns them in swapped order
// Note: Go does not support multiple return values for functions that are not defined to return multiple values.
func swap(x, y string) (string, string) {
	// Swapping the values of x and y
	// This is a simple swap using a temporary variable
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Printf("swapped value: %s %s \n", a, b)
}

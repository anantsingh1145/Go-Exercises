package main

import (
	"fmt"
)

// This program demonstrates the use of the make function to create a slice of strings
// and the append function to add elements to the slice.
// The make function is used to create a slice with a specified length, and the append function
// is used to add elements to the slice, which can grow dynamically.
func main() {
	s1 := make([]string, 0, 5)               // Create a slice of strings with length 5
	fmt.Println("s1 before appending: ", s1) // Output: [     ] (5 empty strings)
	fmt.Println("Length of s1: ", len(s1))   // Output: 0
	fmt.Println("Capacity of s1: ", cap(s1)) // Output: 5

	s1 = append(s1, "a", "b", "c")                    // Append three strings to the slice
	fmt.Printf("s1 after appending a,b,c: %s \n", s1) // Output: [  a b c]

	s1 = append(s1, "d", "e", "f")                 // Append three more strings to the slice
	fmt.Printf("s1 after appending d,e,f: %s", s1) // Output: [  a b c d e f]

	// The slice can grow beyond its initial capacity, and the underlying array will be resized as needed.
	fmt.Println("\nLength of s1 after appending: ", len(s1)) // Output: 6
	fmt.Println("Capacity of s1 after appending: ", cap(s1))
	// Output: 10 (the capacity may increase as the slice grows, slice get multiple of 2,3,4 of initial capacity)
}

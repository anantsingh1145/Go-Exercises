package main

import (
	"fmt"
)

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Original Slice:", s)
	s = append(s[:2], s[3:]...) // Delete the value at index 2 "..." is used to unpack the slice
	// This creates a new slice that includes all elements before index 2 and all
	// elements after index 2, effectively removing the value at index 2.
	fmt.Println("Slice after deleting value at index 2:", s)
	s = append(s[:5], s[6:]...) // Delete the value at index 5
	// This creates a new slice that includes all elements before index 5 and all
	// elements after index 5, effectively removing the value at index 5.
	fmt.Println("Slice after deleting value at index 5:", s)
}
